
SHELL := bash
.SHELLFLAGS := -eu -o pipefail -c
.ONESHELL:
# ifeq ($(origin .RECIPEPREFIX), undefined)
#   $(error This Make does not support .RECIPEPREFIX. Please use GNU Make 4.0 or later)
# endif
# .RECIPEPREFIX = >
.DELETE_ON_ERROR:
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules
ROOT_DIR := $(CURDIR)

HOSTNAME=github.com
NAMESPACE=thomasfinstad
PROVIDER_NAME=vyos-rolling
DIST_DIR="/dist"
GO_IMPORT_ROOT=${HOSTNAME}/${NAMESPACE}/terraform-provider-${PROVIDER_NAME}
ADDRESS="registry.terraform.io/${NAMESPACE}/${PROVIDER_NAME}"
BUILD_ARCH=amd64 386 arm arm64
BUILD_OS=freebsd windows linux darwin

default: help

#####
# Dynamic parameter hacks
#####

# Target: why
ifeq (why,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for target
  INPUT_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(INPUT_ARGS):;@:)
endif

# Target test
ifeq (test,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for target
  INPUT_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(INPUT_ARGS):;@:)
endif

######
# Targets
######

# TODO support rolling and LTS builds
#  we need 1 for rolling releases and 1 for LTS 1.4 releases
#  figure out if these should be separate repos, branches or what else
#  milestone: 5


# Fetch and format newest successful rolling ISO build time
# ref: https://vyos.dev/T6156
data/vyos-1x-info.txt:
	@echo Make data/vyos-1x-info.txt

	mkdir -p data || true
	-mkdir -p .build || true

	curl -L \
		-H "Accept: application/vnd.github+json" \
		-H "X-GitHub-Api-Version: 2022-11-28" \
		--url "https://api.github.com/repos/vyos/vyos-rolling-nightly-builds/actions/runs?status=success&per_page=10" \
			| jq -r '[.workflow_runs[] | select(.name | "VyOS rolling nightly build")] | first .run_started_at' \
				> ".build/latest-vyos-repo-version.txt"

	echo "Currently based on version: $$(cat data/vyos-1x-info.txt)"
	echo "Newest built version: $$(cat .build/latest-vyos-repo-version.txt)"

	if [ ! -f data/vyos-1x-info.txt ]; then
		echo "creating timestamp file";
		mv ".build/latest-vyos-repo-version.txt" "data/vyos-1x-info.txt";
	elif ! diff -w "data/vyos-1x-info.txt" ".build/latest-vyos-repo-version.txt" > /dev/null; then
		echo "updating timestamp file";
		mv ".build/latest-vyos-repo-version.txt" "data/vyos-1x-info.txt";
	else
		echo "keeping current version";
	fi

###
# VyOS src repo at correct commit
.build/vyos-1x: data/vyos-1x-info.txt
	@echo Make .build/vyos-1x

	mkdir -p .build || true

	# Clone repo if needed
	if [ ! -d ".build/vyos-1x" ]; then
		cd .build
		git clone "https://github.com/vyos/vyos-1x.git" vyos-1x;
		cd $(ROOT_DIR)
	fi

	# Cache the commit sha
	cd .build/vyos-1x
	git rev-list --date=iso-strict -n 1 --before="$$(cat "../../data/vyos-1x-info.txt")" "origin/current" > ../vyos-1x.sha

	# Checkout the commit
	git checkout current
	git pull
	git checkout "$$(cat ../vyos-1x.sha)"

###
# Autogenerate Schemas

# Convert from relaxng to XSD
.build/schema-definitions.xsd: data/vyos-1x-info.txt |.build/vyos-1x
	@echo Make .build/schema-definitions.xsd

	mkdir -p .build || true

	java -jar tools/trang-20091111/trang.jar -I rnc -O xsd .build/vyos-1x/schema/interface_definition.rnc .build/schema-definitions.xsd
	xmllint --format --recover --output '.build/schema-definitions.xsd' '.build/schema-definitions.xsd'

