// Package globalvpnpptpremoteaccessauthenticationradiusratelimit code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnpptpremoteaccessauthenticationradiusratelimit

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vpnPptpRemoteAccessAuthenticationRadiusRateLimit{}
	_ resource.ResourceWithConfigure = &vpnPptpRemoteAccessAuthenticationRadiusRateLimit{}
)

// var _ resource.ResourceWithConfigValidators = &vpnPptpRemoteAccessAuthenticationRadiusRateLimit{}
// var _ resource.ResourceWithModifyPlan = &vpnPptpRemoteAccessAuthenticationRadiusRateLimit{}
// var _ resource.ResourceWithUpgradeState = &vpnPptpRemoteAccessAuthenticationRadiusRateLimit{}
// var _ resource.ResourceWithValidateConfig = &vpnPptpRemoteAccessAuthenticationRadiusRateLimit{}
// var _ resource.ResourceWithImportState = &vpnPptpRemoteAccessAuthenticationRadiusRateLimit{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnPptpRemoteAccessAuthenticationRadiusRateLimit{}