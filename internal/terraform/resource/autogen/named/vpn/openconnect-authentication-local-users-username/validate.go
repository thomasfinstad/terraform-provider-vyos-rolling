// Package namedvpnopenconnectauthenticationlocalusersusername code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnopenconnectauthenticationlocalusersusername

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vpnOpenconnectAuthenticationLocalUsersUsername{}
	_ resource.ResourceWithConfigure = &vpnOpenconnectAuthenticationLocalUsersUsername{}
)

// var _ resource.ResourceWithConfigValidators = &vpnOpenconnectAuthenticationLocalUsersUsername{}
// var _ resource.ResourceWithModifyPlan = &vpnOpenconnectAuthenticationLocalUsersUsername{}
// var _ resource.ResourceWithUpgradeState = &vpnOpenconnectAuthenticationLocalUsersUsername{}
// var _ resource.ResourceWithValidateConfig = &vpnOpenconnectAuthenticationLocalUsersUsername{}
// var _ resource.ResourceWithImportState = &vpnOpenconnectAuthenticationLocalUsersUsername{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnOpenconnectAuthenticationLocalUsersUsername{}