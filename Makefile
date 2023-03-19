
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

# Fetch the ISO download webpage
.build/vyos/iso-download-page-hourly-$(TS_HOUR).html: .build/vyos
	mkdir -p .build/vyos

	find ".build/vyos" -name "iso-download-page-hourly-*.html" -delete;

	curl -s "https://vyos.net/get/nightly-builds/" > .build/vyos/iso-download-page-hourly-$(TS_HOUR).html

	tidy_ret="$$(tidy -indent -ashtml -quiet -modify .build/vyos/iso-download-page-hourly-$(TS_HOUR).html 1>&2; echo $$?)"; \
	if [ "$$tidy_ret" -eq 1 ]; then \
		echo "Tidy had warnings"; \
	elif [ "$$tidy_ret" -eq 2 ]; then \
		echo "Tidy had errors"; \
		exit 1; \
	else \
		echo "Tidy unkown exit code: '$$tidy_ret'"; \
		exit 1; \
	fi

# Extract ISO urls
.build/vyos/iso-url.txt: .build/vyos/iso-download-page-hourly-$(TS_HOUR).html
	xmllint \
		--html \
		--nowarning \
		--xpath "/html/body/main/div[@id='content']/div[@id='rolling-current']/ul/*/a/@href" \
		.build/vyos/iso-download-page-hourly-$(TS_HOUR).html 2>/dev/null \
		| cut -d'"' -f2 \
		| grep -v "vyos-rolling-latest.iso" \
		| sort \
		| tail -n1 \
		| tee ".build/vyos/iso-url.txt";

# Fetch and format newest ISO modification timestamp
data/vyos/rolling-iso-time.txt: .build/vyos/iso-url.txt
	curl -s --location --head \
		"$(shell cat .build/vyos/iso-url.txt)" \
		> ".build/vyos/iso-headers.txt"

	mkdir -p data/vyos
	date \
		-d "$(shell sed -n '/last-modified:/s/last-modified: //p' ".build/vyos/iso-headers.txt")" \
		-Iseconds \
		> "data/vyos/rolling-iso-time.txt"

###
# VyOS src repo
data/vyos/vyos-1x: data/vyos/rolling-iso-time.txt
	git submodule update --init --single-branch -- data/vyos/vyos-1x

	cd data/vyos/vyos-1x && \
	git checkout "$$(git rev-list --date=iso-strict -n 1 --before="$(shell cat "data/vyos/rolling-iso-time.txt")" current)"

###
# Autogenerate Schemas

# Convert from relaxng to XSD
.build/vyos/schema/interface-definition.xsd: .build/vyos/vyos-1x
	mkdir -p .build/vyos/schema/
	java -jar scripts/trang-20091111/trang.jar -I rnc -O xsd data/vyos/vyos-1x/schema/interface_definition.rnc .build/vyos/schema/interface-definition.xsd

# Generate go structs from XSD
vyos_api/schema/autogen-interface-definition.go: .build/vyos/schema/interface-definition.xsd
	go generate vyos_api/schema/interface-definition.go

###
# Dump definitions to a more human readable format
# TODO next
# Interface definitions
.build/vyos/vyos-1x/venv: data/vyos/vyos-1x
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
