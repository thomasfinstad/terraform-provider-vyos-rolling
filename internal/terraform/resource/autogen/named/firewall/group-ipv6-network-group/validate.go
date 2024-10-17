// Package namedfirewallgroupipvsixnetworkgroup code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallgroupipvsixnetworkgroup

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &firewallGroupIPvsixNetworkGroup{}
	_ resource.ResourceWithConfigure   = &firewallGroupIPvsixNetworkGroup{}
	_ resource.ResourceWithImportState = &firewallGroupIPvsixNetworkGroup{}
)

// var _ resource.ResourceWithConfigValidators = &firewallGroupIPvsixNetworkGroup{}
// var _ resource.ResourceWithModifyPlan = &firewallGroupIPvsixNetworkGroup{}
// var _ resource.ResourceWithUpgradeState = &firewallGroupIPvsixNetworkGroup{}
// var _ resource.ResourceWithValidateConfig = &firewallGroupIPvsixNetworkGroup{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallGroupIPvsixNetworkGroup{}
