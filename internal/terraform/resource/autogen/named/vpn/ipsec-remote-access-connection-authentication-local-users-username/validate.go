// Package namedvpnipsecremoteaccessconnectionauthenticationlocalusersusername code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnipsecremoteaccessconnectionauthenticationlocalusersusername

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vpnIPsecRemoteAccessConnectionAuthenticationLocalUsersUsername{}
	_ resource.ResourceWithConfigure   = &vpnIPsecRemoteAccessConnectionAuthenticationLocalUsersUsername{}
	_ resource.ResourceWithImportState = &vpnIPsecRemoteAccessConnectionAuthenticationLocalUsersUsername{}
)

// var _ resource.ResourceWithConfigValidators = &vpnIPsecRemoteAccessConnectionAuthenticationLocalUsersUsername{}
// var _ resource.ResourceWithModifyPlan = &vpnIPsecRemoteAccessConnectionAuthenticationLocalUsersUsername{}
// var _ resource.ResourceWithUpgradeState = &vpnIPsecRemoteAccessConnectionAuthenticationLocalUsersUsername{}
// var _ resource.ResourceWithValidateConfig = &vpnIPsecRemoteAccessConnectionAuthenticationLocalUsersUsername{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnIPsecRemoteAccessConnectionAuthenticationLocalUsersUsername{}
