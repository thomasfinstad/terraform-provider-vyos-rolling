/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedfirewallgroupipvsixaddressgroup code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallgroupipvsixaddressgroup

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (ipv6-address-group) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &firewallGroupIPvsixAddressGroup{}
	_ resource.ResourceWithConfigure   = &firewallGroupIPvsixAddressGroup{}
	_ resource.ResourceWithImportState = &firewallGroupIPvsixAddressGroup{}
)

// var _ resource.ResourceWithConfigValidators = &firewallGroupIPvsixAddressGroup{}
// var _ resource.ResourceWithModifyPlan = &firewallGroupIPvsixAddressGroup{}
// var _ resource.ResourceWithUpgradeState = &firewallGroupIPvsixAddressGroup{}
// var _ resource.ResourceWithValidateConfig = &firewallGroupIPvsixAddressGroup{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallGroupIPvsixAddressGroup{}
