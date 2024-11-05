SHELL := bash
# .SHELLFLAGS := -eu -x -o pipefail -c
.SHELLFLAGS := -eu -o pipefail -c
.ONESHELL:
.DELETE_ON_ERROR:
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules
ROOT_DIR := $(CURDIR)

DIST_DIR:=$(shell sed -n -e '/^DIST_DIR=/s/.*=//p' ../../Makefile | tr -d '"')
PROVIDER_DIR:=$(DIST_DIR)/$(shell sed -n -e '/^\s*source\s*=/s/.*=\s"\(.*\)"/\1/p' local_override.tf)

$(info DIST_DIR: ${DIST_DIR})
$(info PROVIDER_DIR: ${PROVIDER_DIR})

default: plan

$(PROVIDER_DIR):
	@echo -e "\n\n###########################################################################"
	echo Make $(PROVIDER_DIR)

	cd ../..
	rm build || true
	make build

init: $(PROVIDER_DIR) $(shell find ${PROVIDER_DIR} -type f) $(shell find . -type f -name "*.tf" -execdir grep -l provider "{}" \; | sort -u)
	@echo -e "\n\n###########################################################################"
	echo Make init

	rm -rf .terraform .terraform.lock.hcl || true
	TF_LOG_PROVIDER=TRACE TF_LOG_PATH=trace.log tofu init -upgrade

	# Cache timestamp
	date > init

validate: init $(shell find . -type f -name "*.tf")
	@echo -e "\n\n###########################################################################"
	echo Make validate

	TF_LOG_PROVIDER=TRACE TF_LOG_PATH=trace.log tofu validate

	# Cache timestamp
	date > validate

# Ref: "Optional dynamic parameters" section below
# Example usage: make plan -- -target vyos_banana.nomnom
ifeq (plan,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for target
  DYN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(DYN_ARGS):;@:)
endif
plan: validate $(shell find . -type f -name "*.tf") .tmp/container
	@echo -e "\n\n###########################################################################"
	echo Make plan

	rm -fv trace.log || true
	TF_LOG_PROVIDER=TRACE TF_LOG_PATH=trace.log tofu plan -out=plan $(DYN_ARGS)

# Ref: "Optional dynamic parameters" section below
# Example usage: make apply -- -target vyos_banana.nomnom
ifeq (apply,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for target
  DYN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(DYN_ARGS):;@:)
endif
apply: plan
	@echo -e "\n\n###########################################################################"
	echo Make apply

	rm -fv trace.log || true
	TF_LOG_PROVIDER=TRACE TF_LOG_PATH=trace.log tofu apply -auto-approve $(DYN_ARGS) plan

	rm plan

	# Cache timestamp
	date > apply

# Ref: "Optional dynamic parameters" section below
# Example usage: make destroy -- -target vyos_banana.nomnom
ifeq (destroy,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for target
  DYN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(DYN_ARGS):;@:)
endif
destroy: init
	@echo -e "\n\n###########################################################################"
	echo Make destroy

	rm -fv trace.log || true
	TF_LOG_PROVIDER=TRACE TF_LOG_PATH=trace.log tofu apply -destroy -auto-approve $(DYN_ARGS) || true

	rm plan || true
	rm apply || true

.PHONY: clean
clean: destroy
	@echo -e "\n\n###########################################################################"
	echo Make clean

	set +e

	rm init plan apply validate

	rm -rf .terraform .terraform.lock.hcl
	rm -fv trace.log
	# keeping the iso file to reduce redundant downloads
	sudo rm -rf .tmp/filesystem.squashfs
	docker stop vyos
	sudo rm -f .tmp/container
	sleep 2
	docker rmi -f $$(docker images -a -q vyos)
	sudo rm -f .tmp/container-img

.PHONY: test
test:
	@echo -e "\n\n###########################################################################"
	echo Make test

	echo ""
	echo "Initial apply------------------------------------------------------"
	make apply || (echo "View ./trace.log" >&2; exit 1)
	echo "Secondary apply------------------------------------------------------"
	make apply || (echo "View ./trace.log" >&2; exit 1)

