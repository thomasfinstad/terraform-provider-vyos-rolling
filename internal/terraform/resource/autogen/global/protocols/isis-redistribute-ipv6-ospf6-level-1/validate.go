// Package globalprotocolsisisredistributeipvsixospfsixlevelone code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsisisredistributeipvsixospfsixlevelone

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsIsisRedistributeIPvsixOspfsixLevelOne{}
	_ resource.ResourceWithConfigure = &protocolsIsisRedistributeIPvsixOspfsixLevelOne{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsIsisRedistributeIPvsixOspfsixLevelOne{}
// var _ resource.ResourceWithModifyPlan = &protocolsIsisRedistributeIPvsixOspfsixLevelOne{}
// var _ resource.ResourceWithUpgradeState = &protocolsIsisRedistributeIPvsixOspfsixLevelOne{}
// var _ resource.ResourceWithValidateConfig = &protocolsIsisRedistributeIPvsixOspfsixLevelOne{}
// var _ resource.ResourceWithImportState = &protocolsIsisRedistributeIPvsixOspfsixLevelOne{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsIsisRedistributeIPvsixOspfsixLevelOne{}
