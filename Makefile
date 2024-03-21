
SHELL:=/usr/bin/bash

HOSTNAME=github.com
NAMESPACE=thomasfinstad
NAME=vyos
BINARY_PREFIX=terraform-provider-
VERSION=1.5.0
VERSION_ROLLING=$$(date +%Y.%m.%d)
OS_ARCH=linux_amd64
BIN_DIR=dist

GO_IMPORT_ROOT:=${HOSTNAME}/${NAMESPACE}/terraform-provider-vyos

# TODO improve build situation
#  we need 1 for rolling releases and 1 for LTS 1.4 releases
#  figure out if these should be seperate repos, branches or what else
#  Also look into cleaning up this Makefil and its targets to streamline usage
#  and improve the help target to only show the ones intended for people
#  by default.

###
# Default helper target
help:
	@echo "Most targets are not ment for manual usage, these are the private targets starting with dot"
	@echo "Valid targets listed below"
	@echo "<target>: <dependency> <dependency ...>" | egrep --color '^[^ ]*:'
	@make -rpn | sed -n -e '/^$$/ { n ; /^[^ .#][^ ]*:/p ; }' | egrep --color '^[^ ]*:'

###
# VyOS ISO Upload timestamp

# Fetch and format newest successful rolling ISO build time
# TODO use release page for version and commit
#  ref: https://github.com/vyos/vyos-rolling-nightly-builds/releases
data/vyos/rolling-iso-build.txt:
	-mkdir -p .build
	curl -L \
		-H "Accept: application/vnd.github+json" \
		-H "X-GitHub-Api-Version: 2022-11-28" \
		--url "https://api.github.com/repos/vyos/vyos-rolling-nightly-builds/actions/runs?status=success&per_page=10" \
			| jq -r '[.workflow_runs[] | select(.name | "VyOS rolling nightly build")].[0].run_started_at' \
				> ".build/rolling-iso-build.txt"

	-head -n1 "data/vyos/rolling-iso-build.txt"
	-head -n1 ".build/rolling-iso-build.txt"

	if [ "$$(head -n1 "data/vyos/rolling-iso-build.txt")" != "$$(head -n1 ".build/rolling-iso-build.txt")" ]; then \
		echo "updating timestamp file"; \
		mv ".build/rolling-iso-build.txt" "data/vyos/rolling-iso-build.txt"; \
	fi

###
# VyOS src repo at correct commit
#  Requireing this also means the first step likely needs to be
#  `make git-submodules` to make sure the submodule is populated.
#  It is done this way because I do not know make well enough to
#  ensure the submodule exists without depending on a file within it
#  since that is problematic when it comes to the timestamps of the files
#  and pulls in more build steps than needed.
data/vyos/submodule.sha: data/vyos/rolling-iso-build.txt
	git submodule update --init --single-branch -- data/vyos/vyos-1x

	cd data/vyos/vyos-1x && \
	commit="$$(git rev-list --date=iso-strict -n 1 --before="$$(cat "../rolling-iso-build.txt")" "origin/current")" && \
	git checkout "$$commit" && \
	echo "$$commit" > ../submodule.sha

	git add data/vyos/vyos-1x

.PHONY:git-submodules
git-submodules: data/vyos/submodule.sha
	git submodule update --init --single-branch -- data/vyos/vyos-1x

###
# Autogenerate Schemas

# Convert from relaxng to XSD
data/vyos/schema/interface-definition.xsd: data/vyos/submodule.sha
	make git-submodules

	mkdir -p data/vyos/schema/
	java -jar tools/trang-20091111/trang.jar -I rnc -O xsd data/vyos/vyos-1x/schema/interface_definition.rnc data/vyos/schema/interface-definition.xsd
	xmllint --format --recover --output 'data/vyos/schema/interface-definition.xsd' 'data/vyos/schema/interface-definition.xsd'

