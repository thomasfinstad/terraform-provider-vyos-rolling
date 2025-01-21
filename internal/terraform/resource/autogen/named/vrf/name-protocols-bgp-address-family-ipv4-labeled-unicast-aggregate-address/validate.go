/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedvrfnameprotocolsbgpaddressfamilyipvfourlabeledunicastaggregateaddress code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsbgpaddressfamilyipvfourlabeledunicastaggregateaddress

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (aggregate-address) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastAggregateAddress{}
	_ resource.ResourceWithConfigure   = &vrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastAggregateAddress{}
	_ resource.ResourceWithImportState = &vrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastAggregateAddress{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastAggregateAddress{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastAggregateAddress{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastAggregateAddress{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastAggregateAddress{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastAggregateAddress{}
