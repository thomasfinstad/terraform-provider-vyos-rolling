// Package globalprotocolsospfvthreedistanceospfvthree code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsospfvthreedistanceospfvthree

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsOspfvthreeDistanceOspfvthree{}
	_ resource.ResourceWithConfigure = &protocolsOspfvthreeDistanceOspfvthree{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOspfvthreeDistanceOspfvthree{}
// var _ resource.ResourceWithModifyPlan = &protocolsOspfvthreeDistanceOspfvthree{}
// var _ resource.ResourceWithUpgradeState = &protocolsOspfvthreeDistanceOspfvthree{}
// var _ resource.ResourceWithValidateConfig = &protocolsOspfvthreeDistanceOspfvthree{}
// var _ resource.ResourceWithImportState = &protocolsOspfvthreeDistanceOspfvthree{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOspfvthreeDistanceOspfvthree{}
