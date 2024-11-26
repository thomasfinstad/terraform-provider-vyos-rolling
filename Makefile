
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
	@echo -e "\n\n###########################################################################"
	echo Make data/vyos-1x-info.txt

	mkdir -p data || true
	mkdir -p .build || true

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
	@echo -e "\n\n###########################################################################"
	echo Make .build/vyos-1x

	mkdir -p .build || true

	echo Clone repo if needed
	if [ ! -d ".build/vyos-1x" ]; then
		cd .build
		git clone "https://github.com/vyos/vyos-1x.git" vyos-1x;
		cd $(ROOT_DIR)
	fi

	echo Cache the commit sha
	cd .build/vyos-1x
	git rev-list --date=iso-strict -n 1 --before="$$(cat "../../data/vyos-1x-info.txt")" "origin/current" > ../vyos-1x.sha

	echo Checkout the commit
	git checkout current
	git pull
	git checkout "$$(cat ../vyos-1x.sha)"

###
# Autogenerate Schemas

# Convert from relaxng to XSD
.build/schema-definitions.xsd: data/vyos-1x-info.txt |.build/vyos-1x
	@echo -e "\n\n###########################################################################"
	echo Make .build/schema-definitions.xsd

	mkdir -p .build || true

	java -jar tools/trang-20091111/trang.jar -I rnc -O xsd .build/vyos-1x/schema/interface_definition.rnc .build/schema-definitions.xsd
	xmllint --format --recover --output '.build/schema-definitions.xsd' '.build/schema-definitions.xsd'

# Generate go structs from XSD
internal/vyos/schemadefinition/autogen-structs.go: data/vyos-1x-info.txt internal/vyos/schemadefinition/interface-definition.go .build/schema-definitions.xsd
	@echo -e "\n\n###########################################################################"
	echo Make internal/vyos/schemadefinition/autogen-structs.go

	[ ! -e internal/vyos/schemadefinition/autogen-structs.go ] || rm -v internal/vyos/schemadefinition/autogen-structs.go
	mkdir -p internal/vyos/schemadefinition || true
	mkdir .build || true

	# Generate structs from schema
	xgen -p schemadefinition -i .build/schema-definitions.xsd -o internal/vyos/schemadefinition/autogen-structs.go -l Go

	# TODO convert from sed mangling to go mangling
	#  create a go tool package and use dst in the same way
	#  as tools/generate-vyos-infterface-definition-structs/main.go
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
	@echo -e "\n\n###########################################################################"
	echo Make .build/interface-definitions

	[ ! -e ".build/interface-definitions" ] || rm -rf ".build/interface-definitions"
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

