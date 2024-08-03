// Package namedloadbalancingreverseproxybackendrule code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedloadbalancingreverseproxybackendrule

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &loadBalancingReverseProxyBackendRule{}
	_ resource.ResourceWithConfigure = &loadBalancingReverseProxyBackendRule{}
)

// var _ resource.ResourceWithConfigValidators = &loadBalancingReverseProxyBackendRule{}
// var _ resource.ResourceWithModifyPlan = &loadBalancingReverseProxyBackendRule{}
// var _ resource.ResourceWithUpgradeState = &loadBalancingReverseProxyBackendRule{}
// var _ resource.ResourceWithValidateConfig = &loadBalancingReverseProxyBackendRule{}
// var _ resource.ResourceWithImportState = &loadBalancingReverseProxyBackendRule{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &loadBalancingReverseProxyBackendRule{}
