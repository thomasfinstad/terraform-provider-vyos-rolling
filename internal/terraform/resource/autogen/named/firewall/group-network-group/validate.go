// Package namedfirewallgroupnetworkgroup code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallgroupnetworkgroup

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &firewallGroupNetworkGroup{}
	_ resource.ResourceWithConfigure = &firewallGroupNetworkGroup{}
)

// var _ resource.ResourceWithConfigValidators = &firewallGroupNetworkGroup{}
// var _ resource.ResourceWithModifyPlan = &firewallGroupNetworkGroup{}
// var _ resource.ResourceWithUpgradeState = &firewallGroupNetworkGroup{}
// var _ resource.ResourceWithValidateConfig = &firewallGroupNetworkGroup{}
// var _ resource.ResourceWithImportState = &firewallGroupNetworkGroup{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallGroupNetworkGroup{}
