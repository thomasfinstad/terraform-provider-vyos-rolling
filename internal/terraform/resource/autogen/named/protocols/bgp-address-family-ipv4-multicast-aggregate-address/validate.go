/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedprotocolsbgpaddressfamilyipvfourmulticastaggregateaddress code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsbgpaddressfamilyipvfourmulticastaggregateaddress

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (aggregate-address) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsBgpAddressFamilyIPvfourMulticastAggregateAddress{}
	_ resource.ResourceWithConfigure   = &protocolsBgpAddressFamilyIPvfourMulticastAggregateAddress{}
	_ resource.ResourceWithImportState = &protocolsBgpAddressFamilyIPvfourMulticastAggregateAddress{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBgpAddressFamilyIPvfourMulticastAggregateAddress{}
// var _ resource.ResourceWithModifyPlan = &protocolsBgpAddressFamilyIPvfourMulticastAggregateAddress{}
// var _ resource.ResourceWithUpgradeState = &protocolsBgpAddressFamilyIPvfourMulticastAggregateAddress{}
// var _ resource.ResourceWithValidateConfig = &protocolsBgpAddressFamilyIPvfourMulticastAggregateAddress{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBgpAddressFamilyIPvfourMulticastAggregateAddress{}
