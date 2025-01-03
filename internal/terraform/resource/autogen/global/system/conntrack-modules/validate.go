/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalsystemconntrackmodules code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalsystemconntrackmodules

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &systemConntrackModules{}
	_ resource.ResourceWithConfigure   = &systemConntrackModules{}
	_ resource.ResourceWithImportState = &systemConntrackModules{}
)

// var _ resource.ResourceWithConfigValidators = &systemConntrackModules{}
// var _ resource.ResourceWithModifyPlan = &systemConntrackModules{}
// var _ resource.ResourceWithUpgradeState = &systemConntrackModules{}
// var _ resource.ResourceWithValidateConfig = &systemConntrackModules{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &systemConntrackModules{}
