
SHELL:=/usr/bin/bash

HOSTNAME=github.com
NAMESPACE=thomasfinstad
NAME=vyos
BINARY_PREFIX=terraform-provider-
VERSION=$(shell cut -d"T" -f1 data/vyos-1x-info.txt | tr '-' '.')
OS_ARCH=linux_amd64
BIN_DIR=dist
GO_IMPORT_ROOT:=${HOSTNAME}/${NAMESPACE}/terraform-provider-${NAME}-rolling
ADDRESS := "registry.terraform.io/${NAMESPACE}/${NAME}"

# TODO support rolling and LTS builds
#  we need 1 for rolling releases and 1 for LTS 1.4 releases
#  figure out if these should be separate repos, branches or what else
#  milestone: 5

###
# Default helper target
# TODO improve Makefile help
#  should print only targets intended for human usage
#  should print preceding comment on those targets when spessific help is requested
#  should print all targets when requested
#  milestone: 3
help:
	@echo "Most targets are not meant for manual usage, these are the private targets starting with dot"
	@echo "Valid targets listed below"
	@echo "<target>: <dependency> <dependency ...>" | egrep --color '^[^ ]*:'
	@make -rpn | sed -n -e '/^$$/ { n ; /^[^ .#][^ ]*:/p ; }' | egrep --color '^[^ ]*:'


# Fetch and format newest successful rolling ISO build time
# TODO use release page for version and commit
#  ref: https://vyos.dev/T6156
#  ref: https://github.com/vyos/vyos-rolling-nightly-builds/releases
#  milestone: 4
data/vyos-1x-info.txt:
	# data/vyos-1x-info.txt

	-mkdir -p data
	-mkdir -p .build

	curl -L \
		-H "Accept: application/vnd.github+json" \
		-H "X-GitHub-Api-Version: 2022-11-28" \
		--url "https://api.github.com/repos/vyos/vyos-rolling-nightly-builds/actions/runs?status=success&per_page=10" \
			| jq -r '[.workflow_runs[] | select(.name | "VyOS rolling nightly build")].[0].run_started_at' \
				> ".build/latest-vyos-repo-version.txt"

	if [ ! -f data/vyos-1x-info.txt ]; then \
		mv ".build/latest-vyos-repo-version.txt" "data/vyos-1x-info.txt"; \
	elif ! diff -w "data/vyos-1x-info.txt" ".build/latest-vyos-repo-version.txt" > /dev/null; then \
		echo "updating timestamp file"; \
		mv ".build/latest-vyos-repo-version.txt" "data/vyos-1x-info.txt"; \
	fi

###
# VyOS src repo at correct commit
.build/vyos-1x: data/vyos-1x-info.txt
	# .build/vyos-1x

	-mkdir -p .build

	# Clone repo if needed
	if [ ! -d ".build/vyos-1x" ]; then \
		cd .build && \
		git clone "https://github.com/vyos/vyos-1x.git" vyos-1x; \
	fi

	# Cache the commit sha
	cd .build/vyos-1x && \
	git rev-list --date=iso-strict -n 1 --before="$$(cat "../../data/vyos-1x-info.txt")" "origin/current" > ../vyos-1x.sha

	# Checkout the commit
	cd .build/vyos-1x && \
	git checkout master && \
	git pull && \
	git checkout "$$(cat ../vyos-1x.sha)"

###
# Autogenerate Schemas

# Convert from relaxng to XSD
.build/schema-definitions.xsd: data/vyos-1x-info.txt |.build/vyos-1x
	-mkdir -p .build

	java -jar tools/trang-20091111/trang.jar -I rnc -O xsd .build/vyos-1x/schema/interface_definition.rnc .build/schema-definitions.xsd
	xmllint --format --recover --output '.build/schema-definitions.xsd' '.build/schema-definitions.xsd'

# Generate go structs from XSD
internal/vyos/schemadefinition/autogen-structs.go: data/vyos-1x-info.txt internal/vyos/schemadefinition/interface-definition.go |.build/schema-definitions.xsd

	@-rm -v internal/vyos/schemadefinition/autogen-structs.go
	-mkdir -p internal/vyos/schemadefinition
	-mkdir .build

	# Generate structs from schema
	go run github.com/xuri/xgen/cmd/xgen -p schemadefinition -i .build/schema-definitions.xsd -o internal/vyos/schemadefinition/autogen-structs.go -l Go

	# TODO convert from sed mangling to go mangling
	#  create a go tool package and use dst in the same way
	#  as tools/build-vyos-infterface-definition-structs/main.go
	#  milestone: 6

	# Ensure the nodes name atter will be properly unmarshaled from xml
	sed -i 's|\*NodeNameAttr.*|string `xml:"name,attr,omitempty"`|' internal/vyos/schemadefinition/autogen-structs.go

	# Convert any undefined value as string type to stop unmarshalling from breaking
	sed -i 's|interface{}|string|' internal/vyos/schemadefinition/autogen-structs.go

	# Add option to set if this is used as a root node
	sed -i 's|\(type [A-Za-z]*Node struct {\)|\1\nIsBaseNode bool|' internal/vyos/schemadefinition/autogen-structs.go

	# Add a parent value to node structs
	sed -i 's|\(type [A-Za-z]*Node struct {\)|\1\nParent NodeParent|' internal/vyos/schemadefinition/autogen-structs.go

	# Format output
	gofumpt -w ./internal/vyos/schemadefinition/

