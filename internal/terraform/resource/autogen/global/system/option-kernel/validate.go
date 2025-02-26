/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalsystemoptionkernel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalsystemoptionkernel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (kernel) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &systemOptionKernel{}
	_ resource.ResourceWithConfigure   = &systemOptionKernel{}
	_ resource.ResourceWithImportState = &systemOptionKernel{}
)

// var _ resource.ResourceWithConfigValidators = &systemOptionKernel{}
// var _ resource.ResourceWithModifyPlan = &systemOptionKernel{}
// var _ resource.ResourceWithUpgradeState = &systemOptionKernel{}
// var _ resource.ResourceWithValidateConfig = &systemOptionKernel{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &systemOptionKernel{}
