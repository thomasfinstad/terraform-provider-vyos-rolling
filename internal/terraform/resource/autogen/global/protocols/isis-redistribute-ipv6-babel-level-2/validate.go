// Package globalprotocolsisisredistributeipvsixbabelleveltwo code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsisisredistributeipvsixbabelleveltwo

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsIsisRedistributeIPvsixBabelLevelTwo{}
	_ resource.ResourceWithConfigure = &protocolsIsisRedistributeIPvsixBabelLevelTwo{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsIsisRedistributeIPvsixBabelLevelTwo{}
// var _ resource.ResourceWithModifyPlan = &protocolsIsisRedistributeIPvsixBabelLevelTwo{}
// var _ resource.ResourceWithUpgradeState = &protocolsIsisRedistributeIPvsixBabelLevelTwo{}
// var _ resource.ResourceWithValidateConfig = &protocolsIsisRedistributeIPvsixBabelLevelTwo{}
// var _ resource.ResourceWithImportState = &protocolsIsisRedistributeIPvsixBabelLevelTwo{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsIsisRedistributeIPvsixBabelLevelTwo{}