// Package namedfirewallgroupdynamicgroupipvsixaddressgroup code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallgroupdynamicgroupipvsixaddressgroup

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &firewallGroupDynamicGroupIPvsixAddressGroup{}
	_ resource.ResourceWithConfigure = &firewallGroupDynamicGroupIPvsixAddressGroup{}
)

// var _ resource.ResourceWithConfigValidators = &firewallGroupDynamicGroupIPvsixAddressGroup{}
// var _ resource.ResourceWithModifyPlan = &firewallGroupDynamicGroupIPvsixAddressGroup{}
// var _ resource.ResourceWithUpgradeState = &firewallGroupDynamicGroupIPvsixAddressGroup{}
// var _ resource.ResourceWithValidateConfig = &firewallGroupDynamicGroupIPvsixAddressGroup{}
// var _ resource.ResourceWithImportState = &firewallGroupDynamicGroupIPvsixAddressGroup{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallGroupDynamicGroupIPvsixAddressGroup{}
