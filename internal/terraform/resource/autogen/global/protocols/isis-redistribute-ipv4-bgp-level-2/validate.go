// Package globalprotocolsisisredistributeipvfourbgpleveltwo code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsisisredistributeipvfourbgpleveltwo

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsIsisRedistributeIPvfourBgpLevelTwo{}
	_ resource.ResourceWithConfigure = &protocolsIsisRedistributeIPvfourBgpLevelTwo{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsIsisRedistributeIPvfourBgpLevelTwo{}
// var _ resource.ResourceWithModifyPlan = &protocolsIsisRedistributeIPvfourBgpLevelTwo{}
// var _ resource.ResourceWithUpgradeState = &protocolsIsisRedistributeIPvfourBgpLevelTwo{}
// var _ resource.ResourceWithValidateConfig = &protocolsIsisRedistributeIPvfourBgpLevelTwo{}
// var _ resource.ResourceWithImportState = &protocolsIsisRedistributeIPvfourBgpLevelTwo{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsIsisRedistributeIPvfourBgpLevelTwo{}