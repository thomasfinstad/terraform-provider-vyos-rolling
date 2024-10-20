/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedfirewallgroupdynamicgroupipvsixaddressgroup code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallgroupdynamicgroupipvsixaddressgroup

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &firewallGroupDynamicGroupIPvsixAddressGroup{}
	_ resource.ResourceWithConfigure   = &firewallGroupDynamicGroupIPvsixAddressGroup{}
	_ resource.ResourceWithImportState = &firewallGroupDynamicGroupIPvsixAddressGroup{}
)

// var _ resource.ResourceWithConfigValidators = &firewallGroupDynamicGroupIPvsixAddressGroup{}
// var _ resource.ResourceWithModifyPlan = &firewallGroupDynamicGroupIPvsixAddressGroup{}
// var _ resource.ResourceWithUpgradeState = &firewallGroupDynamicGroupIPvsixAddressGroup{}
// var _ resource.ResourceWithValidateConfig = &firewallGroupDynamicGroupIPvsixAddressGroup{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallGroupDynamicGroupIPvsixAddressGroup{}
