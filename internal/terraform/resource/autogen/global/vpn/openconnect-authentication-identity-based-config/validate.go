// Package globalvpnopenconnectauthenticationidentitybasedconfig code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnopenconnectauthenticationidentitybasedconfig

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vpnOpenconnectAuthenticationIDentityBasedConfig{}
	_ resource.ResourceWithConfigure   = &vpnOpenconnectAuthenticationIDentityBasedConfig{}
	_ resource.ResourceWithImportState = &vpnOpenconnectAuthenticationIDentityBasedConfig{}
)

// var _ resource.ResourceWithConfigValidators = &vpnOpenconnectAuthenticationIDentityBasedConfig{}
// var _ resource.ResourceWithModifyPlan = &vpnOpenconnectAuthenticationIDentityBasedConfig{}
// var _ resource.ResourceWithUpgradeState = &vpnOpenconnectAuthenticationIDentityBasedConfig{}
// var _ resource.ResourceWithValidateConfig = &vpnOpenconnectAuthenticationIDentityBasedConfig{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnOpenconnectAuthenticationIDentityBasedConfig{}
