// Package namedvrfnameprotocolsbgpaddressfamilyipvsixmulticastaggregateaddress code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsbgpaddressfamilyipvsixmulticastaggregateaddress

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vrfNameProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress{}
	_ resource.ResourceWithConfigure = &vrfNameProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress{}
// var _ resource.ResourceWithImportState = &vrfNameProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress{}
