// Package namedqospolicyratecontrol code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedqospolicyratecontrol

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &qosPolicyRateControl{}
	_ resource.ResourceWithConfigure = &qosPolicyRateControl{}
)

// var _ resource.ResourceWithConfigValidators = &qosPolicyRateControl{}
// var _ resource.ResourceWithModifyPlan = &qosPolicyRateControl{}
// var _ resource.ResourceWithUpgradeState = &qosPolicyRateControl{}
// var _ resource.ResourceWithValidateConfig = &qosPolicyRateControl{}
// var _ resource.ResourceWithImportState = &qosPolicyRateControl{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &qosPolicyRateControl{}
