// Package namedprotocolsmplsldpneighbor code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsmplsldpneighbor

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsMplsLdpNeighbor{}
	_ resource.ResourceWithConfigure = &protocolsMplsLdpNeighbor{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsMplsLdpNeighbor{}
// var _ resource.ResourceWithModifyPlan = &protocolsMplsLdpNeighbor{}
// var _ resource.ResourceWithUpgradeState = &protocolsMplsLdpNeighbor{}
// var _ resource.ResourceWithValidateConfig = &protocolsMplsLdpNeighbor{}
// var _ resource.ResourceWithImportState = &protocolsMplsLdpNeighbor{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsMplsLdpNeighbor{}