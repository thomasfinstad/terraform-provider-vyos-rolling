// Package globalprotocolsbgpaddressfamilyipvsixunicastmaximumpaths code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbgpaddressfamilyipvsixunicastmaximumpaths

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsBgpAddressFamilyIPvsixUnicastMaximumPaths{}
	_ resource.ResourceWithConfigure = &protocolsBgpAddressFamilyIPvsixUnicastMaximumPaths{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBgpAddressFamilyIPvsixUnicastMaximumPaths{}
// var _ resource.ResourceWithModifyPlan = &protocolsBgpAddressFamilyIPvsixUnicastMaximumPaths{}
// var _ resource.ResourceWithUpgradeState = &protocolsBgpAddressFamilyIPvsixUnicastMaximumPaths{}
// var _ resource.ResourceWithValidateConfig = &protocolsBgpAddressFamilyIPvsixUnicastMaximumPaths{}
// var _ resource.ResourceWithImportState = &protocolsBgpAddressFamilyIPvsixUnicastMaximumPaths{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBgpAddressFamilyIPvsixUnicastMaximumPaths{}