# Generate go structs from XSD
internal/vyos/schema/interfacedefinition/autogen-structs.go: data/vyos/schema/interface-definition.xsd internal/vyos/schema/interfacedefinition/interface-definition.go

	@-rm -v internal/vyos/schema/interfacedefinition/autogen-structs.go

	# Generate structs from schema
	go run github.com/xuri/xgen/cmd/xgen -p interfacedefinition -i data/vyos/schema/interface-definition.xsd -o internal/vyos/schema/interfacedefinition/autogen-structs.go -l Go

	# Ensure the nodes name atter will be properly unmarshaled from xml
	sed -i 's|\*NodeNameAttr.*|string `xml:"name,attr,omitempty"`|' internal/vyos/schema/interfacedefinition/autogen-structs.go

	# Convert any undefined value as string type to stop unmarshaling from breaking
	sed -i 's|interface{}|string|' internal/vyos/schema/interfacedefinition/autogen-structs.go

	# Add option to set if this is used as a root node
	sed -i 's|\(type [A-Za-z]*Node struct {\)|\1\nIsBaseNode bool|' internal/vyos/schema/interfacedefinition/autogen-structs.go

	# Add a parent value to node structs
	sed -i 's|\(type [A-Za-z]*Node struct {\)|\1\nParent NodeParent|' internal/vyos/schema/interfacedefinition/autogen-structs.go

	# Format output
	gofumpt -w ./internal/vyos/schema/interfacedefinition/


###
# Terraform Resource Schemas

# Compile interface devfinitions
data/vyos/interface-definitions: data/vyos/submodule.sha
	make git-submodules

	@rm -rf "data/vyos/interface-definitions"
	python -m venv .build/vyos/vyos-1x/venv

	bash -c " \
		source .build/vyos/vyos-1x/venv/bin/activate && \
		cd data/vyos/vyos-1x && \
		pip install -r test-requirements.txt \
	"

	# A bit unfortunate, but we have to "extract" python requirements from the deb control to be able to make the interface defs
	# as they are not all listed in the test-requirements.txt
	bash -c " \
		source .build/vyos/vyos-1x/venv/bin/activate && \
		cd data/vyos/vyos-1x && \
		sed -r -n '/python3*-/s/^[[:space:]]*python3*-(\w*).*/\1/p' debian/control | xargs -n1 pip install \
	"

	# Seems the build script no longer works cleanly without an assumed library function:
	# /usr/lib/libvyosconfig.so.0: cannot open shared object file: No such file or directory
	# TODO investigate blessed method of building interface definitions
	-bash -c " \
		source .build/vyos/vyos-1x/venv/bin/activate && \
		cd data/vyos/vyos-1x && \
		make interface_definitions \
	"

	mv data/vyos/vyos-1x/build/interface-definitions data/vyos/interface-definitions
	find data/vyos/interface-definitions/ -type f -name "*.xml" -execdir xmllint --format --recover --output '{}' '{}' \;

