
SHELL:=/usr/bin/bash

###
# Default helper target
help:
	@echo "Most targets are not ment for manual usage, these are the private targets starting with dot"
	@echo "Valid targets listed below"
	@echo "<target>: <dependency> <dependency ...>" | egrep --color '^[^ ]*:'
	@make -rpn | sed -n -e '/^$$/ { n ; /^[^ .#][^ ]*:/p ; }' | egrep --color '^[^ ]*:'

###
# VyOS ISO Upload timestamp

# Fetch and format newest ISO modification timestamp
data/vyos/rolling-iso-time.txt:
	curl -s --location --head \
		"https://s3-us.vyos.io/rolling/current/vyos-rolling-latest.iso" \
		> ".build/vyos/iso-headers.txt"

	mkdir -p data/vyos
	date \
		-d "$(shell sed -n '/last-modified:/s/last-modified: //p' ".build/vyos/iso-headers.txt")" \
		-Iseconds \
		> "data/vyos/rolling-iso-time.txt"

###
# VyOS src repo at correct commit
data/vyos/vyos-1x/submodule.log: data/vyos/rolling-iso-time.txt
	git submodule update --init --single-branch -- data/vyos/vyos-1x

	cd data/vyos/vyos-1x && \
	commit="$$(git rev-list --date=iso-strict -n 1 --before="$(shell cat "data/vyos/rolling-iso-time.txt")" "current")" && \
	git checkout "$$commit" && \
	echo "$$commit" > submodule.log

###
# Autogenerate Schemas

# Convert from relaxng to XSD
.build/vyos/schema/interface-definition.xsd: data/vyos/vyos-1x/submodule.log
	mkdir -p .build/vyos/schema/
	java -jar scripts/trang-20091111/trang.jar -I rnc -O xsd data/vyos/vyos-1x/schema/interface_definition.rnc .build/vyos/schema/interface-definition.xsd

# Generate go structs from XSD
internal/provider/vyos/schema/interfacedefinition/autogen-structs.go: .build/vyos/schema/interface-definition.xsd internal/provider/vyos/schema/interfacedefinition/interface-definition.go
	go run github.com/xuri/xgen/cmd/xgen -p interfacedefinition -i .build/vyos/schema/interface-definition.xsd -o internal/provider/vyos/schema/interfacedefinition/autogen-structs.go -l Go
	sed -i 's|interface{}|string|' internal/provider/vyos/schema/interfacedefinition/autogen-structs.go


###
# Terraform Resource Schemas

# Compile interface devfinitions
.build/vyos/interface-definitions: data/vyos/vyos-1x/submodule.log

	@rm -rf ".build/vyos/interface-definitions"
	python -m venv .build/vyos/vyos-1x/venv

	bash -c " \
		source .build/vyos/vyos-1x/venv/bin/activate && \
		cd data/vyos/vyos-1x && \
		pip install -r test-requirements.txt && \
		make interface_definitions \
	"

	mv data/vyos/vyos-1x/build/interface-definitions .build/vyos/interface-definitions

internal/provider/vyos/interface/package.go: .build/vyos/interface-definitions
	mkdir -p "internal/provider/vyos/interface"

	for xmlFile in $(shell ls ".build/vyos/interface-definitions/"); do \
		echo -en "Input xml: '$${xmlFile}'\t"; \
		go run tools/build-resource-terraform-schemas/main.go ".build/vyos/interface-definitions/$${xmlFile}" "internal/provider/vyos/interface" "vyosinterface"; \
	done

	echo -e "// Package vyosinterface generated by Makefile on $(shell date -u -Iseconds). DO NOT EDIT.\npackage vyosinterface\n " > "internal/provider/vyos/interface/package.go"

	gofumpt -l -w "internal/provider/vyos/interface"
	gofmt -l -w -s "./internal/provider/vyos/interface/"

clean:
	rm -rfv .build
	git submodule deinit --all
