/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalhighavailabilityvrrpglobalparameters code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalhighavailabilityvrrpglobalparameters

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (global-parameters) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &highAvailabilityVrrpGlobalParameters{}
	_ resource.ResourceWithConfigure   = &highAvailabilityVrrpGlobalParameters{}
	_ resource.ResourceWithImportState = &highAvailabilityVrrpGlobalParameters{}
)

// var _ resource.ResourceWithConfigValidators = &highAvailabilityVrrpGlobalParameters{}
// var _ resource.ResourceWithModifyPlan = &highAvailabilityVrrpGlobalParameters{}
// var _ resource.ResourceWithUpgradeState = &highAvailabilityVrrpGlobalParameters{}
// var _ resource.ResourceWithValidateConfig = &highAvailabilityVrrpGlobalParameters{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &highAvailabilityVrrpGlobalParameters{}
