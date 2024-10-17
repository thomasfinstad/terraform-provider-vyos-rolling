// Package namedqospolicylimiter code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedqospolicylimiter

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &qosPolicyLimiter{}
	_ resource.ResourceWithConfigure   = &qosPolicyLimiter{}
	_ resource.ResourceWithImportState = &qosPolicyLimiter{}
)

// var _ resource.ResourceWithConfigValidators = &qosPolicyLimiter{}
// var _ resource.ResourceWithModifyPlan = &qosPolicyLimiter{}
// var _ resource.ResourceWithUpgradeState = &qosPolicyLimiter{}
// var _ resource.ResourceWithValidateConfig = &qosPolicyLimiter{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &qosPolicyLimiter{}
