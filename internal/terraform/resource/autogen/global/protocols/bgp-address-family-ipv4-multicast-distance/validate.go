// Package globalprotocolsbgpaddressfamilyipvfourmulticastdistance code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbgpaddressfamilyipvfourmulticastdistance

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsBgpAddressFamilyIPvfourMulticastDistance{}
	_ resource.ResourceWithConfigure = &protocolsBgpAddressFamilyIPvfourMulticastDistance{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBgpAddressFamilyIPvfourMulticastDistance{}
// var _ resource.ResourceWithModifyPlan = &protocolsBgpAddressFamilyIPvfourMulticastDistance{}
// var _ resource.ResourceWithUpgradeState = &protocolsBgpAddressFamilyIPvfourMulticastDistance{}
// var _ resource.ResourceWithValidateConfig = &protocolsBgpAddressFamilyIPvfourMulticastDistance{}
// var _ resource.ResourceWithImportState = &protocolsBgpAddressFamilyIPvfourMulticastDistance{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBgpAddressFamilyIPvfourMulticastDistance{}
