/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedfirewallipvfourpreroutingrawrule code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallipvfourpreroutingrawrule

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (rule) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &firewallIPvfourPreroutingRawRule{}
	_ resource.ResourceWithConfigure   = &firewallIPvfourPreroutingRawRule{}
	_ resource.ResourceWithImportState = &firewallIPvfourPreroutingRawRule{}
)

// var _ resource.ResourceWithConfigValidators = &firewallIPvfourPreroutingRawRule{}
// var _ resource.ResourceWithModifyPlan = &firewallIPvfourPreroutingRawRule{}
// var _ resource.ResourceWithUpgradeState = &firewallIPvfourPreroutingRawRule{}
// var _ resource.ResourceWithValidateConfig = &firewallIPvfourPreroutingRawRule{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallIPvfourPreroutingRawRule{}
