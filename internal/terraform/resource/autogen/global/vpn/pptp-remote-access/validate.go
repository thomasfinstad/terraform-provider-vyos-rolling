// Package globalvpnpptpremoteaccess code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnpptpremoteaccess

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vpnPptpRemoteAccess{}
	_ resource.ResourceWithConfigure = &vpnPptpRemoteAccess{}
)

// var _ resource.ResourceWithConfigValidators = &vpnPptpRemoteAccess{}
// var _ resource.ResourceWithModifyPlan = &vpnPptpRemoteAccess{}
// var _ resource.ResourceWithUpgradeState = &vpnPptpRemoteAccess{}
// var _ resource.ResourceWithValidateConfig = &vpnPptpRemoteAccess{}
// var _ resource.ResourceWithImportState = &vpnPptpRemoteAccess{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnPptpRemoteAccess{}
