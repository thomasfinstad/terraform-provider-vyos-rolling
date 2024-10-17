// Package namedpolicyrouterule code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedpolicyrouterule

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &policyRouteRule{}
	_ resource.ResourceWithConfigure   = &policyRouteRule{}
	_ resource.ResourceWithImportState = &policyRouteRule{}
)

// var _ resource.ResourceWithConfigValidators = &policyRouteRule{}
// var _ resource.ResourceWithModifyPlan = &policyRouteRule{}
// var _ resource.ResourceWithUpgradeState = &policyRouteRule{}
// var _ resource.ResourceWithValidateConfig = &policyRouteRule{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &policyRouteRule{}
