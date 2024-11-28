SHELL := bash
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
	@echo -e "\n\n# examples/provider #######################################################"
	echo Make $(PROVIDER_DIR)

	cd ../..
	rm build || true
	make build

init: \
	$(PROVIDER_DIR) \
	$(shell find ${PROVIDER_DIR} -type f) \
	$(shell find . -not \( -path ./.tmp -prune \) -type f -name "*.tf" -execdir grep -l provider "{}" \; | sort -u)

	@echo -e "\n\n# examples/provider #######################################################"
	echo Make init

	rm -rf .terraform .terraform.lock.hcl || true
	TF_LOG_PROVIDER=TRACE TF_LOG_PATH=trace.log tofu init -upgrade

	# Cache timestamp
	date > init

validate: init $(shell find . -not \( -path ./.tmp -prune \) -type f -name "*.tf")
	@echo -e "\n\n# examples/provider #######################################################"
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
plan: validate $(shell find . -not \( -path ./.tmp -prune \) -type f -name "*.tf") .tmp/container
	@echo -e "\n\n# examples/provider #######################################################"
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
	@echo -e "\n\n# examples/provider #######################################################"
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
	@echo -e "\n\n# examples/provider #######################################################"
	echo Make destroy

	rm -fv trace.log || true
	TF_LOG_PROVIDER=TRACE TF_LOG_PATH=trace.log tofu apply -destroy -auto-approve $(DYN_ARGS) || true

	rm plan || true
	rm apply || true

.PHONY: clean
clean: destroy
	@echo -e "\n\n# examples/provider #######################################################"
	echo Make clean

	set +e

	rm init plan apply validate

	rm -rf .terraform .terraform.lock.hcl || true
	rm -fv trace.log || true
	# keeping the iso file to reduce redundant downloads
	sudo rm -rf .tmp/filesystem.squashfs || true
	sudo docker stop vyos-container || true
	sudo rm -f .tmp/container || true
	sleep 2
	sudo docker rmi -f $$(sudo docker images -a -q vyos-local) || true
	sudo rm -f .tmp/container-img || true

.PHONY: test
test:
	@echo -e "\n\n# examples/provider #######################################################"
	echo Make test

	echo ""
	echo "Initial apply------------------------------------------------------"
	make apply || (echo "View ./trace.log" >&2; exit 1)
	echo "Secondary apply------------------------------------------------------"
	make apply || (echo "View ./trace.log" >&2; exit 1)

.PHONY: why
why:
	@echo -e "\n\n# examples/provider #######################################################"
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
	@echo -e "\n\n# examples/provider #######################################################"
	echo Make .tmp/provider-schema.json
	mkdir -p .tmp/
	tofu providers schema -json | jq '.' > .tmp/provider-schema.json

.tmp/vyos.iso: ../../data/vyos-1x-info.txt
	@echo -e "\n\n# examples/provider #######################################################"
	echo Make .tmp/vyos.iso

	build_date="$$(sed -n 's/T.*$$//p' "../../data/vyos-1x-info.txt")"

	mkdir -p .tmp/
	download_url="$$( \
		curl -L \
			-H "Accept: application/vnd.github+json" \
			-H "X-GitHub-Api-Version: 2022-11-28" \
			"https://api.github.com/repos/vyos/vyos-nightly-build/releases?per_page=30" \
				| jq -r --arg var_date "$${build_date}" '.[] | select(.published_at|startswith($$var_date)) | .assets[] | select(.name|endswith(".iso")) | .browser_download_url' \
				| head -n1)"

	if [ -z "$$download_url" ]; then
		download_url="$$( \
			curl -L \
				-H "Accept: application/vnd.github+json" \
				-H "X-GitHub-Api-Version: 2022-11-28" \
				"https://api.github.com/repos/vyos/vyos-nightly-build/releases?per_page=30" \
					| jq -r '.[] | .assets[] | select(.name|endswith(".iso")) | .browser_download_url' \
					| head -n1)"
	fi

	wget -O .tmp/vyos.iso "$${download_url}"
	touch .tmp/vyos.iso

.tmp/filesystem.squashfs: .tmp/vyos.iso
	@echo -e "\n\n# examples/provider #######################################################"
	echo Make .tmp/filesystem.squashfs
	mkdir -p .tmp/
	build_date="$$(sed -n 's/T.*$$//p' "../../data/vyos-1x-info.txt")"

	cd .tmp
	rm -f filesystem.squashfs
	xorriso -osirrox on -indev vyos.iso -extract /live/filesystem.squashfs filesystem.squashfs
	touch filesystem.squashfs

.tmp/container-img: .tmp/filesystem.squashfs
	@echo -e "\n\n# examples/provider #######################################################"
	echo Make .tmp/vyos-unsquashfs
	mkdir -p .tmp/

	build_date="$$(sed -n 's/T.*$$//p' "../../data/vyos-1x-info.txt")"

	sudo rm -rf .tmp/vyos-unsquashfs || true
	mkdir -p .tmp/vyos-unsquashfs
	cd .tmp

	sudo unsquashfs -f -d "vyos-unsquashfs" -excludes filesystem.squashfs /dev
	sudo touch vyos-unsquashfs
	sudo tar -C vyos-unsquashfs -c . | sudo docker import - "vyos-local:$${build_date}"
	sudo docker tag "vyos-local:$${build_date}" vyos-local:latest
	date > container-img
	sudo rm -rf vyos-unsquashfs


.tmp/container: .tmp/container-img assets/container-setup-api.vbash
	@echo -e "\n\n# examples/provider #######################################################"
	echo Make .tmp/container
	mkdir -p .tmp/

	build_date="$$(sed -n 's/T.*$$//p' "../../data/vyos-1x-info.txt")"

	until [ -z "$$(sudo docker ps -f 'name=^vyos-container$$' -q)" ]; do
		sudo docker stop vyos-container
		sleep 5
	done

	sudo docker rm vyos-container || true

	sudo docker run --rm --name vyos-container --privileged -p 8443:443 --hostname localhost -t -d "vyos-local:$${build_date}" /sbin/init

	date > .tmp/container

	for i in {1..999}; do
		echo "Waiting for container config service to come up..."
		if sudo docker exec vyos-container systemctl is-active vyos-configd.service; then
			break
		elif [ $$i -eq 25 ]; then
			echo "container config service not started" >&2
			exit 1
		else
			sleep 5
		fi
	done

	bash -c 'while [ -n "$$(sudo docker ps -f "name=^vyos-container$$" -q)" ]; do sleep 5; done; rm .tmp/container' &

	sudo docker cp assets/container-setup-api.vbash vyos-container:/container-setup-api.vbash
	sudo docker exec -t vyos-container chmod +x /container-setup-api.vbash
	sudo docker exec -t vyos-container /container-setup-api.vbash

	for i in {1..999}; do
		echo "Waiting for container api service to come up..."
		if sudo docker exec vyos-container systemctl is-active vyos-http-api.service; then
			break
		elif [ $$i -eq 25 ]; then
			echo "container api service not started" >&2
			exit 1
		else
			sleep 5
		fi
	done

	vyos_ip="$$(sudo docker inspect --format json --type container vyos-container | jq -r '.[].NetworkSettings.IPAddress')"

	echo "Adding vyos container to hosts -> $${vyos_ip}  vyos.local"
	echo -ne "X\n$${vyos_ip}  vyos.local" | sudo tee -a /etc/hosts >/dev/null
