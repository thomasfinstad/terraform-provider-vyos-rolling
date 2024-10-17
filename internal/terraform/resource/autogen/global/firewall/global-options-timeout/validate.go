// Package globalfirewallglobaloptionstimeout code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalfirewallglobaloptionstimeout

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &firewallGlobalOptionsTimeout{}
	_ resource.ResourceWithConfigure   = &firewallGlobalOptionsTimeout{}
	_ resource.ResourceWithImportState = &firewallGlobalOptionsTimeout{}
)

// var _ resource.ResourceWithConfigValidators = &firewallGlobalOptionsTimeout{}
// var _ resource.ResourceWithModifyPlan = &firewallGlobalOptionsTimeout{}
// var _ resource.ResourceWithUpgradeState = &firewallGlobalOptionsTimeout{}
// var _ resource.ResourceWithValidateConfig = &firewallGlobalOptionsTimeout{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallGlobalOptionsTimeout{}
