// Package globalprotocolsripngredistributekernel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsripngredistributekernel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsRIPngRedistributeKernel{}
	_ resource.ResourceWithConfigure   = &protocolsRIPngRedistributeKernel{}
	_ resource.ResourceWithImportState = &protocolsRIPngRedistributeKernel{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsRIPngRedistributeKernel{}
// var _ resource.ResourceWithModifyPlan = &protocolsRIPngRedistributeKernel{}
// var _ resource.ResourceWithUpgradeState = &protocolsRIPngRedistributeKernel{}
// var _ resource.ResourceWithValidateConfig = &protocolsRIPngRedistributeKernel{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsRIPngRedistributeKernel{}
