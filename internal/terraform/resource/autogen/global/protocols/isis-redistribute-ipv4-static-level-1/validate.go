// Package globalprotocolsisisredistributeipvfourstaticlevelone code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsisisredistributeipvfourstaticlevelone

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsIsisRedistributeIPvfourStaticLevelOne{}
	_ resource.ResourceWithConfigure = &protocolsIsisRedistributeIPvfourStaticLevelOne{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsIsisRedistributeIPvfourStaticLevelOne{}
// var _ resource.ResourceWithModifyPlan = &protocolsIsisRedistributeIPvfourStaticLevelOne{}
// var _ resource.ResourceWithUpgradeState = &protocolsIsisRedistributeIPvfourStaticLevelOne{}
// var _ resource.ResourceWithValidateConfig = &protocolsIsisRedistributeIPvfourStaticLevelOne{}
// var _ resource.ResourceWithImportState = &protocolsIsisRedistributeIPvfourStaticLevelOne{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsIsisRedistributeIPvfourStaticLevelOne{}