internal/vyos/vyosinterface/auto-package.go:  $(shell find data/vyos/interface-definitions/ -type f) $(shell find tools/build-vyos-infterface-definition-structs/ -type f) internal/vyos/schema/interfacedefinition/autogen-structs.go
	mkdir -p "internal/vyos/vyosinterface"

	@rm -fv internal/vyos/vyosinterface/auto-package.go
	@rm -fv internal/vyos/vyosinterface/autogen-*.go

	# Generate interfaces, skip xml component version metadata file
	for xmlFile in $$(ls "data/vyos/interface-definitions/" | grep -v "xml-component-version.xml"); do \
		echo -en "Input xml: '$${xmlFile}'\t"; \
		go run tools/build-vyos-infterface-definition-structs/*.go "data/vyos/interface-definitions/$${xmlFile}" "internal/vyos/vyosinterface" "vyosinterface" || exit 1; \
	done

	echo -e "// Package vyosinterface generated by Makefile. DO NOT EDIT." > "internal/vyos/vyosinterface/auto-package.go"
	echo -e "package vyosinterface\n\n" >> "internal/vyos/vyosinterface/auto-package.go"

	echo -e 'import "$(GO_IMPORT_ROOT)/internal/vyos/schema/interfacedefinition"\n\n' >> "internal/vyos/vyosinterface/auto-package.go"

	echo -e "// GetInterfaces returns all autogenerated interface definitions" >> "internal/vyos/vyosinterface/auto-package.go"
	echo -e "func GetInterfaces() []interfacedefinition.InterfaceDefinition {" >> "internal/vyos/vyosinterface/auto-package.go"
	echo -e "return []interfacedefinition.InterfaceDefinition{" >> "internal/vyos/vyosinterface/auto-package.go"
	grep -r -o -h "func [a-z]*()" internal/vyos/vyosinterface/| sort | sed 's/func \(.*\)/\1,/g' >> "internal/vyos/vyosinterface/auto-package.go"
	echo -e "}" >> "internal/vyos/vyosinterface/auto-package.go"
	echo -e "}" >> "internal/vyos/vyosinterface/auto-package.go"

	gofumpt -w ./internal/vyos/vyosinterface/*.go

internal/terraform/resource/autogen/timestamp.txt: \
													internal/vyos/vyosinterface/auto-package.go \
													$(shell find internal/vyos/schema/interfacedefinition/ -type f) \
													$(shell find internal/vyos/vyosinterface/ -type f) \
													$(shell find tools/build-terraform-resource-full/ -type f)
	# Prep dirs
	rm -rf internal/terraform/resource/autogen
	mkdir -p "internal/terraform/resource/autogen/named"
	mkdir -p "internal/terraform/resource/autogen/global"

	# Generate resources
	go run tools/build-terraform-resource-full/main.go \
		internal/terraform/resource/autogen \
		$(GO_IMPORT_ROOT) \
		"system/conntrack-timeout-custom-rule"

	# Clean up code
	gofumpt -w internal/terraform/resource/autogen/
	goimports -w "./internal/terraform/resource/autogen"

	date > "internal/terraform/resource/autogen/timestamp.txt"

## Ref: https://stackoverflow.com/a/45003119
# problems of method 2:
# 	If an argument has same name as an existing target then make will print a warning that it is being overwritten.
# 	no workaround that I know of
#
# 	Arguments with an equal sign will still be interpreted by make and not passed
# 	no workaround
#
# 	Arguments with spaces is still awkward
# 	no workaround
#
# 	Arguments with space breaks eval trying to create do nothing targets.
# 	workaround: create the global catch all target doing nothing as above. with the problem as above that it will again silently ignore mistyped legitimate targets.
#
# 	it uses eval to modify the makefile at runtime. how much worse can you go in terms of readability and debugability and the Principle of least astonishment.
# 	workaround: don't!
#
# If the first argument is "test"...
# Example usage: make test -- -test.run TestPolicyAccessListEmptyResponseFromApi
ifeq (test,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for target
  INPUT_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(INPUT_ARGS):;@:)
endif
test: Makefile internal/terraform/resource/autogen/timestamp.txt $(shell find . -type f -name "*.go")
	@echo Input Args: $(INPUT_ARGS)

	# VyOS API can often take ~1 second to respond to a configure request.
	# This means we attempt to tune retrys and delays around this.
	# The end result is that to be able to test retry functionality we will need a bit of head room,
	# so 5s timeout should be plenty for any test by using a context with 2 or 3 seconds timeout.
	go test -count=1 -failfast -timeout 30s $(shell find . -type f -name "*_test.go" -exec dirname {} \; | sort -u) $(INPUT_ARGS)

	# Caching timestamp
	@date > test

update-rolling:
	make --always-make data/vyos/rolling-iso-build.txt

build-rolling: test
	-rm -rf "${BIN_DIR}"
	-mkdir -p "${BIN_DIR}/local/providers/${NAME}/${VERSION_ROLLING}/${OS_ARCH}/"
	go build -o ${BIN_DIR}/local/providers/${NAME}/${VERSION_ROLLING}/${OS_ARCH}/${BINARY_PREFIX}${NAME} $(RUN_ARGS)

	# Caching timestamp
	@date > build-rolling

# .PHONY: build
# build: test
# 	-rm -rf ${BIN_DIR}
# 	-mkdir -p "${BIN_DIR}/local/providers/${NAME}/${VERSION}/${OS_ARCH}/"
# 	go build -o ${BIN_DIR}/local/providers/${NAME}/${VERSION}/${OS_ARCH}/${BINARY_PREFIX}${NAME}

.PHONY: clean
clean:
	rm -rfv .build
	git submodule deinit --all -f

.PHONY: provider-schema
provider-schema: build-rolling
	mkdir -p data/provider-schema
	cd examples/provider && make init
	terraform -chdir=examples/provider providers schema -json | jq '.' > data/provider-schema/${VERSION_ROLLING}.json

	# if diff data/provider-schema/$(shell ls data/provider-schema | sort -V | tail -n1) data/provider-schema/${VERSION_ROLLING}.json; then \
	# 	rm -v data/provider-schema/${VERSION_ROLLING}.json; \
	# fi

docs/index.md: \
				build-rolling \
				internal/terraform/resource/autogen/timestamp.txt \
				$(shell find templates/ -type f)

	# Prep dirs
	rm -rf "docs/"

	# Create docs
	tfplugindocs generate

.PHONY: stage
stage: docs/index.md provider-schema
	git add -A
	-pre-commit run
	git add -A
