// Package namedvpnltwotpremoteaccessclientipvsixpoolprefix code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnltwotpremoteaccessclientipvsixpoolprefix

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vpnLtwotpRemoteAccessClientIPvsixPoolPrefix{}
	_ resource.ResourceWithConfigure   = &vpnLtwotpRemoteAccessClientIPvsixPoolPrefix{}
	_ resource.ResourceWithImportState = &vpnLtwotpRemoteAccessClientIPvsixPoolPrefix{}
)

// var _ resource.ResourceWithConfigValidators = &vpnLtwotpRemoteAccessClientIPvsixPoolPrefix{}
// var _ resource.ResourceWithModifyPlan = &vpnLtwotpRemoteAccessClientIPvsixPoolPrefix{}
// var _ resource.ResourceWithUpgradeState = &vpnLtwotpRemoteAccessClientIPvsixPoolPrefix{}
// var _ resource.ResourceWithValidateConfig = &vpnLtwotpRemoteAccessClientIPvsixPoolPrefix{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnLtwotpRemoteAccessClientIPvsixPoolPrefix{}
