// Package globalvpnsstpauthenticationradiusratelimit code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnsstpauthenticationradiusratelimit

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vpnSstpAuthenticationRadiusRateLimit{}
	_ resource.ResourceWithConfigure = &vpnSstpAuthenticationRadiusRateLimit{}
)

// var _ resource.ResourceWithConfigValidators = &vpnSstpAuthenticationRadiusRateLimit{}
// var _ resource.ResourceWithModifyPlan = &vpnSstpAuthenticationRadiusRateLimit{}
// var _ resource.ResourceWithUpgradeState = &vpnSstpAuthenticationRadiusRateLimit{}
// var _ resource.ResourceWithValidateConfig = &vpnSstpAuthenticationRadiusRateLimit{}
// var _ resource.ResourceWithImportState = &vpnSstpAuthenticationRadiusRateLimit{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnSstpAuthenticationRadiusRateLimit{}