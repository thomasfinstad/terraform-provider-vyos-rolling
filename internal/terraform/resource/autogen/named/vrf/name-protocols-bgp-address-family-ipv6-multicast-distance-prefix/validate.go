// Package namedvrfnameprotocolsbgpaddressfamilyipvsixmulticastdistanceprefix code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsbgpaddressfamilyipvsixmulticastdistanceprefix

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vrfNameProtocolsBgpAddressFamilyIPvsixMulticastDistancePrefix{}
	_ resource.ResourceWithConfigure = &vrfNameProtocolsBgpAddressFamilyIPvsixMulticastDistancePrefix{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsBgpAddressFamilyIPvsixMulticastDistancePrefix{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsBgpAddressFamilyIPvsixMulticastDistancePrefix{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsBgpAddressFamilyIPvsixMulticastDistancePrefix{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsBgpAddressFamilyIPvsixMulticastDistancePrefix{}
// var _ resource.ResourceWithImportState = &vrfNameProtocolsBgpAddressFamilyIPvsixMulticastDistancePrefix{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsBgpAddressFamilyIPvsixMulticastDistancePrefix{}
