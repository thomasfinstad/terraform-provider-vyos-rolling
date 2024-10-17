// Package globalprotocolsospfvthreeredistributeconnected code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsospfvthreeredistributeconnected

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsOspfvthreeRedistributeConnected{}
	_ resource.ResourceWithConfigure   = &protocolsOspfvthreeRedistributeConnected{}
	_ resource.ResourceWithImportState = &protocolsOspfvthreeRedistributeConnected{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOspfvthreeRedistributeConnected{}
// var _ resource.ResourceWithModifyPlan = &protocolsOspfvthreeRedistributeConnected{}
// var _ resource.ResourceWithUpgradeState = &protocolsOspfvthreeRedistributeConnected{}
// var _ resource.ResourceWithValidateConfig = &protocolsOspfvthreeRedistributeConnected{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOspfvthreeRedistributeConnected{}
