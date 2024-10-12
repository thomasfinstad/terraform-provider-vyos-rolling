// Package namedprotocolsisissegmentroutingprefix code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsisissegmentroutingprefix

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsIsisSegmentRoutingPrefix{}
	_ resource.ResourceWithConfigure = &protocolsIsisSegmentRoutingPrefix{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsIsisSegmentRoutingPrefix{}
// var _ resource.ResourceWithModifyPlan = &protocolsIsisSegmentRoutingPrefix{}
// var _ resource.ResourceWithUpgradeState = &protocolsIsisSegmentRoutingPrefix{}
// var _ resource.ResourceWithValidateConfig = &protocolsIsisSegmentRoutingPrefix{}
// var _ resource.ResourceWithImportState = &protocolsIsisSegmentRoutingPrefix{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsIsisSegmentRoutingPrefix{}
