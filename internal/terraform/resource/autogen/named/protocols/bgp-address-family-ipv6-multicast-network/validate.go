// Package namedprotocolsbgpaddressfamilyipvsixmulticastnetwork code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsbgpaddressfamilyipvsixmulticastnetwork

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsBgpAddressFamilyIPvsixMulticastNetwork{}
	_ resource.ResourceWithConfigure   = &protocolsBgpAddressFamilyIPvsixMulticastNetwork{}
	_ resource.ResourceWithImportState = &protocolsBgpAddressFamilyIPvsixMulticastNetwork{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBgpAddressFamilyIPvsixMulticastNetwork{}
// var _ resource.ResourceWithModifyPlan = &protocolsBgpAddressFamilyIPvsixMulticastNetwork{}
// var _ resource.ResourceWithUpgradeState = &protocolsBgpAddressFamilyIPvsixMulticastNetwork{}
// var _ resource.ResourceWithValidateConfig = &protocolsBgpAddressFamilyIPvsixMulticastNetwork{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBgpAddressFamilyIPvsixMulticastNetwork{}
