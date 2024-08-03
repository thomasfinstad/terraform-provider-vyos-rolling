// Package globalprotocolsospfsegmentroutinglocalblock code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsospfsegmentroutinglocalblock

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsOspfSegmentRoutingLocalBlock{}
	_ resource.ResourceWithConfigure = &protocolsOspfSegmentRoutingLocalBlock{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOspfSegmentRoutingLocalBlock{}
// var _ resource.ResourceWithModifyPlan = &protocolsOspfSegmentRoutingLocalBlock{}
// var _ resource.ResourceWithUpgradeState = &protocolsOspfSegmentRoutingLocalBlock{}
// var _ resource.ResourceWithValidateConfig = &protocolsOspfSegmentRoutingLocalBlock{}
// var _ resource.ResourceWithImportState = &protocolsOspfSegmentRoutingLocalBlock{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOspfSegmentRoutingLocalBlock{}
