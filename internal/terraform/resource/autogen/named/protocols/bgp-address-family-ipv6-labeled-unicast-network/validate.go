// Package namedprotocolsbgpaddressfamilyipvsixlabeledunicastnetwork code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsbgpaddressfamilyipvsixlabeledunicastnetwork

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork{}
	_ resource.ResourceWithConfigure = &protocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork{}
// var _ resource.ResourceWithModifyPlan = &protocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork{}
// var _ resource.ResourceWithUpgradeState = &protocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork{}
// var _ resource.ResourceWithValidateConfig = &protocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork{}
// var _ resource.ResourceWithImportState = &protocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork{}
