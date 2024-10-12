// Package globalvpnopenconnectauthenticationmode code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnopenconnectauthenticationmode

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vpnOpenconnectAuthenticationMode{}
	_ resource.ResourceWithConfigure = &vpnOpenconnectAuthenticationMode{}
)

// var _ resource.ResourceWithConfigValidators = &vpnOpenconnectAuthenticationMode{}
// var _ resource.ResourceWithModifyPlan = &vpnOpenconnectAuthenticationMode{}
// var _ resource.ResourceWithUpgradeState = &vpnOpenconnectAuthenticationMode{}
// var _ resource.ResourceWithValidateConfig = &vpnOpenconnectAuthenticationMode{}
// var _ resource.ResourceWithImportState = &vpnOpenconnectAuthenticationMode{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnOpenconnectAuthenticationMode{}
