// Package namedvrfnameprotocolsbgpneighborlocalrole code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsbgpneighborlocalrole

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vrfNameProtocolsBgpNeighborLocalRole{}
	_ resource.ResourceWithConfigure = &vrfNameProtocolsBgpNeighborLocalRole{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsBgpNeighborLocalRole{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsBgpNeighborLocalRole{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsBgpNeighborLocalRole{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsBgpNeighborLocalRole{}
// var _ resource.ResourceWithImportState = &vrfNameProtocolsBgpNeighborLocalRole{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsBgpNeighborLocalRole{}
