// Package globalprotocolsbgpaddressfamilyipvsixunicastredistributeconnected code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbgpaddressfamilyipvsixunicastredistributeconnected

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsBgpAddressFamilyIPvsixUnicastRedistributeConnected{}
	_ resource.ResourceWithConfigure   = &protocolsBgpAddressFamilyIPvsixUnicastRedistributeConnected{}
	_ resource.ResourceWithImportState = &protocolsBgpAddressFamilyIPvsixUnicastRedistributeConnected{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBgpAddressFamilyIPvsixUnicastRedistributeConnected{}
// var _ resource.ResourceWithModifyPlan = &protocolsBgpAddressFamilyIPvsixUnicastRedistributeConnected{}
// var _ resource.ResourceWithUpgradeState = &protocolsBgpAddressFamilyIPvsixUnicastRedistributeConnected{}
// var _ resource.ResourceWithValidateConfig = &protocolsBgpAddressFamilyIPvsixUnicastRedistributeConnected{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBgpAddressFamilyIPvsixUnicastRedistributeConnected{}
