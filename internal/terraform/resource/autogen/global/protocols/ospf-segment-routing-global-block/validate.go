// Package globalprotocolsospfsegmentroutingglobalblock code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsospfsegmentroutingglobalblock

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsOspfSegmentRoutingGlobalBlock{}
	_ resource.ResourceWithConfigure = &protocolsOspfSegmentRoutingGlobalBlock{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOspfSegmentRoutingGlobalBlock{}
// var _ resource.ResourceWithModifyPlan = &protocolsOspfSegmentRoutingGlobalBlock{}
// var _ resource.ResourceWithUpgradeState = &protocolsOspfSegmentRoutingGlobalBlock{}
// var _ resource.ResourceWithValidateConfig = &protocolsOspfSegmentRoutingGlobalBlock{}
// var _ resource.ResourceWithImportState = &protocolsOspfSegmentRoutingGlobalBlock{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOspfSegmentRoutingGlobalBlock{}
