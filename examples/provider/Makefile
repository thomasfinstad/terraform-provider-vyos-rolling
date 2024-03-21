
default: plan

.PHONY: init
init: provider.tf ../.terraform-version ../../dist/
	-rm -rf .terraform .terraform.lock.hcl
	-rm -fv trace.log
	TF_LOG_PROVIDER=TRACE TF_LOG_PATH=trace.log terraform init -upgrade

# Ref: "Optional dynamic parameters" section below
# Example usage: make plan -- -target vyos_banana.nomnom
ifeq (plan,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for target
  DYN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(DYN_ARGS):;@:)
endif
.PHONY: plan
plan: init
	-rm -fv trace.log
	TF_LOG_PROVIDER=TRACE TF_LOG_PATH=trace.log terraform plan $(DYN_ARGS)

# Ref: "Optional dynamic parameters" section below
# Example usage: make apply -- -target vyos_banana.nomnom
ifeq (apply,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for target
  DYN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(DYN_ARGS):;@:)
endif
.PHONY: apply
apply: init
	-rm -fv trace.log
	TF_LOG_PROVIDER=TRACE TF_LOG_PATH=trace.log terraform apply -auto-approve $(DYN_ARGS)


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
