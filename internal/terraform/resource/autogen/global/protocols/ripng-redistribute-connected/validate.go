// Package globalprotocolsripngredistributeconnected code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsripngredistributeconnected

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsRIPngRedistributeConnected{}
	_ resource.ResourceWithConfigure = &protocolsRIPngRedistributeConnected{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsRIPngRedistributeConnected{}
// var _ resource.ResourceWithModifyPlan = &protocolsRIPngRedistributeConnected{}
// var _ resource.ResourceWithUpgradeState = &protocolsRIPngRedistributeConnected{}
// var _ resource.ResourceWithValidateConfig = &protocolsRIPngRedistributeConnected{}
// var _ resource.ResourceWithImportState = &protocolsRIPngRedistributeConnected{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsRIPngRedistributeConnected{}
