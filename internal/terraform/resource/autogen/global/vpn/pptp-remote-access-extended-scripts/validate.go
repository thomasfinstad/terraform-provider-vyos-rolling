// Package globalvpnpptpremoteaccessextendedscripts code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnpptpremoteaccessextendedscripts

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vpnPptpRemoteAccessExtendedScrIPts{}
	_ resource.ResourceWithConfigure = &vpnPptpRemoteAccessExtendedScrIPts{}
)

// var _ resource.ResourceWithConfigValidators = &vpnPptpRemoteAccessExtendedScrIPts{}
// var _ resource.ResourceWithModifyPlan = &vpnPptpRemoteAccessExtendedScrIPts{}
// var _ resource.ResourceWithUpgradeState = &vpnPptpRemoteAccessExtendedScrIPts{}
// var _ resource.ResourceWithValidateConfig = &vpnPptpRemoteAccessExtendedScrIPts{}
// var _ resource.ResourceWithImportState = &vpnPptpRemoteAccessExtendedScrIPts{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnPptpRemoteAccessExtendedScrIPts{}