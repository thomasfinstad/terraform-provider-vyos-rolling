// Package globalprotocolsospfaggregation code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsospfaggregation

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsOspfAggregation{}
	_ resource.ResourceWithConfigure = &protocolsOspfAggregation{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOspfAggregation{}
// var _ resource.ResourceWithModifyPlan = &protocolsOspfAggregation{}
// var _ resource.ResourceWithUpgradeState = &protocolsOspfAggregation{}
// var _ resource.ResourceWithValidateConfig = &protocolsOspfAggregation{}
// var _ resource.ResourceWithImportState = &protocolsOspfAggregation{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOspfAggregation{}
