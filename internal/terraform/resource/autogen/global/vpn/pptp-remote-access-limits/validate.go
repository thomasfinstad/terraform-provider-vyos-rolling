// Package globalvpnpptpremoteaccesslimits code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnpptpremoteaccesslimits

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vpnPptpRemoteAccessLimits{}
	_ resource.ResourceWithConfigure = &vpnPptpRemoteAccessLimits{}
)

// var _ resource.ResourceWithConfigValidators = &vpnPptpRemoteAccessLimits{}
// var _ resource.ResourceWithModifyPlan = &vpnPptpRemoteAccessLimits{}
// var _ resource.ResourceWithUpgradeState = &vpnPptpRemoteAccessLimits{}
// var _ resource.ResourceWithValidateConfig = &vpnPptpRemoteAccessLimits{}
// var _ resource.ResourceWithImportState = &vpnPptpRemoteAccessLimits{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnPptpRemoteAccessLimits{}
