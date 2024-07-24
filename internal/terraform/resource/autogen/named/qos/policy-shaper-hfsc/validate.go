// Package namedqospolicyshaperhfsc code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedqospolicyshaperhfsc

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &qosPolicyShaperHfsc{}
	_ resource.ResourceWithConfigure = &qosPolicyShaperHfsc{}
)

// var _ resource.ResourceWithConfigValidators = &qosPolicyShaperHfsc{}
// var _ resource.ResourceWithModifyPlan = &qosPolicyShaperHfsc{}
// var _ resource.ResourceWithUpgradeState = &qosPolicyShaperHfsc{}
// var _ resource.ResourceWithValidateConfig = &qosPolicyShaperHfsc{}
// var _ resource.ResourceWithImportState = &qosPolicyShaperHfsc{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &qosPolicyShaperHfsc{}