# Generate go structs from XSD
internal/vyos/schemadefinition/autogen-structs.go: data/vyos-1x-info.txt internal/vyos/schemadefinition/interface-definition.go .build/schema-definitions.xsd
	@echo Make internal/vyos/schemadefinition/autogen-structs.go

	go get

	rm -v internal/vyos/schemadefinition/autogen-structs.go || true
	mkdir -p internal/vyos/schemadefinition || true
	mkdir .build || true

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
	@echo Make .build/interface-definitions

	rm -rf ".build/interface-definitions"
	mkdir -p .build || true

	# Making a temporary copy means we don't touch
	# the files in the directory, which would cause make to
	# see vyos-1x as newer than the interface-definitions
	# causing interface-definitions to always rebuild
	sudo docker volume create --name repo
	sudo docker container create --name make-interface-definitions -v repo:/docker-volume busybox
	sudo docker cp .build/vyos-1x make-interface-definitions:/docker-volume

	# Build interface definitions using the vyos build container.
	sudo docker run --rm -v repo:/docker-volume -w /docker-volume/vyos-1x docker.io/vyos/vyos-build:current make interface_definitions

	# Clean up the tmp resources
	sudo docker cp make-interface-definitions:/docker-volume/vyos-1x/build/interface-definitions .build/interface-definitions
	sudo docker rm make-interface-definitions
	sudo docker volume rm repo

	sudo chown -R $$(id -u) .build/interface-definitions

	# Pretty format the XML files incase a human needs to inspect them
	find .build/interface-definitions/ -type f -name "*.xml" -execdir xmllint --format --recover --output '{}' '{}' \;

