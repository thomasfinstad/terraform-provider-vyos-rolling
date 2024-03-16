
default: plan

.PHONY: init
init: provider.tf ../.terraform-version ../../dist/
	-rm -rf .terraform .terraform.lock.hcl
	-rm -fv trace.log
	TF_LOG_PROVIDER=TRACE TF_LOG_PATH=trace.log terraform init -upgrade

.PHONY: plan
plan: init
	-rm -fv trace.log
	TF_LOG_PROVIDER=TRACE TF_LOG_PATH=trace.log terraform plan

.PHONY: apply
apply: init
	-rm -fv trace.log
	TF_LOG_PROVIDER=TRACE TF_LOG_PATH=trace.log terraform apply -auto-approve
