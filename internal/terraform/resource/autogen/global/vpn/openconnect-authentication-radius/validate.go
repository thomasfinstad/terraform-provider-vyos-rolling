// Package globalvpnopenconnectauthenticationradius code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnopenconnectauthenticationradius

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vpnOpenconnectAuthenticationRadius{}
	_ resource.ResourceWithConfigure = &vpnOpenconnectAuthenticationRadius{}
)

// var _ resource.ResourceWithConfigValidators = &vpnOpenconnectAuthenticationRadius{}
// var _ resource.ResourceWithModifyPlan = &vpnOpenconnectAuthenticationRadius{}
// var _ resource.ResourceWithUpgradeState = &vpnOpenconnectAuthenticationRadius{}
// var _ resource.ResourceWithValidateConfig = &vpnOpenconnectAuthenticationRadius{}
// var _ resource.ResourceWithImportState = &vpnOpenconnectAuthenticationRadius{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnOpenconnectAuthenticationRadius{}
