/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedfirewallipvsixoutputrawrule code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallipvsixoutputrawrule

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (rule) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &firewallIPvsixOutputRawRule{}
	_ resource.ResourceWithConfigure   = &firewallIPvsixOutputRawRule{}
	_ resource.ResourceWithImportState = &firewallIPvsixOutputRawRule{}
)

// var _ resource.ResourceWithConfigValidators = &firewallIPvsixOutputRawRule{}
// var _ resource.ResourceWithModifyPlan = &firewallIPvsixOutputRawRule{}
// var _ resource.ResourceWithUpgradeState = &firewallIPvsixOutputRawRule{}
// var _ resource.ResourceWithValidateConfig = &firewallIPvsixOutputRawRule{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallIPvsixOutputRawRule{}
