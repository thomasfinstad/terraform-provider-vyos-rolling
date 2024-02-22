
default: plan

.PHONY: init
init: provider.tf .terraform.lock.hcl .terraform ../.terraform-version ~/.terraform.d/plugins/
	rm -rf .terraform .terraform.lock.hcl
	terraform init -upgrade

.PHONY: trace
trace: init
	-rm -fv trace.log
	TF_LOG_PROVIDER=TRACE TF_LOG_PATH=trace.log terraform apply -auto-approve
