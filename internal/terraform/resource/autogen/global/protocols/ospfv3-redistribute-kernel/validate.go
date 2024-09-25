// Package globalprotocolsospfvthreeredistributekernel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsospfvthreeredistributekernel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsOspfvthreeRedistributeKernel{}
	_ resource.ResourceWithConfigure = &protocolsOspfvthreeRedistributeKernel{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOspfvthreeRedistributeKernel{}
// var _ resource.ResourceWithModifyPlan = &protocolsOspfvthreeRedistributeKernel{}
// var _ resource.ResourceWithUpgradeState = &protocolsOspfvthreeRedistributeKernel{}
// var _ resource.ResourceWithValidateConfig = &protocolsOspfvthreeRedistributeKernel{}
// var _ resource.ResourceWithImportState = &protocolsOspfvthreeRedistributeKernel{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOspfvthreeRedistributeKernel{}
