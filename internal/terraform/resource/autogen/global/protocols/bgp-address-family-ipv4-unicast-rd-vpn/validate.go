// Package globalprotocolsbgpaddressfamilyipvfourunicastrdvpn code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbgpaddressfamilyipvfourunicastrdvpn

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsBgpAddressFamilyIPvfourUnicastRdVpn{}
	_ resource.ResourceWithConfigure = &protocolsBgpAddressFamilyIPvfourUnicastRdVpn{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBgpAddressFamilyIPvfourUnicastRdVpn{}
// var _ resource.ResourceWithModifyPlan = &protocolsBgpAddressFamilyIPvfourUnicastRdVpn{}
// var _ resource.ResourceWithUpgradeState = &protocolsBgpAddressFamilyIPvfourUnicastRdVpn{}
// var _ resource.ResourceWithValidateConfig = &protocolsBgpAddressFamilyIPvfourUnicastRdVpn{}
// var _ resource.ResourceWithImportState = &protocolsBgpAddressFamilyIPvfourUnicastRdVpn{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBgpAddressFamilyIPvfourUnicastRdVpn{}
