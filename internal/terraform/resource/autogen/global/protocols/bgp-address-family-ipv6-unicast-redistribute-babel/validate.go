// Package globalprotocolsbgpaddressfamilyipvsixunicastredistributebabel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbgpaddressfamilyipvsixunicastredistributebabel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsBgpAddressFamilyIPvsixUnicastRedistributeBabel{}
	_ resource.ResourceWithConfigure = &protocolsBgpAddressFamilyIPvsixUnicastRedistributeBabel{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBgpAddressFamilyIPvsixUnicastRedistributeBabel{}
// var _ resource.ResourceWithModifyPlan = &protocolsBgpAddressFamilyIPvsixUnicastRedistributeBabel{}
// var _ resource.ResourceWithUpgradeState = &protocolsBgpAddressFamilyIPvsixUnicastRedistributeBabel{}
// var _ resource.ResourceWithValidateConfig = &protocolsBgpAddressFamilyIPvsixUnicastRedistributeBabel{}
// var _ resource.ResourceWithImportState = &protocolsBgpAddressFamilyIPvsixUnicastRedistributeBabel{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBgpAddressFamilyIPvsixUnicastRedistributeBabel{}