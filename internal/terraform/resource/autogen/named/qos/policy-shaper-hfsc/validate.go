/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedqospolicyshaperhfsc code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedqospolicyshaperhfsc

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (shaper-hfsc) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &qosPolicyShaperHfsc{}
	_ resource.ResourceWithConfigure   = &qosPolicyShaperHfsc{}
	_ resource.ResourceWithImportState = &qosPolicyShaperHfsc{}
)

// var _ resource.ResourceWithConfigValidators = &qosPolicyShaperHfsc{}
// var _ resource.ResourceWithModifyPlan = &qosPolicyShaperHfsc{}
// var _ resource.ResourceWithUpgradeState = &qosPolicyShaperHfsc{}
// var _ resource.ResourceWithValidateConfig = &qosPolicyShaperHfsc{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &qosPolicyShaperHfsc{}
