// Package namedvpnipsecremoteaccessradiusserver code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnipsecremoteaccessradiusserver

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vpnIPsecRemoteAccessRadiusServer{}
	_ resource.ResourceWithConfigure = &vpnIPsecRemoteAccessRadiusServer{}
)

// var _ resource.ResourceWithConfigValidators = &vpnIPsecRemoteAccessRadiusServer{}
// var _ resource.ResourceWithModifyPlan = &vpnIPsecRemoteAccessRadiusServer{}
// var _ resource.ResourceWithUpgradeState = &vpnIPsecRemoteAccessRadiusServer{}
// var _ resource.ResourceWithValidateConfig = &vpnIPsecRemoteAccessRadiusServer{}
// var _ resource.ResourceWithImportState = &vpnIPsecRemoteAccessRadiusServer{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnIPsecRemoteAccessRadiusServer{}
