// Package namedprotocolsbgpaddressfamilyipvfourunicastdistanceprefix code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsbgpaddressfamilyipvfourunicastdistanceprefix

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsBgpAddressFamilyIPvfourUnicastDistancePrefix{}
	_ resource.ResourceWithConfigure = &protocolsBgpAddressFamilyIPvfourUnicastDistancePrefix{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBgpAddressFamilyIPvfourUnicastDistancePrefix{}
// var _ resource.ResourceWithModifyPlan = &protocolsBgpAddressFamilyIPvfourUnicastDistancePrefix{}
// var _ resource.ResourceWithUpgradeState = &protocolsBgpAddressFamilyIPvfourUnicastDistancePrefix{}
// var _ resource.ResourceWithValidateConfig = &protocolsBgpAddressFamilyIPvfourUnicastDistancePrefix{}
// var _ resource.ResourceWithImportState = &protocolsBgpAddressFamilyIPvfourUnicastDistancePrefix{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBgpAddressFamilyIPvfourUnicastDistancePrefix{}
