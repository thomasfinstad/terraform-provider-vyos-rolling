/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalfirewallipvsixinputfilter code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalfirewallipvsixinputfilter

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &firewallIPvsixInputFilter{}
	_ resource.ResourceWithConfigure   = &firewallIPvsixInputFilter{}
	_ resource.ResourceWithImportState = &firewallIPvsixInputFilter{}
)

// var _ resource.ResourceWithConfigValidators = &firewallIPvsixInputFilter{}
// var _ resource.ResourceWithModifyPlan = &firewallIPvsixInputFilter{}
// var _ resource.ResourceWithUpgradeState = &firewallIPvsixInputFilter{}
// var _ resource.ResourceWithValidateConfig = &firewallIPvsixInputFilter{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallIPvsixInputFilter{}
