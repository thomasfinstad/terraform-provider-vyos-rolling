// Package namedvpnpptpremoteaccessauthenticationradiusserver code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnpptpremoteaccessauthenticationradiusserver

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vpnPptpRemoteAccessAuthenticationRadiusServer{}
	_ resource.ResourceWithConfigure = &vpnPptpRemoteAccessAuthenticationRadiusServer{}
)

// var _ resource.ResourceWithConfigValidators = &vpnPptpRemoteAccessAuthenticationRadiusServer{}
// var _ resource.ResourceWithModifyPlan = &vpnPptpRemoteAccessAuthenticationRadiusServer{}
// var _ resource.ResourceWithUpgradeState = &vpnPptpRemoteAccessAuthenticationRadiusServer{}
// var _ resource.ResourceWithValidateConfig = &vpnPptpRemoteAccessAuthenticationRadiusServer{}
// var _ resource.ResourceWithImportState = &vpnPptpRemoteAccessAuthenticationRadiusServer{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnPptpRemoteAccessAuthenticationRadiusServer{}
