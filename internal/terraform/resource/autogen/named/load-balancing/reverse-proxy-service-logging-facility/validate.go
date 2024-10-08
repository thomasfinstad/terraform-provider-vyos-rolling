// Package namedloadbalancingreverseproxyserviceloggingfacility code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedloadbalancingreverseproxyserviceloggingfacility

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &loadBalancingReverseProxyServiceLoggingFacility{}
	_ resource.ResourceWithConfigure = &loadBalancingReverseProxyServiceLoggingFacility{}
)

// var _ resource.ResourceWithConfigValidators = &loadBalancingReverseProxyServiceLoggingFacility{}
// var _ resource.ResourceWithModifyPlan = &loadBalancingReverseProxyServiceLoggingFacility{}
// var _ resource.ResourceWithUpgradeState = &loadBalancingReverseProxyServiceLoggingFacility{}
// var _ resource.ResourceWithValidateConfig = &loadBalancingReverseProxyServiceLoggingFacility{}
// var _ resource.ResourceWithImportState = &loadBalancingReverseProxyServiceLoggingFacility{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &loadBalancingReverseProxyServiceLoggingFacility{}
