// Package namedfirewallgroupdynamicgroupaddressgroup code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallgroupdynamicgroupaddressgroup

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &firewallGroupDynamicGroupAddressGroup{}
	_ resource.ResourceWithConfigure   = &firewallGroupDynamicGroupAddressGroup{}
	_ resource.ResourceWithImportState = &firewallGroupDynamicGroupAddressGroup{}
)

// var _ resource.ResourceWithConfigValidators = &firewallGroupDynamicGroupAddressGroup{}
// var _ resource.ResourceWithModifyPlan = &firewallGroupDynamicGroupAddressGroup{}
// var _ resource.ResourceWithUpgradeState = &firewallGroupDynamicGroupAddressGroup{}
// var _ resource.ResourceWithValidateConfig = &firewallGroupDynamicGroupAddressGroup{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallGroupDynamicGroupAddressGroup{}
