// Package globalprotocolsisisredistributeipvfourkernellevelone code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsisisredistributeipvfourkernellevelone

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsIsisRedistributeIPvfourKernelLevelOne{}
	_ resource.ResourceWithConfigure   = &protocolsIsisRedistributeIPvfourKernelLevelOne{}
	_ resource.ResourceWithImportState = &protocolsIsisRedistributeIPvfourKernelLevelOne{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsIsisRedistributeIPvfourKernelLevelOne{}
// var _ resource.ResourceWithModifyPlan = &protocolsIsisRedistributeIPvfourKernelLevelOne{}
// var _ resource.ResourceWithUpgradeState = &protocolsIsisRedistributeIPvfourKernelLevelOne{}
// var _ resource.ResourceWithValidateConfig = &protocolsIsisRedistributeIPvfourKernelLevelOne{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsIsisRedistributeIPvfourKernelLevelOne{}