internal/vyos/vyosinterfaces/autogen.go: \
											data/vyos-1x-info.txt \
											tools/generate-vyos-infterface-definition-structs \
											$(shell find internal/vyos/schemadefinition -type f) \
											.build/interface-definitions \
											internal/vyos/schemadefinition/autogen-structs.go
	@echo -e "\n\n###########################################################################"
	echo Make internal/vyos/vyosinterfaces/autogen.go

	mkdir -p "internal/vyos/vyosinterfaces"

	[ ! -e internal/vyos/vyosinterfaces/autogen.go ] || rm -f internal/vyos/vyosinterfaces/autogen*.go

	# Generate interfaces
	cd tools/generate-vyos-infterface-definition-structs
	go get
	go run ./*.go \
			"../../.build/interface-definitions/" \
			"../../internal/vyos/vyosinterfaces" \
			"vyosinterface"

	# TODO clean up generation of vyosinterfaces/autogen.go
	#  either use HEREDOC to make it prettier or create a gotemplate
	#  milestone: 6
	cat <<- EOF > "../../internal/vyos/vyosinterfaces/autogen.go"
		// Autogenerated by Makefile. DO NOT EDIT.
		package vyosinterface

		import "$(GO_IMPORT_ROOT)/internal/vyos/schemadefinition"

		// GetInterfaces returns all autogenerated interface definitions
		func GetInterfaces() []schemadefinition.InterfaceDefinition {
		return []schemadefinition.InterfaceDefinition{
	EOF

	grep -r -o -h "func [a-z_]*()" ../../internal/vyos/vyosinterfaces/| sort | sed 's/func \(.*\)/\1,/g' \
		>> "../../internal/vyos/vyosinterfaces/autogen.go"

	echo -e "}\n}" >> "../../internal/vyos/vyosinterfaces/autogen.go"

	gofumpt -w ../../internal/vyos/vyosinterfaces/*.go

internal/terraform/resource/autogen/package.go: \
									data/vyos-1x-info.txt \
									internal/vyos/vyosinterfaces/autogen.go \
									$(shell find tools/generate-terraform-resource-full -type f)
	@echo -e "\n\n###########################################################################"
	echo Make internal/terraform/resource/autogen/package.go

	# Prep dirs
	[ ! -e "internal/terraform/resource/autogen/named" ] || rm -rf "internal/terraform/resource/autogen/named"
	mkdir -p "internal/terraform/resource/autogen/named"

	[ ! -e "internal/terraform/resource/autogen/global" ] || rm -rf "internal/terraform/resource/autogen/global"
	mkdir -p "internal/terraform/resource/autogen/global"

	# Generate
	cd tools/generate-terraform-resource-full
	go get
	go run ./main.go \
		"$(ROOT_DIR)" \
		"internal/terraform/resource/autogen" \
		"$(GO_IMPORT_ROOT)" \
		"system/conntrack-timeout-custom-rule"

	# Clean up code
	gofumpt -w "../../internal/terraform/resource/autogen/"
	goimports -w "../../internal/terraform/resource/autogen"

# Pretty-name target for human usage
generate: internal/terraform/resource/autogen/package.go
	@echo -e "\n\n###########################################################################"
	echo Make generate

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
test: generate \
		$(shell find . -not \( -path "*/.build/*" \) -not \( -path ./examples/provider/.tmp -prune \) -type f -name "*_test.go" ) \
		$(shell find . -not \( -path "*autogen*" -or -path "*/.build/*" -or -path "*/tools/*" \) -not \( -path ./examples/provider/.tmp -prune \) -type f -name "*.go" )

	@echo -e "\n\n###########################################################################"
	echo Make test

	# VyOS API can often take ~1 second to respond to a configure request.
	# This means we attempt to tune retrys and delays around this.
	# The end result is that to be able to test retry functionality we will need a bit of head room,
	# so 30s timeout should be enough for all tests combined.

	go test -count=1 -failfast -timeout 30s $(shell find . -not \( -path ./examples/provider/.tmp -prune \) -type f -name "*_test.go" -exec dirname {} \; | sort -u) $(INPUT_ARGS)

	# Caching timestamp
	date > test

build: test
	@echo -e "\n\n###########################################################################"
	echo Make build

	[ ! -e "${DIST_DIR}/providers.localhost" ] || rm -rf "${DIST_DIR}/providers.localhost"

	os=""
	case "$$OSTYPE" in
		linux-gnu)	os="linux" ;;
		darwin)		os="darwin" ;;
		freebsd)	os="freebsd" ;;
		*)			echo "Unsupported OSTYPE: '$$OSTYPE'"; exit 1 ;;
	esac

	arch=""
	case $$(uname -m) in
		i386 | i686)	arch="386" ;;
		x86_64)			arch="amd64" ;;
		arm)			arch="arm" ;;
		aarch64)		arch="arm64" ;;
		*)				echo "Unsupported system architecture: '$$(uname -m)'"; exit 1 ;;
	esac

	version=$(shell date +%Y%m%d.%H%M.%S);
	echo "Building for $${os} ($${arch})"
	build_dir="${DIST_DIR}/providers.localhost/dev/$(PROVIDER_NAME)/$${version}/$${os}_$${arch}"

	mkdir -p "$${build_dir}/";
	GOOS="$$os" \
	GOARCH="$$arch" \
		go build \
			-ldflags "-X main.version=$${version} -X main.address=${ADDRESS}" \
			-o "$${build_dir}/terraform-provider-$(PROVIDER_NAME)_v$${version}";

	# Caching timestamp
	date > build

build-all: test
	@echo -e "\n\n###########################################################################"
	echo Make build-all

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
			GOOS="$$os" \
			GOARCH="$$arch" \
				go build \
					-ldflags "-X main.version=$${version} -X main.address=${ADDRESS}" \
					-o "$${build_dir}/terraform-provider-$(PROVIDER_NAME)_v$${version}";
		done;
	done;

	# Caching timestamp
	date > build-all

docs/index.md: \
				build \
				internal/terraform/resource/autogen/package.go \
				$(shell find templates/ -type f)
	@echo -e "\n\n###########################################################################"
	echo Make docs/index.md

	# Prep dirs
	[ ! -e "docs/" ] || rm -rf "docs/"
	[ ! -e "./examples/resources/" ] || find ./examples/resources/ -name import.sh -execdir bash -c 'rm {}; rmdir --ignore-fail-on-non-empty -p $$PWD' \;

	# Generate import docs
	go run ./tools/generate-import-docs/ examples/resources/

	#
	echo Generate provider json schema
	[ ! -e examples/provider/.tmp/provider-schema.json ] || rm examples/provider/.tmp/provider-schema.json
	cd examples/provider
	make .tmp/provider-schema.json
	cd $(ROOT_DIR)

	# Create docs
	old_name="$$(cat examples/provider/.tmp/provider-schema.json | jq -r '.provider_schemas | keys | .[0]')"
	jq \
		--arg old_name "$$old_name" \
		'.provider_schemas.vyos=.provider_schemas[$$old_name] | del(.provider_schemas[$$old_name])' \
		examples/provider/.tmp/provider-schema.json \
		> examples/provider/.tmp/provider-schema-modified.json

	tfplugindocs generate \
		--provider-name "vyos" \
		--providers-schema "examples/provider/.tmp/provider-schema-modified.json" \
		--rendered-provider-name vyos \
		| grep -v '^rendering "resources/.*\.md\.tmpl"$$' \
		| grep -v 'resource ".*" fallback template exists, creating template'

	###
	echo Reverse html escaping for known parts that should stay as html code

	echo Fix comments
	sed -r -i 's/&lt;!-- (.*) --&gt;/<!-- \1 -->/' docs/resources/*.md

	echo Add whitespace to make the tables readable on docs website
	sed -r -i 's/^(\s*[|].*[^-])[|](.*[|])$$/\1\&emsp;|\2/' docs/resources/*.md

	echo Fix header links
	sed -r -i 's|&lt;a id=&#34;([A-Za-z0-9_-]+)&#34;&gt;&lt;/a&gt;|<a id="\1"></a>|' docs/resources/*.md

	echo Fix description/help linebreaks
	sed -r -i 's|&lt;/br&gt;|</br>|g' docs/resources/*.md

	echo Turn top level attributes into headers
	sed -r -i '/## Schema/,/### Nested Schema for/s/- `([A-Za-z0-9_]+)`/#### \1\n- `\1`/' docs/resources/*.md

	echo Enforce formatting of doc files
	(pre-commit run --files docs/* docs/resources/* || true) | grep -v '^Fixing docs/.*\.md$$'

.PHONY: changelog
changelog:
	@echo -e "\n\n###########################################################################"
	echo Make changelog

	echo Verify that there are no unstaged changes
	if [ -n "$$(git status -s | head)" ]; then
		git status | head -n 20
		echo "Can not create version when git is not in a clean, committed, and pushed state" >&2
		exit 1
	fi

	mkdir -p ".build" || true

	echo Generate provider json schema
	cd examples/provider
	[ ! -e "../../.build/new-provider-schema.json" ] || rm -rf "../../.build/new-provider-schema.json"
	make .tmp/provider-schema.json
	cp .tmp/provider-schema.json ../../.build/new-provider-schema.json
	cd $(ROOT_DIR)

	echo Last 5 releases:
	git tag -l | sort -V | tail -n5

	echo Move old changelog into .build dir
	mv CHANGELOG.md .build/old-CHANGELOG.md

	echo Copy out provider schema from last tagged release
	git show "$$(git tag | sort -V | tail -n1)":data/provider-schema.json > .build/old-provider-schema.json
	cd tools/generate-changelog
	echo Generate changelog
	go run *.go ../../.build/old-provider-schema.json ../../.build/new-provider-schema.json
	cd $(ROOT_DIR)

	echo Extract latest version from changelog
	prev_version="$$(cat VERSION)"
	grep -E -o "[0-9]+\.[0-9]+\.[0-9]+" .build/CHANGELOG.md | sort -V | tail -n1 > VERSION
	new_version="$$(cat VERSION)"

	echo "Old version: $$prev_version"
	echo "New version: $$new_version"

	echo Overwrite old json schema with new
	mv .build/new-provider-schema.json data/provider-schema.json

	echo Check if previous MAJOR version is different than the new/current one
	prev_major="$$(echo $$prev_version | cut -d. -f1)"
	new_major="$$(echo $$new_version | cut -d. -f1)"
	if [ "$$prev_major" != "$$new_major" ]; then
		echo Move old changelog into archive
		mkdir -p "data/changelogs"
		mv ".build/old-CHANGELOG.md" "data/changelogs/CHANGELOG-v$$prev_major.md"
		cat .build/CHANGELOG.md > CHANGELOG.md
		cat <<- EOF >> "CHANGELOG.md"
			## Previous changelogs

			For previous version see [changelog for v$$prev_major](data/changelogs/CHANGELOG-v$$prev_major.md) or older archives [directory](data/changelogs/)
		EOF

		sed -i "s|data/changelogs/CHANGELOG|CHANGELOG|" data/changelogs/*.md
	else
		echo Copy new changelog into place
		cat .build/CHANGELOG.md > CHANGELOG.md
		echo Append old changelog to end of new changelog
		echo "" >> CHANGELOG.md
		cat .build/old-CHANGELOG.md | sed -n '/^##/,$$p' >> CHANGELOG.md
	fi

.PHONY: prepare-git-for-release
prepare-git-for-release: changelog
	@echo -e "\n\n###########################################################################"
	echo Make prepare-git-for-release

	echo Stage files in git
	git add -v -A | sed 's|/.*$$|/|' | uniq -c | sort -n | column -t

	echo Run pre-commit and stage any changed files
	pre-commit run || git add -v -A | sed 's|/.*$$|/|' | uniq -c | sort -n | column -t

	echo Commit meta files for release to git
	git commit -m "chore: Prepare for release v$$(cat VERSION)"

	echo "Create git release tag v$$(cat VERSION)"
	git tag "v$$(cat VERSION)"

	cp .build/CHANGELOG.md CHANGELOG-FOR-GO-RELEASER.md.tmp

.PHONY: ci-update
ci-update:
	@echo -e "\n\n###########################################################################"
	echo Make ci-update

	echo "Previous release: $$(git tag -l v* | sort -V | tail -n1)"

	git config --global user.name "Github Action"
	git config --global user.email "noreply@github.com"

	make --always-make data/vyos-1x-info.txt
	if ! git diff --exit-code --stat "data/vyos-1x-info.txt"; then
		echo "Detected new rolling release: $$(cat data/vyos-1x-info.txt)"
		git add -v "data/vyos-1x-info.txt"
	fi

	echo Generate files
	make generate

	echo Run tests
	make test

	echo Render provider documentation
	make docs/index.md

	echo Check for changes since last release
	if git diff --ignore-matching-lines='[0-9-]+T[0-9:]+Z' --exit-code --stat  "$$(git tag -l v* | sort -V | tail -n1)" >/dev/null; then
		echo "No changes to provider files since last release"
		exit 0
	fi

	if [ -n "$$(git status -s)" ]; then
		echo "Changes detected:"
		git status --short | sed 's|/.*$$|/|' | uniq -c | sort -n | column -t

		echo Stage changed files to git:
		git add -v -A | sed 's|/.*$$|/|' | uniq -c | sort -n | column -t

		echo Commit files updated by CI
		git commit -m "ci: update to rolling release $$(cat data/vyos-1x-info.txt)"
	fi

	echo Generate changelog
	make changelog

	echo Add changes to git
	make prepare-git-for-release

	echo Run clean target to free up space before goreleaser runs
	make clean

	echo "Update to rolling release '$$(cat data/vyos-1x-info.txt)' complete, ready to release provider."

.PHONY: clean
clean:
	-@echo -e "\n\n###########################################################################"
	echo Make clean

	set +e

	go clean -modcache
	go clean -testcache
	go clean -fuzzcache
	rm -rf "${DIST_DIR}/*"
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
