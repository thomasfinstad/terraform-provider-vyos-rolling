// Package namedqospolicydroptail code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedqospolicydroptail

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &qosPolicyDropTail{}
	_ resource.ResourceWithConfigure = &qosPolicyDropTail{}
)

// var _ resource.ResourceWithConfigValidators = &qosPolicyDropTail{}
// var _ resource.ResourceWithModifyPlan = &qosPolicyDropTail{}
// var _ resource.ResourceWithUpgradeState = &qosPolicyDropTail{}
// var _ resource.ResourceWithValidateConfig = &qosPolicyDropTail{}
// var _ resource.ResourceWithImportState = &qosPolicyDropTail{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &qosPolicyDropTail{}
