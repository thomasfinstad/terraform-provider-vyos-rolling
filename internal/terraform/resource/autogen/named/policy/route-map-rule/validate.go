// Package namedpolicyroutemaprule code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedpolicyroutemaprule

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &policyRouteMapRule{}
	_ resource.ResourceWithConfigure = &policyRouteMapRule{}
)

// var _ resource.ResourceWithConfigValidators = &policyRouteMapRule{}
// var _ resource.ResourceWithModifyPlan = &policyRouteMapRule{}
// var _ resource.ResourceWithUpgradeState = &policyRouteMapRule{}
// var _ resource.ResourceWithValidateConfig = &policyRouteMapRule{}
// var _ resource.ResourceWithImportState = &policyRouteMapRule{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &policyRouteMapRule{}