internal/vyos/vyosinterfaces/autogen.go: data/vyos-1x-info.txt internal/vyos/schemadefinition/autogen-structs.go tools/build-vyos-infterface-definition-structs .build/interface-definitions
	@echo Make internal/vyos/vyosinterfaces/autogen.go

	mkdir -p "internal/vyos/vyosinterfaces"

	rm -f internal/vyos/vyosinterfaces/autogen.go
	rm -f internal/vyos/vyosinterfaces/autogen-*.go

	# Generate interfaces, skip xml component version metadata file
	for xmlFile in $$(ls ".build/interface-definitions/" | grep -v "xml-component-version.xml"); do
		echo -en "Input xml: '$${xmlFile}'\t";
		go run tools/build-vyos-infterface-definition-structs/*.go \
			".build/interface-definitions/$${xmlFile}" \
			"internal/vyos/vyosinterfaces" \
			"vyosinterface"
		echo;
	done

	# TODO clean up generation of vyosinterfaces/autogen.go
	#  either use HEREDOC to make it prettier or create a gotemplate
	#  milestone: 6
	echo -e "// Autogenerated by Makefile. DO NOT EDIT." > "internal/vyos/vyosinterfaces/autogen.go"
	echo -e "package vyosinterface\n\n" >> "internal/vyos/vyosinterfaces/autogen.go"

	echo -e 'import "$(GO_IMPORT_ROOT)/internal/vyos/schemadefinition"\n\n' >> "internal/vyos/vyosinterfaces/autogen.go"

	echo -e "// GetInterfaces returns all autogenerated interface definitions" >> "internal/vyos/vyosinterfaces/autogen.go"
	echo -e "func GetInterfaces() []schemadefinition.InterfaceDefinition {" >> "internal/vyos/vyosinterfaces/autogen.go"
	echo -e "return []schemadefinition.InterfaceDefinition{" >> "internal/vyos/vyosinterfaces/autogen.go"
	grep -r -o -h "func [a-z]*()" internal/vyos/vyosinterfaces/| sort | sed 's/func \(.*\)/\1,/g' >> "internal/vyos/vyosinterfaces/autogen.go"
	echo -e "}" >> "internal/vyos/vyosinterfaces/autogen.go"
	echo -e "}" >> "internal/vyos/vyosinterfaces/autogen.go"

	gofumpt -w ./internal/vyos/vyosinterfaces/*.go

internal/terraform/resource/autogen: data/vyos-1x-info.txt internal/vyos/vyosinterfaces/autogen.go
	@echo Make internal/terraform/resource/autogen

	# Prep dirs
	rm -rf internal/terraform/resource/autogen
	mkdir -p "internal/terraform/resource/autogen/named"
	mkdir -p "internal/terraform/resource/autogen/global"

	go run tools/build-terraform-resource-full/main.go \
		internal/terraform/resource/autogen \
		$(GO_IMPORT_ROOT) \
		"system/conntrack-timeout-custom-rule"

	# Clean up code
	gofumpt -w internal/terraform/resource/autogen/
	goimports -w "./internal/terraform/resource/autogen"

# Pretty-name target for human usage
generate: internal/terraform/resource/autogen
	@echo Make generate

	# Caching timestamp
	date > generate

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
test:	internal/terraform/resource/autogen/ \
		$(shell find . -type f -name "*_test.go"  -not \( -path "*/.build/*" \) ) \
		$(shell find . -type f -name "*.go" -not \( -path "*autogen*" -or -path "*/.build/*" -or -path "*/tools/*" \) )

	@echo Make test

	echo Input Args: $(INPUT_ARGS)

	# VyOS API can often take ~1 second to respond to a configure request.
	# This means we attempt to tune retrys and delays around this.
	# The end result is that to be able to test retry functionality we will need a bit of head room,
	# so 30s timeout should be enough for all tests combined.

	go test -count=1 -failfast -timeout 30s $(shell find . -type f -name "*_test.go" -exec dirname {} \; | sort -u) $(INPUT_ARGS)

	# Caching timestamp
	date > test

build: test
	@echo Make build

	rm -rf "${DIST_DIR}/providers.localhost" || true

	version=$(shell date +%Y%m%d.%H%M.%S);
	for os in $(BUILD_OS); do
		for arch in $(BUILD_ARCH); do
			if [ "$${os}/$${arch}" == "darwin/arm" -o "$${os}/$${arch}" == "darwin/386" ]; then
				echo "Skipping unsupported os/arch combination: $${os}/$${arch}";
				continue;
			fi;
			echo "Building for $${os} ($${arch})"
			build_dir="${DIST_DIR}/providers.localhost/dev/$(PROVIDER_NAME)/$${version}/$${os}_$${arch}"

			mkdir -p "$${build_dir}/";
			go build \
				-ldflags "-X main.version=$${version} -X main.address=${ADDRESS}" \
				-o "$${build_dir}/terraform-provider-$(PROVIDER_NAME)_v$${version}";
		done;
	done;

	# Caching timestamp
	date > build

docs/index.md: \
				build \
				internal/terraform/resource/autogen \
				$(shell find templates/ -type f)
	@echo Make docs/index.md

	# Prep dirs
	rm -rf "docs/"

	# Create docs
	tfplugindocs generate --provider-name vyos > /dev/null

.PHONY: version
version:
	@echo Make version

	if [ -n "$(shell git status -s | head)" ]; then
		git status -s | head;
		echo "Can not create version when git is not in a clean, committed, and pushed state" >&2;
		exit 1;
	fi

	mkdir -p ".build" || true

	cd examples/provider
	make init
	terraform providers schema -json | jq '.' > ../../.build/provider-schema.json
	cd $(ROOT_DIR)

	pwd
	git tag -l
	cd tools/generate-changelog && go run *.go ../../.build/provider-schema.json ../../data/provider-schema.json
	cd $(ROOT_DIR)

	mv .build/provider-schema.json data/provider-schema.json
	mv CHANGELOG.md .build/CHANGELOG.md.old
	cat .build/CHANGELOG.md > CHANGELOG.md
	tail -n +2  .build/CHANGELOG.md.old >> CHANGELOG.md
	grep "^##" CHANGELOG.md | head -n1 | cut -d" " -f2 > VERSION

	git add -A
	-pre-commit run
	git add -A

	git commit -m "chore: Prepare for release v$$(cat VERSION)"
	git tag "v$$(cat VERSION)"

release: version
	@echo Make release
	# This is depriciated, release should be done via github actions

	rm -rf "${DIST_DIR}/publish" || true
	mkdir -p "${DIST_DIR}/publish" || true

	version="$$(cat VERSION)";
	build_dir=".build/release-build";
	mkdir -p "$${build_dir}/";
	for os in $(BUILD_OS); do
		for arch in $(BUILD_ARCH); do
			if [ "$${os}/$${arch}" == "darwin/arm" -o "$${os}/$${arch}" == "darwin/386" ]; then
				echo "Skipping unsupported os/arch combination: $${os}/$${arch}";
				continue;
			fi;
			echo -n "Packaging for $${os} ($${arch})"
			pub_dir="${DIST_DIR}/publish"

			go build \
				-ldflags "-X main.version=$${version} -X main.address=${ADDRESS}" \
				-o "$${build_dir}/terraform-provider-$(PROVIDER_NAME)_v$${version}";
			cd "$${build_dir}"
			zip \
				"$${pub_dir}/terraform-provider-$(PROVIDER_NAME)_$${version}_$${os}_$${arch}.zip" \
				"terraform-provider-$(PROVIDER_NAME)_v$${version}"
			rm "terraform-provider-$(PROVIDER_NAME)_v$${version}"
			cd "$(ROOT_DIR)";
		done;
	done;

	echo '{"version":1,"metadata":{"protocol_versions":["6.0"]}}' > "dist/publish/terraform-provider-$(PROVIDER_NAME)_$$(cat VERSION)_manifest.json"

	cd dist/publish && shasum -a 256 * > "terraform-provider-$(PROVIDER_NAME)_$$(cat ../../VERSION)_SHA256SUMS"
	gpg --detach-sign "terraform-provider-$(PROVIDER_NAME)_$$(cat VERSION)_SHA256SUMS"
	cd $(ROOT_DIR)

	gh release create \
		"v$$(cat VERSION)" \
		--title "v$$(cat VERSION)" \
		--notes-file .build/CHANGELOG.md \
		dist/publish/*


	# Caching timestamp
	date > release

.PHONY: ci-update
ci-update:
	@echo Make ci-update

	id

	git config --global user.name "Github Action"
	git config --global user.email "noreply@github.com"

	make --always-make data/vyos-1x-info.txt
	if [ -n "$$(git diff --stat "data/vyos-1x-info.txt")" ]; then
		git add "data/vyos-1x-info.txt"
		git commit -m "refactor: update to rolling release $$(cat data/vyos-1x-info.txt)"
	fi

	make generate
	make test
	make docs/index.md

	if [ -z "$$(git diff --stat)" ]; then
		echo "No changes to provider files were detected"
		exit 0
	fi

	git add -A
	git commit -m "ci: regenerate files"

	make version
	echo "Updated to rolling release $$(cat data/vyos-1x-info.txt)"

.PHONY: clean
clean:
	-@echo Make clean

	go clean -modcache
	go clean -testcache
	go clean -fuzzcache
	rm -rf dist
	rm -rf .build
	rm generate
	rm test
	rm build
	rm version
	rm release


################
# Helper targets
################

###
# Default helper target
# TODO improve Makefile help
#  should print only targets intended for human usage
#  should print preceding comment on those targets when spessific help is requested
#  should print all targets when requested
#  milestone: 3
help:
	echo "Most targets are not meant for manual usage, these are the private targets starting with dot"
	echo "Valid targets listed below"
	echo "<target>: <dependency> <dependency ...>" | egrep --color '^[^ ]*:'
	make -rpn | sed -n -e '/^$$/ { n ; /^[^ .#][^ ]*:/p ; }' | egrep --color '^[^ ]*:'

.PHONY: why
why:
	make -nd $(INPUT_ARGS) \
		| sed -rn 's/^ *([A-Za-z ]*Must remake target.*| Prerequisite.*is newer than.*)/\1/p' \
			| tac

.PHONY: diff
diff:
	git diff --word-diff --word-diff-regex='\w+' -- . ':!docs' ':!data/provider-schema' ':!*autogen*'
