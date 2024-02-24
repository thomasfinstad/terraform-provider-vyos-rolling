
default: plan

.PHONY: init
init: provider.tf ../.terraform-version ../../dist/
	-rm -rf .terraform .terraform.lock.hcl
	terraform init -upgrade

.PHONY: plan
plan: init
	terraform plan

.PHONY: apply
apply: init
	terraform apply -auto-approve

.PHONY: trace
trace: init
	-rm -fv trace.log
	TF_LOG_PROVIDER=TRACE TF_LOG_PATH=trace.log terraform apply -auto-approve
