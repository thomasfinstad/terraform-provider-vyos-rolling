// Package globalprotocolsisisredistributeipvsixconnectedlevelone code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsisisredistributeipvsixconnectedlevelone

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsIsisRedistributeIPvsixConnectedLevelOne{}
	_ resource.ResourceWithConfigure = &protocolsIsisRedistributeIPvsixConnectedLevelOne{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsIsisRedistributeIPvsixConnectedLevelOne{}
// var _ resource.ResourceWithModifyPlan = &protocolsIsisRedistributeIPvsixConnectedLevelOne{}
// var _ resource.ResourceWithUpgradeState = &protocolsIsisRedistributeIPvsixConnectedLevelOne{}
// var _ resource.ResourceWithValidateConfig = &protocolsIsisRedistributeIPvsixConnectedLevelOne{}
// var _ resource.ResourceWithImportState = &protocolsIsisRedistributeIPvsixConnectedLevelOne{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsIsisRedistributeIPvsixConnectedLevelOne{}
