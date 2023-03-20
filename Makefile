
TS_HOUR := $(shell date -u +%Y-%m-%d--%H)

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
	go generate internal/provider/vyos/schema/interfacedefinition/interface-definition.go

###
# Dump definitions to a more human readable format
# TODO next
# Interface definitions
.build/vyos/vyos-1x/venv: data/vyos/vyos-1x/submodule.log
	python -m venv .build/vyos/vyos-1x/venv

	bash -c " \
		source .build/vyos/vyos-1x/venv/bin/activate && \
		cd data/vyos/vyos-1x && \
		pip install -r test-requirements.txt \
	"

.build/vyos/interface-definitions: .build/vyos/vyos-1x/venv
	bash -c " \
		source .build/vyos/vyos-1x/venv/bin/activate && \
		cd data/vyos/vyos-1x && \
		make interface_definitions \
	"

	mv data/vyos/vyos-1x/interface-definitions .build/vyos/interface-definitions

clean:
	rm -rfv .build
	git submodule deinit --all
