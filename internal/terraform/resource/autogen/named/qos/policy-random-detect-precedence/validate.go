/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedqospolicyrandomdetectprecedence code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedqospolicyrandomdetectprecedence

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (precedence) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &qosPolicyRandomDetectPrecedence{}
	_ resource.ResourceWithConfigure   = &qosPolicyRandomDetectPrecedence{}
	_ resource.ResourceWithImportState = &qosPolicyRandomDetectPrecedence{}
)

// var _ resource.ResourceWithConfigValidators = &qosPolicyRandomDetectPrecedence{}
// var _ resource.ResourceWithModifyPlan = &qosPolicyRandomDetectPrecedence{}
// var _ resource.ResourceWithUpgradeState = &qosPolicyRandomDetectPrecedence{}
// var _ resource.ResourceWithValidateConfig = &qosPolicyRandomDetectPrecedence{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &qosPolicyRandomDetectPrecedence{}
