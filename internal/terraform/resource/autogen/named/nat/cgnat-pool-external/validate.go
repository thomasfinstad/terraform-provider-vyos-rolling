/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namednatcgnatpoolexternal code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namednatcgnatpoolexternal

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (external) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &natCgnatPoolExternal{}
	_ resource.ResourceWithConfigure   = &natCgnatPoolExternal{}
	_ resource.ResourceWithImportState = &natCgnatPoolExternal{}
)

// var _ resource.ResourceWithConfigValidators = &natCgnatPoolExternal{}
// var _ resource.ResourceWithModifyPlan = &natCgnatPoolExternal{}
// var _ resource.ResourceWithUpgradeState = &natCgnatPoolExternal{}
// var _ resource.ResourceWithValidateConfig = &natCgnatPoolExternal{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &natCgnatPoolExternal{}
