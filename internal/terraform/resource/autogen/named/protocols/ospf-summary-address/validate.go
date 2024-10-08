// Package namedprotocolsospfsummaryaddress code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsospfsummaryaddress

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsOspfSummaryAddress{}
	_ resource.ResourceWithConfigure = &protocolsOspfSummaryAddress{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOspfSummaryAddress{}
// var _ resource.ResourceWithModifyPlan = &protocolsOspfSummaryAddress{}
// var _ resource.ResourceWithUpgradeState = &protocolsOspfSummaryAddress{}
// var _ resource.ResourceWithValidateConfig = &protocolsOspfSummaryAddress{}
// var _ resource.ResourceWithImportState = &protocolsOspfSummaryAddress{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOspfSummaryAddress{}
