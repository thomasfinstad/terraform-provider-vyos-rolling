// Package namedqospolicyshaperhfscclass code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedqospolicyshaperhfscclass

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &qosPolicyShaperHfscClass{}
	_ resource.ResourceWithConfigure = &qosPolicyShaperHfscClass{}
)

// var _ resource.ResourceWithConfigValidators = &qosPolicyShaperHfscClass{}
// var _ resource.ResourceWithModifyPlan = &qosPolicyShaperHfscClass{}
// var _ resource.ResourceWithUpgradeState = &qosPolicyShaperHfscClass{}
// var _ resource.ResourceWithValidateConfig = &qosPolicyShaperHfscClass{}
// var _ resource.ResourceWithImportState = &qosPolicyShaperHfscClass{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &qosPolicyShaperHfscClass{}
