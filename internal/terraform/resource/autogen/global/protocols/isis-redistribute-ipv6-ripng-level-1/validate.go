// Package globalprotocolsisisredistributeipvsixripnglevelone code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsisisredistributeipvsixripnglevelone

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsIsisRedistributeIPvsixRIPngLevelOne{}
	_ resource.ResourceWithConfigure = &protocolsIsisRedistributeIPvsixRIPngLevelOne{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsIsisRedistributeIPvsixRIPngLevelOne{}
// var _ resource.ResourceWithModifyPlan = &protocolsIsisRedistributeIPvsixRIPngLevelOne{}
// var _ resource.ResourceWithUpgradeState = &protocolsIsisRedistributeIPvsixRIPngLevelOne{}
// var _ resource.ResourceWithValidateConfig = &protocolsIsisRedistributeIPvsixRIPngLevelOne{}
// var _ resource.ResourceWithImportState = &protocolsIsisRedistributeIPvsixRIPngLevelOne{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsIsisRedistributeIPvsixRIPngLevelOne{}
