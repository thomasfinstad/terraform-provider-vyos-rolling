// Package globalprotocolsospfvthreeautocost code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsospfvthreeautocost

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsOspfvthreeAutoCost{}
	_ resource.ResourceWithConfigure = &protocolsOspfvthreeAutoCost{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOspfvthreeAutoCost{}
// var _ resource.ResourceWithModifyPlan = &protocolsOspfvthreeAutoCost{}
// var _ resource.ResourceWithUpgradeState = &protocolsOspfvthreeAutoCost{}
// var _ resource.ResourceWithValidateConfig = &protocolsOspfvthreeAutoCost{}
// var _ resource.ResourceWithImportState = &protocolsOspfvthreeAutoCost{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOspfvthreeAutoCost{}