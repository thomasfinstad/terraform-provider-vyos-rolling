// Package namedvrfnameprotocolsbgpaddressfamilyipvsixlabeledunicastaggregateaddress code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsbgpaddressfamilyipvsixlabeledunicastaggregateaddress

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicastAggregateAddress{}
	_ resource.ResourceWithConfigure   = &vrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicastAggregateAddress{}
	_ resource.ResourceWithImportState = &vrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicastAggregateAddress{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicastAggregateAddress{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicastAggregateAddress{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicastAggregateAddress{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicastAggregateAddress{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicastAggregateAddress{}
