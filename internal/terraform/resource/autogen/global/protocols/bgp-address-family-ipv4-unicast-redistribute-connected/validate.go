// Package globalprotocolsbgpaddressfamilyipvfourunicastredistributeconnected code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbgpaddressfamilyipvfourunicastredistributeconnected

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsBgpAddressFamilyIPvfourUnicastRedistributeConnected{}
	_ resource.ResourceWithConfigure = &protocolsBgpAddressFamilyIPvfourUnicastRedistributeConnected{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBgpAddressFamilyIPvfourUnicastRedistributeConnected{}
// var _ resource.ResourceWithModifyPlan = &protocolsBgpAddressFamilyIPvfourUnicastRedistributeConnected{}
// var _ resource.ResourceWithUpgradeState = &protocolsBgpAddressFamilyIPvfourUnicastRedistributeConnected{}
// var _ resource.ResourceWithValidateConfig = &protocolsBgpAddressFamilyIPvfourUnicastRedistributeConnected{}
// var _ resource.ResourceWithImportState = &protocolsBgpAddressFamilyIPvfourUnicastRedistributeConnected{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBgpAddressFamilyIPvfourUnicastRedistributeConnected{}