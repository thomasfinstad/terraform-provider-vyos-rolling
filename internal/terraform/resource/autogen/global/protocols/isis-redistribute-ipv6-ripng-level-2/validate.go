// Package globalprotocolsisisredistributeipvsixripngleveltwo code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsisisredistributeipvsixripngleveltwo

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsIsisRedistributeIPvsixRIPngLevelTwo{}
	_ resource.ResourceWithConfigure = &protocolsIsisRedistributeIPvsixRIPngLevelTwo{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsIsisRedistributeIPvsixRIPngLevelTwo{}
// var _ resource.ResourceWithModifyPlan = &protocolsIsisRedistributeIPvsixRIPngLevelTwo{}
// var _ resource.ResourceWithUpgradeState = &protocolsIsisRedistributeIPvsixRIPngLevelTwo{}
// var _ resource.ResourceWithValidateConfig = &protocolsIsisRedistributeIPvsixRIPngLevelTwo{}
// var _ resource.ResourceWithImportState = &protocolsIsisRedistributeIPvsixRIPngLevelTwo{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsIsisRedistributeIPvsixRIPngLevelTwo{}