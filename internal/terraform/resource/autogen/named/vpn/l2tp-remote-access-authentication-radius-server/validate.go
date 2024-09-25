// Package namedvpnltwotpremoteaccessauthenticationradiusserver code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnltwotpremoteaccessauthenticationradiusserver

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vpnLtwotpRemoteAccessAuthenticationRadiusServer{}
	_ resource.ResourceWithConfigure = &vpnLtwotpRemoteAccessAuthenticationRadiusServer{}
)

// var _ resource.ResourceWithConfigValidators = &vpnLtwotpRemoteAccessAuthenticationRadiusServer{}
// var _ resource.ResourceWithModifyPlan = &vpnLtwotpRemoteAccessAuthenticationRadiusServer{}
// var _ resource.ResourceWithUpgradeState = &vpnLtwotpRemoteAccessAuthenticationRadiusServer{}
// var _ resource.ResourceWithValidateConfig = &vpnLtwotpRemoteAccessAuthenticationRadiusServer{}
// var _ resource.ResourceWithImportState = &vpnLtwotpRemoteAccessAuthenticationRadiusServer{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnLtwotpRemoteAccessAuthenticationRadiusServer{}