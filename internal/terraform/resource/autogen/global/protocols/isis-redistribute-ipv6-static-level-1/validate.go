// Package globalprotocolsisisredistributeipvsixstaticlevelone code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsisisredistributeipvsixstaticlevelone

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsIsisRedistributeIPvsixStaticLevelOne{}
	_ resource.ResourceWithConfigure   = &protocolsIsisRedistributeIPvsixStaticLevelOne{}
	_ resource.ResourceWithImportState = &protocolsIsisRedistributeIPvsixStaticLevelOne{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsIsisRedistributeIPvsixStaticLevelOne{}
// var _ resource.ResourceWithModifyPlan = &protocolsIsisRedistributeIPvsixStaticLevelOne{}
// var _ resource.ResourceWithUpgradeState = &protocolsIsisRedistributeIPvsixStaticLevelOne{}
// var _ resource.ResourceWithValidateConfig = &protocolsIsisRedistributeIPvsixStaticLevelOne{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsIsisRedistributeIPvsixStaticLevelOne{}
