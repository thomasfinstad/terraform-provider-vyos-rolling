// Package globalprotocolsisissegmentroutingglobalblock code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsisissegmentroutingglobalblock

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsIsisSegmentRoutingGlobalBlock{}
	_ resource.ResourceWithConfigure = &protocolsIsisSegmentRoutingGlobalBlock{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsIsisSegmentRoutingGlobalBlock{}
// var _ resource.ResourceWithModifyPlan = &protocolsIsisSegmentRoutingGlobalBlock{}
// var _ resource.ResourceWithUpgradeState = &protocolsIsisSegmentRoutingGlobalBlock{}
// var _ resource.ResourceWithValidateConfig = &protocolsIsisSegmentRoutingGlobalBlock{}
// var _ resource.ResourceWithImportState = &protocolsIsisSegmentRoutingGlobalBlock{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsIsisSegmentRoutingGlobalBlock{}
