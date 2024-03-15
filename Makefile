
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
data/vyos/submodule.sha: data/vyos/rolling-iso-build.txt
	git submodule update --init --single-branch -- data/vyos/vyos-1x

	cd data/vyos/vyos-1x && \
	commit="$$(git rev-list --date=iso-strict -n 1 --before="$$(cat "../rolling-iso-build.txt")" "origin/current")" && \
	git checkout "$$commit" && \
	echo "$$commit" > ../submodule.sha

	git add data/vyos/vyos-1x

data/vyos/Makefile: data/vyos/submodule.sha
	git submodule update --init --single-branch -- data/vyos/vyos-1x

###
# Autogenerate Schemas

# Convert from relaxng to XSD
data/vyos/schema/interface-definition.xsd: data/vyos/submodule.sha
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
data/vyos/interface-definitions: data/vyos/submodule.sha data/vyos/Makefile

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

internal/vyos/vyosinterface/auto-package.go: data/vyos/interface-definitions tools/build-vyos-infterface-definition-structs/main.go
	mkdir -p "internal/vyos/vyosinterface"

	@rm -fv internal/vyos/vyosinterface/auto-package.go
	@rm -fv internal/vyos/vyosinterface/autogen-*.go

	# Generate interfaces, skip xml component version metadata file
	for xmlFile in $$(ls "data/vyos/interface-definitions/" | grep -v "xml-component-version.xml"); do \
		echo -en "Input xml: '$${xmlFile}'\t"; \
		go run tools/build-vyos-infterface-definition-structs/*.go "data/vyos/interface-definitions/$${xmlFile}" "internal/vyos/vyosinterface" "vyosinterface" || exit 1; \
	done

	echo -e "// Package vyosinterface generated by Makefile on $$(date -u -Iseconds). DO NOT EDIT." > "internal/vyos/vyosinterface/auto-package.go"
	echo -e "package vyosinterface\n\n" >> "internal/vyos/vyosinterface/auto-package.go"

	echo -e 'import "$(GO_IMPORT_ROOT)/internal/vyos/schema/interfacedefinition"\n\n' >> "internal/vyos/vyosinterface/auto-package.go"

	echo -e "// GetInterfaces returns all autogenerated interface definitions" >> "internal/vyos/vyosinterface/auto-package.go"
	echo -e "func GetInterfaces() []interfacedefinition.InterfaceDefinition {" >> "internal/vyos/vyosinterface/auto-package.go"
	echo -e "return []interfacedefinition.InterfaceDefinition{" >> "internal/vyos/vyosinterface/auto-package.go"
	grep -r -o -h "func [a-z]*()" internal/vyos/vyosinterface/| sort | sed 's/func \(.*\)/\1,/g' >> "internal/vyos/vyosinterface/auto-package.go"
	echo -e "}" >> "internal/vyos/vyosinterface/auto-package.go"
	echo -e "}" >> "internal/vyos/vyosinterface/auto-package.go"

	gofumpt -w ./internal/vyos/vyosinterface/*.go

internal/terraform/resource/autogen: internal/vyos/vyosinterface/auto-package.go tools/build-terraform-resource-full/named-template.gotmpl tools/build-terraform-resource-full/global-template.gotmpl

	# Prep dirs
	@rm -rfv internal/terraform/resource/autogen
	@rm -rfv docs/resources
	mkdir -p "internal/terraform/resource/autogen/named"
	mkdir -p "internal/terraform/resource/autogen/global"

	# Generate code
	go run tools/build-terraform-resource-full/main.go \
		internal/terraform/resource/autogen \
		$(GO_IMPORT_ROOT) \
		"system/conntrack-timeout-custom-rule"

	# Clean up code
	gofumpt -w internal/terraform/resource/autogen/
	goimports -w "./internal/terraform/resource/autogen"

	# Create docs
	go generate main.go
	# https://github.com/hashicorp/terraform-plugin-docs/issues/156
	for f in $$(find docs -name "*.md"); do \
		SC=$$(basename "$$f" | cut -d_ -f1); \
		sed -i "s|subcategory:.*|subcategory: \"$$SC\"|" "$$f"; \
	done

.PHONY: test
test: internal/terraform/resource/autogen
	go test -failfast -timeout 5s ./internal/terraform/tests/... ./internal/terraform/helpers/...

.PHONY: build-rolling
build-rolling:
	make clean
	make --always-make data/vyos/rolling-iso-build.txt
	make test

	-rm -rf "${BIN_DIR}"
	-mkdir -p "${BIN_DIR}/local/providers/${NAME}/${VERSION_ROLLING}/${OS_ARCH}/"
	go build -o ${BIN_DIR}/local/providers/${NAME}/${VERSION_ROLLING}/${OS_ARCH}/${BINARY_PREFIX}${NAME}

	make clean

.PHONY: build
build: test
	-rm -rf ${BIN_DIR}
	-mkdir -p "${BIN_DIR}/local/providers/${NAME}/${VERSION}/${OS_ARCH}/"
	go build -o ${BIN_DIR}/local/providers/${NAME}/${VERSION}/${OS_ARCH}/${BINARY_PREFIX}${NAME}

.PHONY: clean
clean:
	rm -rfv .build
	git submodule deinit --all -f

.PHONY: example-clean
example-clean:
	-rm -rfv ~/.terraform.d/plugins/

.PHONY: provider-schema
provider-schema: build-rolling
	mkdir -p data/provider-schema
	cd examples/provider; make init
	terraform -chdir=examples/provider providers schema -json | jq '.' > data/provider-schema/${VERSION_ROLLING}.json

	# if diff data/provider-schema/$(shell ls data/provider-schema | sort -V | tail -n1) data/provider-schema/${VERSION_ROLLING}.json; then \
	# 	rm -v data/provider-schema/${VERSION_ROLLING}.json; \
	# fi

.PHONY: stage
stage: provider-schema
	git add -A
	-pre-commit run
	git add -A
