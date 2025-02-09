/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedfirewallipvfournamerule code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallipvfournamerule

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (rule) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &firewallIPvfourNameRule{}
	_ resource.ResourceWithConfigure   = &firewallIPvfourNameRule{}
	_ resource.ResourceWithImportState = &firewallIPvfourNameRule{}
)

// var _ resource.ResourceWithConfigValidators = &firewallIPvfourNameRule{}
// var _ resource.ResourceWithModifyPlan = &firewallIPvfourNameRule{}
// var _ resource.ResourceWithUpgradeState = &firewallIPvfourNameRule{}
// var _ resource.ResourceWithValidateConfig = &firewallIPvfourNameRule{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallIPvfourNameRule{}
