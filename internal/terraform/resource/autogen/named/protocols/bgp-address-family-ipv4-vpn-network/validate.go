// Package namedprotocolsbgpaddressfamilyipvfourvpnnetwork code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsbgpaddressfamilyipvfourvpnnetwork

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsBgpAddressFamilyIPvfourVpnNetwork{}
	_ resource.ResourceWithConfigure = &protocolsBgpAddressFamilyIPvfourVpnNetwork{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBgpAddressFamilyIPvfourVpnNetwork{}
// var _ resource.ResourceWithModifyPlan = &protocolsBgpAddressFamilyIPvfourVpnNetwork{}
// var _ resource.ResourceWithUpgradeState = &protocolsBgpAddressFamilyIPvfourVpnNetwork{}
// var _ resource.ResourceWithValidateConfig = &protocolsBgpAddressFamilyIPvfourVpnNetwork{}
// var _ resource.ResourceWithImportState = &protocolsBgpAddressFamilyIPvfourVpnNetwork{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBgpAddressFamilyIPvfourVpnNetwork{}
