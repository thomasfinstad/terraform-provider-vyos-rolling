// Package globalprotocolsmplsldptargetedneighboripvfour code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsmplsldptargetedneighboripvfour

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsMplsLdpTargetedNeighborIPvfour{}
	_ resource.ResourceWithConfigure   = &protocolsMplsLdpTargetedNeighborIPvfour{}
	_ resource.ResourceWithImportState = &protocolsMplsLdpTargetedNeighborIPvfour{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsMplsLdpTargetedNeighborIPvfour{}
// var _ resource.ResourceWithModifyPlan = &protocolsMplsLdpTargetedNeighborIPvfour{}
// var _ resource.ResourceWithUpgradeState = &protocolsMplsLdpTargetedNeighborIPvfour{}
// var _ resource.ResourceWithValidateConfig = &protocolsMplsLdpTargetedNeighborIPvfour{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsMplsLdpTargetedNeighborIPvfour{}
