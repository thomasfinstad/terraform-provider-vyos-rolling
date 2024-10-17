// Package namedloadbalancinghaproxybackendserver code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedloadbalancinghaproxybackendserver

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &loadBalancingHaproxyBackendServer{}
	_ resource.ResourceWithConfigure   = &loadBalancingHaproxyBackendServer{}
	_ resource.ResourceWithImportState = &loadBalancingHaproxyBackendServer{}
)

// var _ resource.ResourceWithConfigValidators = &loadBalancingHaproxyBackendServer{}
// var _ resource.ResourceWithModifyPlan = &loadBalancingHaproxyBackendServer{}
// var _ resource.ResourceWithUpgradeState = &loadBalancingHaproxyBackendServer{}
// var _ resource.ResourceWithValidateConfig = &loadBalancingHaproxyBackendServer{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &loadBalancingHaproxyBackendServer{}
