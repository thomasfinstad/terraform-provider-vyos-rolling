/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalprotocolsospfaggregation code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsospfaggregation

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (aggregation) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsOspfAggregation{}
	_ resource.ResourceWithConfigure   = &protocolsOspfAggregation{}
	_ resource.ResourceWithImportState = &protocolsOspfAggregation{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOspfAggregation{}
// var _ resource.ResourceWithModifyPlan = &protocolsOspfAggregation{}
// var _ resource.ResourceWithUpgradeState = &protocolsOspfAggregation{}
// var _ resource.ResourceWithValidateConfig = &protocolsOspfAggregation{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOspfAggregation{}
