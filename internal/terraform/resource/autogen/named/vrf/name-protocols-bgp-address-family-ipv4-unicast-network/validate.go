// Package namedvrfnameprotocolsbgpaddressfamilyipvfourunicastnetwork code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsbgpaddressfamilyipvfourunicastnetwork

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vrfNameProtocolsBgpAddressFamilyIPvfourUnicastNetwork{}
	_ resource.ResourceWithConfigure = &vrfNameProtocolsBgpAddressFamilyIPvfourUnicastNetwork{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsBgpAddressFamilyIPvfourUnicastNetwork{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsBgpAddressFamilyIPvfourUnicastNetwork{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsBgpAddressFamilyIPvfourUnicastNetwork{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsBgpAddressFamilyIPvfourUnicastNetwork{}
// var _ resource.ResourceWithImportState = &vrfNameProtocolsBgpAddressFamilyIPvfourUnicastNetwork{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsBgpAddressFamilyIPvfourUnicastNetwork{}
