// Package globalprotocolsbgpaddressfamilyipvfourunicastredistributerip code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbgpaddressfamilyipvfourunicastredistributerip

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsBgpAddressFamilyIPvfourUnicastRedistributeRIP{}
	_ resource.ResourceWithConfigure = &protocolsBgpAddressFamilyIPvfourUnicastRedistributeRIP{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBgpAddressFamilyIPvfourUnicastRedistributeRIP{}
// var _ resource.ResourceWithModifyPlan = &protocolsBgpAddressFamilyIPvfourUnicastRedistributeRIP{}
// var _ resource.ResourceWithUpgradeState = &protocolsBgpAddressFamilyIPvfourUnicastRedistributeRIP{}
// var _ resource.ResourceWithValidateConfig = &protocolsBgpAddressFamilyIPvfourUnicastRedistributeRIP{}
// var _ resource.ResourceWithImportState = &protocolsBgpAddressFamilyIPvfourUnicastRedistributeRIP{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBgpAddressFamilyIPvfourUnicastRedistributeRIP{}
