/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalsystemsflow code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalsystemsflow

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (sflow) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &systemSflow{}
	_ resource.ResourceWithConfigure   = &systemSflow{}
	_ resource.ResourceWithImportState = &systemSflow{}
)

// var _ resource.ResourceWithConfigValidators = &systemSflow{}
// var _ resource.ResourceWithModifyPlan = &systemSflow{}
// var _ resource.ResourceWithUpgradeState = &systemSflow{}
// var _ resource.ResourceWithValidateConfig = &systemSflow{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &systemSflow{}
