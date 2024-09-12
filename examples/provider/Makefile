DIST_DIR:=$(shell sed -n -e '/^DIST_DIR=/s/.*=//p' ../../Makefile)
PROVIDER_DIR:=$(DIST_DIR)/$(shell sed -n -e '/^\s*source\s*=/s/.*=\s"\(.*\)"/\1/p' local_override.tf)

default: plan

$(PROVIDER_DIR):
	cd ../.. && make build

init: $(PROVIDER_DIR) $(shell find $(PROVIDER_DIR) -type f) $(shell find . -type f -name "*.tf" -execdir grep -l provider "{}" \; | sort -u)
	-rm -rf .terraform .terraform.lock.hcl
	TF_LOG_PROVIDER=TRACE TF_LOG_PATH=trace.log tofu init -upgrade

	# Cache timestamp
	date > init

# Ref: "Optional dynamic parameters" section below
# Example usage: make plan -- -target vyos_banana.nomnom
ifeq (plan,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for target
  DYN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(DYN_ARGS):;@:)
endif
plan: init $(shell find . -type f -name "*.tf")
	-rm -fv trace.log
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
	-rm -fv trace.log
	TF_LOG_PROVIDER=TRACE TF_LOG_PATH=trace.log tofu apply -auto-approve $(DYN_ARGS) plan

	# Cache timestamp
	date > apply

.PHONY: clean
clean:
	-rm init plan apply

	-rm -rf .terraform .terraform.lock.hcl
	-rm -fv trace.log

.PHONY: why
why:
	@make -and apply \
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
	@echo Make .tmp/provider-schema.json
	tofu providers schema -json | jq '.' > .tmp/provider-schema.json
