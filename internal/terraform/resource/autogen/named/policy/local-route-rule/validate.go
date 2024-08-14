// Package namedpolicylocalrouterule code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedpolicylocalrouterule

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &policyLocalRouteRule{}
	_ resource.ResourceWithConfigure = &policyLocalRouteRule{}
)

// var _ resource.ResourceWithConfigValidators = &policyLocalRouteRule{}
// var _ resource.ResourceWithModifyPlan = &policyLocalRouteRule{}
// var _ resource.ResourceWithUpgradeState = &policyLocalRouteRule{}
// var _ resource.ResourceWithValidateConfig = &policyLocalRouteRule{}
// var _ resource.ResourceWithImportState = &policyLocalRouteRule{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &policyLocalRouteRule{}