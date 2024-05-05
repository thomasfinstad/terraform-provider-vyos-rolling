// Package namedqospolicyroundrobinclass code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedqospolicyroundrobinclass

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &qosPolicyRoundRobinClass{}
	_ resource.ResourceWithConfigure = &qosPolicyRoundRobinClass{}
)

// var _ resource.ResourceWithConfigValidators = &qosPolicyRoundRobinClass{}
// var _ resource.ResourceWithModifyPlan = &qosPolicyRoundRobinClass{}
// var _ resource.ResourceWithUpgradeState = &qosPolicyRoundRobinClass{}
// var _ resource.ResourceWithValidateConfig = &qosPolicyRoundRobinClass{}
// var _ resource.ResourceWithImportState = &qosPolicyRoundRobinClass{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &qosPolicyRoundRobinClass{}