# Compile interface devfinitions
.build/interface-definitions: data/vyos-1x-info.txt |.build/vyos-1x
	@rm -rf ".build/interface-definitions"
	-mkdir -p .build

	# Making a temporary copy means we don't touch
	# the files in the directory, which would cause make to
	# see vyos-1x as newer than the interface-definitions
	# causing interface-definitions to always rebuild
	cp -a .build/vyos-1x .build/vyos-1x-tmp

	# Build interface definitions using the vyos build container.
	docker run --rm -v .:/repo -w /repo docker.io/vyos/vyos-build:current bash -c "cd .build/vyos-1x-tmp && make interface_definitions"

	# Clean up the tmp directory
	mv .build/vyos-1x-tmp/build/interface-definitions .build/interface-definitions
	rm -rf .build/vyos-1x-tmp

	# Pretty format the XML files incase a human needs to inspect them
	find .build/interface-definitions/ -type f -name "*.xml" -execdir xmllint --format --recover --output '{}' '{}' \;

.build/vyosinterfaces/auto-package.go: data/vyos-1x-info.txt internal/vyos/schemadefinition/autogen-structs.go tools/build-vyos-infterface-definition-structs |.build/interface-definitions
	mkdir -p ".build/vyosinterfaces"

	@rm -fv .build/vyosinterfaces/auto-package.go
	@rm -fv .build/vyosinterfaces/autogen-*.go

	# Generate interfaces, skip xml component version metadata file
	for xmlFile in $$(ls ".build/interface-definitions/" | grep -v "xml-component-version.xml"); do \
		echo -en "Input xml: '$${xmlFile}'\t"; \
		go run tools/build-vyos-infterface-definition-structs/*.go ".build/interface-definitions/$${xmlFile}" ".build/vyosinterfaces" "vyosinterface" || exit 1; \
	done

	# TODO clean up generation of vyosinterfaces/auto-package.go
	#  either use HEREDOC to make it prettier or create a gotemplate
	#  milestone: 6
	echo -e "// Autogenerated by Makefile. DO NOT EDIT." > ".build/vyosinterfaces/auto-package.go"
	echo -e "package vyosinterface\n\n" >> ".build/vyosinterfaces/auto-package.go"

	echo -e 'import "$(GO_IMPORT_ROOT)/internal/vyos/schemadefinition"\n\n' >> ".build/vyosinterfaces/auto-package.go"

	echo -e "// GetInterfaces returns all autogenerated interface definitions" >> ".build/vyosinterfaces/auto-package.go"
	echo -e "func GetInterfaces() []schemadefinition.InterfaceDefinition {" >> ".build/vyosinterfaces/auto-package.go"
	echo -e "return []schemadefinition.InterfaceDefinition{" >> ".build/vyosinterfaces/auto-package.go"
	grep -r -o -h "func [a-z]*()" .build/vyosinterfaces/| sort | sed 's/func \(.*\)/\1,/g' >> ".build/vyosinterfaces/auto-package.go"
	echo -e "}" >> ".build/vyosinterfaces/auto-package.go"
	echo -e "}" >> ".build/vyosinterfaces/auto-package.go"

	gofumpt -w ./.build/vyosinterfaces/*.go

internal/terraform/resource/autogen: data/vyos-1x-info.txt |.build/vyosinterfaces/auto-package.go
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
test:	internal/terraform/resource/autogen \
		$(shell find . -type f -name "*_test.go") \
		$(shell find . -type f -name "*.go" -not \( -path "*autogen*" -or -path "*/.build/*" -or -path "*/tools/*" \) )

	@echo Input Args: $(INPUT_ARGS)

	# VyOS API can often take ~1 second to respond to a configure request.
	# This means we attempt to tune retrys and delays around this.
	# The end result is that to be able to test retry functionality we will need a bit of head room,
	# so 30s timeout should be enough for all tests combined.

	go test -count=1 -failfast -timeout 30s $(shell find . -type f -name "*_test.go" -exec dirname {} \; | sort -u) $(INPUT_ARGS)

	# Caching timestamp
	@date > test

update-rolling:
	make --always-make data/vyos-1x-info.txt

build: test
	-rm -rf "${BIN_DIR}"
	-mkdir -p "${BIN_DIR}/local/providers/${NAME}/${VERSION}/${OS_ARCH}/"
	go build \
		-ldflags "-X main.version=$(VERSION) -X main.address=${ADDRESS}" \
		-o "${BIN_DIR}/local/providers/${NAME}/${VERSION}/${OS_ARCH}/${BINARY_PREFIX}${NAME}"

	# Caching timestamp
	@date > build

docs/index.md: \
				build \
				internal/terraform/resource/autogen \
				$(shell find templates/ -type f)

	# Prep dirs
	rm -rf "docs/"

	# Create docs
	tfplugindocs generate --provider-name vyos

data/provider-schema/$(VERSION).json: build
	mkdir -p data/provider-schema
	cd examples/provider && make init
	terraform -chdir=examples/provider providers schema -json | jq '.' > data/provider-schema/${VERSION}.json

	# if diff data/provider-schema/$(shell ls data/provider-schema | sort -V | tail -n1) data/provider-schema/${VERSION}.json; then \
	# 	rm -v data/provider-schema/${VERSION}.json; \
	# fi

#CHANGELOG.md:
	#go run tools/generate-changelog/*.go data/provider-schema/2024.03.21.json data/provider-schema/2024.03.22.json

.PHONY: clean
clean:
	-rm -rf dist
	-rm -rf .build
	-rm test
	-rm build

.PHONY: full
full: build docs/index.md data/provider-schema/$(VERSION).json
	-pre-commit run --all-files

.PHONY: why
why:
	@make -and full \
		| sed -rn 's/^ *([A-Za-z ]*Must remake target.*| Prerequisite.*is newer than.*)/\1/p' \
			| tac

.PHONY: diff
diff:
	git diff --word-diff --word-diff-regex='\w+' -- . ':!docs' ':!data/provider-schema' ':!*autogen*'
