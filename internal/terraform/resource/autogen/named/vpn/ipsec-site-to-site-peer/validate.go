// Package namedvpnipsecsitetositepeer code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnipsecsitetositepeer

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vpnIPsecSiteToSitePeer{}
	_ resource.ResourceWithConfigure   = &vpnIPsecSiteToSitePeer{}
	_ resource.ResourceWithImportState = &vpnIPsecSiteToSitePeer{}
)

// var _ resource.ResourceWithConfigValidators = &vpnIPsecSiteToSitePeer{}
// var _ resource.ResourceWithModifyPlan = &vpnIPsecSiteToSitePeer{}
// var _ resource.ResourceWithUpgradeState = &vpnIPsecSiteToSitePeer{}
// var _ resource.ResourceWithValidateConfig = &vpnIPsecSiteToSitePeer{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnIPsecSiteToSitePeer{}
