// Package globalfirewallglobaloptionsstatepolicyestablished code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalfirewallglobaloptionsstatepolicyestablished

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &firewallGlobalOptionsStatePolicyEstablished{}
	_ resource.ResourceWithConfigure = &firewallGlobalOptionsStatePolicyEstablished{}
)

// var _ resource.ResourceWithConfigValidators = &firewallGlobalOptionsStatePolicyEstablished{}
// var _ resource.ResourceWithModifyPlan = &firewallGlobalOptionsStatePolicyEstablished{}
// var _ resource.ResourceWithUpgradeState = &firewallGlobalOptionsStatePolicyEstablished{}
// var _ resource.ResourceWithValidateConfig = &firewallGlobalOptionsStatePolicyEstablished{}
// var _ resource.ResourceWithImportState = &firewallGlobalOptionsStatePolicyEstablished{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallGlobalOptionsStatePolicyEstablished{}
