// Package globalprotocolsospfsegmentrouting code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsospfsegmentrouting

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsOspfSegmentRouting{}
	_ resource.ResourceWithConfigure = &protocolsOspfSegmentRouting{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOspfSegmentRouting{}
// var _ resource.ResourceWithModifyPlan = &protocolsOspfSegmentRouting{}
// var _ resource.ResourceWithUpgradeState = &protocolsOspfSegmentRouting{}
// var _ resource.ResourceWithValidateConfig = &protocolsOspfSegmentRouting{}
// var _ resource.ResourceWithImportState = &protocolsOspfSegmentRouting{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOspfSegmentRouting{}
