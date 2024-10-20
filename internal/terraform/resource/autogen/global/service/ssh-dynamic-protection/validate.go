/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalservicetcpdynamicprotection code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicetcpdynamicprotection

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceTCPDynamicProtection{}
	_ resource.ResourceWithConfigure   = &serviceTCPDynamicProtection{}
	_ resource.ResourceWithImportState = &serviceTCPDynamicProtection{}
)

// var _ resource.ResourceWithConfigValidators = &serviceTCPDynamicProtection{}
// var _ resource.ResourceWithModifyPlan = &serviceTCPDynamicProtection{}
// var _ resource.ResourceWithUpgradeState = &serviceTCPDynamicProtection{}
// var _ resource.ResourceWithValidateConfig = &serviceTCPDynamicProtection{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceTCPDynamicProtection{}
