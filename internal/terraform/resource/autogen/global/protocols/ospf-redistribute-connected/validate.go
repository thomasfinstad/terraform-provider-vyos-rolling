// Package globalprotocolsospfredistributeconnected code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsospfredistributeconnected

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsOspfRedistributeConnected{}
	_ resource.ResourceWithConfigure   = &protocolsOspfRedistributeConnected{}
	_ resource.ResourceWithImportState = &protocolsOspfRedistributeConnected{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOspfRedistributeConnected{}
// var _ resource.ResourceWithModifyPlan = &protocolsOspfRedistributeConnected{}
// var _ resource.ResourceWithUpgradeState = &protocolsOspfRedistributeConnected{}
// var _ resource.ResourceWithValidateConfig = &protocolsOspfRedistributeConnected{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOspfRedistributeConnected{}
