// Package globalloadbalancingreverseproxyglobalparameters code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalloadbalancingreverseproxyglobalparameters

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &loadBalancingReverseProxyGlobalParameters{}
	_ resource.ResourceWithConfigure = &loadBalancingReverseProxyGlobalParameters{}
)

// var _ resource.ResourceWithConfigValidators = &loadBalancingReverseProxyGlobalParameters{}
// var _ resource.ResourceWithModifyPlan = &loadBalancingReverseProxyGlobalParameters{}
// var _ resource.ResourceWithUpgradeState = &loadBalancingReverseProxyGlobalParameters{}
// var _ resource.ResourceWithValidateConfig = &loadBalancingReverseProxyGlobalParameters{}
// var _ resource.ResourceWithImportState = &loadBalancingReverseProxyGlobalParameters{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &loadBalancingReverseProxyGlobalParameters{}