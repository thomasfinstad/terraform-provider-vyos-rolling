// Package globalprotocolsripngredistributebabel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsripngredistributebabel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsRIPngRedistributeBabel{}
	_ resource.ResourceWithConfigure = &protocolsRIPngRedistributeBabel{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsRIPngRedistributeBabel{}
// var _ resource.ResourceWithModifyPlan = &protocolsRIPngRedistributeBabel{}
// var _ resource.ResourceWithUpgradeState = &protocolsRIPngRedistributeBabel{}
// var _ resource.ResourceWithValidateConfig = &protocolsRIPngRedistributeBabel{}
// var _ resource.ResourceWithImportState = &protocolsRIPngRedistributeBabel{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsRIPngRedistributeBabel{}
