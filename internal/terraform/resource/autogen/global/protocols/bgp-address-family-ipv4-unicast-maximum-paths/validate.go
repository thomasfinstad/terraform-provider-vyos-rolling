// Package globalprotocolsbgpaddressfamilyipvfourunicastmaximumpaths code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbgpaddressfamilyipvfourunicastmaximumpaths

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsBgpAddressFamilyIPvfourUnicastMaximumPaths{}
	_ resource.ResourceWithConfigure = &protocolsBgpAddressFamilyIPvfourUnicastMaximumPaths{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBgpAddressFamilyIPvfourUnicastMaximumPaths{}
// var _ resource.ResourceWithModifyPlan = &protocolsBgpAddressFamilyIPvfourUnicastMaximumPaths{}
// var _ resource.ResourceWithUpgradeState = &protocolsBgpAddressFamilyIPvfourUnicastMaximumPaths{}
// var _ resource.ResourceWithValidateConfig = &protocolsBgpAddressFamilyIPvfourUnicastMaximumPaths{}
// var _ resource.ResourceWithImportState = &protocolsBgpAddressFamilyIPvfourUnicastMaximumPaths{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBgpAddressFamilyIPvfourUnicastMaximumPaths{}
