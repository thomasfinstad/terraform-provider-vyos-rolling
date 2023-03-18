
TS_HOUR := $(shell date -u +%Y-%m-%d--%H)

default: data/vyos/vyos-1x

.build/vyos:
	mkdir -p .build/vyos

.build/vyos/iso-download-page-hourly-$(TS_HOUR).html: .build/vyos
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


.build/vyos/iso-time.txt: .build/vyos/iso-url.txt
	curl -s --location --head \
		"$(shell cat .build/vyos/iso-url.txt)" \
		> ".build/vyos/iso-headers.txt"

	date \
		-d "$(shell sed -n '/last-modified:/s/last-modified: //p' ".build/vyos/iso-headers.txt")" \
		-Iseconds \
		> .build/vyos/iso-time.txt

data/vyos:
	mkdir -p data/vyos

data/vyos/rolling-iso-time.txt: data/vyos .build/vyos/iso-time.txt
	cp ".build/vyos/iso-time.txt" "data/vyos/rolling-iso-time.txt"

data/vyos/vyos-1x: data/vyos/rolling-iso-time.txt
	git submodule update --init --single-branch -- data/vyos/vyos-1x

	cd data/vyos/vyos-1x && \
	git checkout "$$(git rev-list --date=iso-strict -n 1 --before="$(shell cat "data/vyos/rolling-iso-time.txt")" current)"
