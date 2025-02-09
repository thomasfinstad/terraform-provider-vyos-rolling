/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalhighavailabilityvrrpglobalparametersgarp code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalhighavailabilityvrrpglobalparametersgarp

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (garp) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &highAvailabilityVrrpGlobalParametersGarp{}
	_ resource.ResourceWithConfigure   = &highAvailabilityVrrpGlobalParametersGarp{}
	_ resource.ResourceWithImportState = &highAvailabilityVrrpGlobalParametersGarp{}
)

// var _ resource.ResourceWithConfigValidators = &highAvailabilityVrrpGlobalParametersGarp{}
// var _ resource.ResourceWithModifyPlan = &highAvailabilityVrrpGlobalParametersGarp{}
// var _ resource.ResourceWithUpgradeState = &highAvailabilityVrrpGlobalParametersGarp{}
// var _ resource.ResourceWithValidateConfig = &highAvailabilityVrrpGlobalParametersGarp{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &highAvailabilityVrrpGlobalParametersGarp{}
