// Package globalprotocolsbgpaddressfamilyipvsixunicastredistributestatic code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbgpaddressfamilyipvsixunicastredistributestatic

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsBgpAddressFamilyIPvsixUnicastRedistributeStatic{}
	_ resource.ResourceWithConfigure = &protocolsBgpAddressFamilyIPvsixUnicastRedistributeStatic{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBgpAddressFamilyIPvsixUnicastRedistributeStatic{}
// var _ resource.ResourceWithModifyPlan = &protocolsBgpAddressFamilyIPvsixUnicastRedistributeStatic{}
// var _ resource.ResourceWithUpgradeState = &protocolsBgpAddressFamilyIPvsixUnicastRedistributeStatic{}
// var _ resource.ResourceWithValidateConfig = &protocolsBgpAddressFamilyIPvsixUnicastRedistributeStatic{}
// var _ resource.ResourceWithImportState = &protocolsBgpAddressFamilyIPvsixUnicastRedistributeStatic{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBgpAddressFamilyIPvsixUnicastRedistributeStatic{}
