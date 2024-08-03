// Package globalvpnsstpextendedscripts code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnsstpextendedscripts

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vpnSstpExtendedScrIPts{}
	_ resource.ResourceWithConfigure = &vpnSstpExtendedScrIPts{}
)

// var _ resource.ResourceWithConfigValidators = &vpnSstpExtendedScrIPts{}
// var _ resource.ResourceWithModifyPlan = &vpnSstpExtendedScrIPts{}
// var _ resource.ResourceWithUpgradeState = &vpnSstpExtendedScrIPts{}
// var _ resource.ResourceWithValidateConfig = &vpnSstpExtendedScrIPts{}
// var _ resource.ResourceWithImportState = &vpnSstpExtendedScrIPts{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnSstpExtendedScrIPts{}
