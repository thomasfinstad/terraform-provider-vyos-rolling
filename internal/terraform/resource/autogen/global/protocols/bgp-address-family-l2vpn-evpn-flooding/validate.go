// Package globalprotocolsbgpaddressfamilyltwovpnevpnflooding code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbgpaddressfamilyltwovpnevpnflooding

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsBgpAddressFamilyLtwovpnEvpnFlooding{}
	_ resource.ResourceWithConfigure = &protocolsBgpAddressFamilyLtwovpnEvpnFlooding{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBgpAddressFamilyLtwovpnEvpnFlooding{}
// var _ resource.ResourceWithModifyPlan = &protocolsBgpAddressFamilyLtwovpnEvpnFlooding{}
// var _ resource.ResourceWithUpgradeState = &protocolsBgpAddressFamilyLtwovpnEvpnFlooding{}
// var _ resource.ResourceWithValidateConfig = &protocolsBgpAddressFamilyLtwovpnEvpnFlooding{}
// var _ resource.ResourceWithImportState = &protocolsBgpAddressFamilyLtwovpnEvpnFlooding{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBgpAddressFamilyLtwovpnEvpnFlooding{}
