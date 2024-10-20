/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedqospolicyshaperhfscclassmatch code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedqospolicyshaperhfscclassmatch

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &qosPolicyShaperHfscClassMatch{}
	_ resource.ResourceWithConfigure   = &qosPolicyShaperHfscClassMatch{}
	_ resource.ResourceWithImportState = &qosPolicyShaperHfscClassMatch{}
)

// var _ resource.ResourceWithConfigValidators = &qosPolicyShaperHfscClassMatch{}
// var _ resource.ResourceWithModifyPlan = &qosPolicyShaperHfscClassMatch{}
// var _ resource.ResourceWithUpgradeState = &qosPolicyShaperHfscClassMatch{}
// var _ resource.ResourceWithValidateConfig = &qosPolicyShaperHfscClassMatch{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &qosPolicyShaperHfscClassMatch{}