.PHONY: why
why:
	@echo -e "\n\n###########################################################################"
	echo Make why

	make -nd apply \
		| sed -rn 's/^ *([A-Za-z ]*Must remake target.*| Prerequisite.*is newer than.*)/\1/p' \
			| tac

###
# Optional dynamic parameters
#
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

.tmp/provider-schema.json: init
	@echo -e "\n\n###########################################################################"
	echo Make .tmp/provider-schema.json
	mkdir -p .tmp/
	tofu providers schema -json | jq '.' > .tmp/provider-schema.json

.tmp/vyos.iso: ../../data/vyos-1x-info.txt
	@echo -e "\n\n###########################################################################"
	echo Make .tmp/vyos.iso

	set -x

	build_date="$$(sed -n 's/T.*$$//p' "../../data/vyos-1x-info.txt")"

	mkdir -p .tmp/
	download_url="$$( \
		curl -L \
			-H "Accept: application/vnd.github+json" \
			-H "X-GitHub-Api-Version: 2022-11-28" \
			"https://api.github.com/repos/vyos/vyos-nightly-build/releases?per_page=100" \
			| jq -r --arg var_date "$${build_date}" '.[] | select(.published_at|startswith($$var_date)) | .assets[] | select(.name|endswith(".iso")) | .browser_download_url' | head -n1)"

	wget -O .tmp/vyos.iso "$${download_url}"
	touch .tmp/vyos.iso

.tmp/filesystem.squashfs: .tmp/vyos.iso
	@echo -e "\n\n###########################################################################"
	echo Make .tmp/filesystem.squashfs
	mkdir -p .tmp/
	build_date="$$(sed -n 's/T.*$$//p' "../../data/vyos-1x-info.txt")"

	cd .tmp
	rm -f filesystem.squashfs
	xorriso -osirrox on -indev vyos.iso -extract /live/filesystem.squashfs filesystem.squashfs
	touch filesystem.squashfs

.tmp/container-img: .tmp/filesystem.squashfs
	@echo -e "\n\n###########################################################################"
	echo Make .tmp/vyos-unsquashfs
	mkdir -p .tmp/

	build_date="$$(sed -n 's/T.*$$//p' "../../data/vyos-1x-info.txt")"

	sudo rm -rf .tmp/vyos-unsquashfs || true
	mkdir -p .tmp/vyos-unsquashfs
	cd .tmp

	sudo unsquashfs -f -d "vyos-unsquashfs" -excludes filesystem.squashfs /dev
	sudo touch vyos-unsquashfs
	sudo tar -C vyos-unsquashfs -c . | docker import - "vyos:$${build_date}"
	docker tag "vyos:$${build_date}" vyos:latest
	touch container-img
	sudo rm -rf vyos-unsquashfs


.tmp/container: .tmp/container-img assets/container-setup-api.vbash
	@echo -e "\n\n###########################################################################"
	echo Make .tmp/container
	mkdir -p .tmp/

	build_date="$$(sed -n 's/T.*$$//p' "../../data/vyos-1x-info.txt")"

	set -x

	until [ -z "$$(docker ps -f "name=vyos" -q)" ]; do
		docker stop vyos
		sleep 5
	done

	bash -c "touch .tmp/container; docker run --rm --name vyos --privileged --network=bridge -p 8443:443 "vyos:latest" /sbin/init; rm -f .tmp/container" &
	sleep 5

	s=5; for i in {1..5}; do [ -n "$$(docker ps -f "name=vyos" -q)" ] && s=0 && break || s=$$? && sleep 5; done; [ $$s -eq 0 ]

	set +x

	docker cp assets/container-setup-api.vbash vyos:/container-setup-api.vbash
	docker exec -t vyos chmod +x /container-setup-api.vbash
	docker exec -t vyos /container-setup-api.vbash

	#vyos_ip="$$(docker exec -t vyos ip -4 --color=never --brief address show dev eth0 | sed -r 's|^eth0[^0-9]+([.0-9]+)/.*$$|\1|')"
	vyos_ip="$$(docker inspect --format json --type container vyos | jq -r '.[].NetworkSettings.IPAddress')"

	echo -ne "X\n$${vyos_ip}  vyos.local" | sudo tee -a /etc/hosts
