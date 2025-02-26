/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalservicepppoeservershaper code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicepppoeservershaper

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (shaper) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &servicePppoeServerShaper{}
	_ resource.ResourceWithConfigure   = &servicePppoeServerShaper{}
	_ resource.ResourceWithImportState = &servicePppoeServerShaper{}
)

// var _ resource.ResourceWithConfigValidators = &servicePppoeServerShaper{}
// var _ resource.ResourceWithModifyPlan = &servicePppoeServerShaper{}
// var _ resource.ResourceWithUpgradeState = &servicePppoeServerShaper{}
// var _ resource.ResourceWithValidateConfig = &servicePppoeServerShaper{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &servicePppoeServerShaper{}
