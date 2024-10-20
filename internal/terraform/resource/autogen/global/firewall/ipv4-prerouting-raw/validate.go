/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalfirewallipvfourpreroutingraw code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalfirewallipvfourpreroutingraw

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &firewallIPvfourPreroutingRaw{}
	_ resource.ResourceWithConfigure   = &firewallIPvfourPreroutingRaw{}
	_ resource.ResourceWithImportState = &firewallIPvfourPreroutingRaw{}
)

// var _ resource.ResourceWithConfigValidators = &firewallIPvfourPreroutingRaw{}
// var _ resource.ResourceWithModifyPlan = &firewallIPvfourPreroutingRaw{}
// var _ resource.ResourceWithUpgradeState = &firewallIPvfourPreroutingRaw{}
// var _ resource.ResourceWithValidateConfig = &firewallIPvfourPreroutingRaw{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallIPvfourPreroutingRaw{}
