// Package globalprotocolsisisredistributeipvfourbabellevelone code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsisisredistributeipvfourbabellevelone

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsIsisRedistributeIPvfourBabelLevelOne{}
	_ resource.ResourceWithConfigure = &protocolsIsisRedistributeIPvfourBabelLevelOne{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsIsisRedistributeIPvfourBabelLevelOne{}
// var _ resource.ResourceWithModifyPlan = &protocolsIsisRedistributeIPvfourBabelLevelOne{}
// var _ resource.ResourceWithUpgradeState = &protocolsIsisRedistributeIPvfourBabelLevelOne{}
// var _ resource.ResourceWithValidateConfig = &protocolsIsisRedistributeIPvfourBabelLevelOne{}
// var _ resource.ResourceWithImportState = &protocolsIsisRedistributeIPvfourBabelLevelOne{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsIsisRedistributeIPvfourBabelLevelOne{}