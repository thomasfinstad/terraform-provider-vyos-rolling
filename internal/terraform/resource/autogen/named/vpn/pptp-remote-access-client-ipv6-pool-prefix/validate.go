// Package namedvpnpptpremoteaccessclientipvsixpoolprefix code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnpptpremoteaccessclientipvsixpoolprefix

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vpnPptpRemoteAccessClientIPvsixPoolPrefix{}
	_ resource.ResourceWithConfigure = &vpnPptpRemoteAccessClientIPvsixPoolPrefix{}
)

// var _ resource.ResourceWithConfigValidators = &vpnPptpRemoteAccessClientIPvsixPoolPrefix{}
// var _ resource.ResourceWithModifyPlan = &vpnPptpRemoteAccessClientIPvsixPoolPrefix{}
// var _ resource.ResourceWithUpgradeState = &vpnPptpRemoteAccessClientIPvsixPoolPrefix{}
// var _ resource.ResourceWithValidateConfig = &vpnPptpRemoteAccessClientIPvsixPoolPrefix{}
// var _ resource.ResourceWithImportState = &vpnPptpRemoteAccessClientIPvsixPoolPrefix{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnPptpRemoteAccessClientIPvsixPoolPrefix{}
