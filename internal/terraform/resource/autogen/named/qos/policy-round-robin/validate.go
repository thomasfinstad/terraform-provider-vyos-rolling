// Package namedqospolicyroundrobin code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedqospolicyroundrobin

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &qosPolicyRoundRobin{}
	_ resource.ResourceWithConfigure = &qosPolicyRoundRobin{}
)

// var _ resource.ResourceWithConfigValidators = &qosPolicyRoundRobin{}
// var _ resource.ResourceWithModifyPlan = &qosPolicyRoundRobin{}
// var _ resource.ResourceWithUpgradeState = &qosPolicyRoundRobin{}
// var _ resource.ResourceWithValidateConfig = &qosPolicyRoundRobin{}
// var _ resource.ResourceWithImportState = &qosPolicyRoundRobin{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &qosPolicyRoundRobin{}
