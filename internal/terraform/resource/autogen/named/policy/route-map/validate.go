// Package namedpolicyroutemap code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedpolicyroutemap

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &policyRouteMap{}
	_ resource.ResourceWithConfigure = &policyRouteMap{}
)

// var _ resource.ResourceWithConfigValidators = &policyRouteMap{}
// var _ resource.ResourceWithModifyPlan = &policyRouteMap{}
// var _ resource.ResourceWithUpgradeState = &policyRouteMap{}
// var _ resource.ResourceWithValidateConfig = &policyRouteMap{}
// var _ resource.ResourceWithImportState = &policyRouteMap{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &policyRouteMap{}
