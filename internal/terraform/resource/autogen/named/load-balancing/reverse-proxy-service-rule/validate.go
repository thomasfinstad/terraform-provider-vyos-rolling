// Package namedloadbalancingreverseproxyservicerule code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedloadbalancingreverseproxyservicerule

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &loadBalancingReverseProxyServiceRule{}
	_ resource.ResourceWithConfigure = &loadBalancingReverseProxyServiceRule{}
)

// var _ resource.ResourceWithConfigValidators = &loadBalancingReverseProxyServiceRule{}
// var _ resource.ResourceWithModifyPlan = &loadBalancingReverseProxyServiceRule{}
// var _ resource.ResourceWithUpgradeState = &loadBalancingReverseProxyServiceRule{}
// var _ resource.ResourceWithValidateConfig = &loadBalancingReverseProxyServiceRule{}
// var _ resource.ResourceWithImportState = &loadBalancingReverseProxyServiceRule{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &loadBalancingReverseProxyServiceRule{}