// Package globalprotocolsisisredistributeipvfourbabelleveltwo code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsisisredistributeipvfourbabelleveltwo

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsIsisRedistributeIPvfourBabelLevelTwo{}
	_ resource.ResourceWithConfigure = &protocolsIsisRedistributeIPvfourBabelLevelTwo{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsIsisRedistributeIPvfourBabelLevelTwo{}
// var _ resource.ResourceWithModifyPlan = &protocolsIsisRedistributeIPvfourBabelLevelTwo{}
// var _ resource.ResourceWithUpgradeState = &protocolsIsisRedistributeIPvfourBabelLevelTwo{}
// var _ resource.ResourceWithValidateConfig = &protocolsIsisRedistributeIPvfourBabelLevelTwo{}
// var _ resource.ResourceWithImportState = &protocolsIsisRedistributeIPvfourBabelLevelTwo{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsIsisRedistributeIPvfourBabelLevelTwo{}
