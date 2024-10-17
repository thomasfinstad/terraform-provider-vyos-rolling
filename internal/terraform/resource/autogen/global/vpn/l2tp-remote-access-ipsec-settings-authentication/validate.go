// Package globalvpnltwotpremoteaccessipsecsettingsauthentication code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnltwotpremoteaccessipsecsettingsauthentication

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vpnLtwotpRemoteAccessIPsecSettingsAuthentication{}
	_ resource.ResourceWithConfigure   = &vpnLtwotpRemoteAccessIPsecSettingsAuthentication{}
	_ resource.ResourceWithImportState = &vpnLtwotpRemoteAccessIPsecSettingsAuthentication{}
)

// var _ resource.ResourceWithConfigValidators = &vpnLtwotpRemoteAccessIPsecSettingsAuthentication{}
// var _ resource.ResourceWithModifyPlan = &vpnLtwotpRemoteAccessIPsecSettingsAuthentication{}
// var _ resource.ResourceWithUpgradeState = &vpnLtwotpRemoteAccessIPsecSettingsAuthentication{}
// var _ resource.ResourceWithValidateConfig = &vpnLtwotpRemoteAccessIPsecSettingsAuthentication{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnLtwotpRemoteAccessIPsecSettingsAuthentication{}
