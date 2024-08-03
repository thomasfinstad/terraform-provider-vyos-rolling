// Package globalprotocolsripredistributestatic code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsripredistributestatic

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsRIPRedistributeStatic{}
	_ resource.ResourceWithConfigure = &protocolsRIPRedistributeStatic{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsRIPRedistributeStatic{}
// var _ resource.ResourceWithModifyPlan = &protocolsRIPRedistributeStatic{}
// var _ resource.ResourceWithUpgradeState = &protocolsRIPRedistributeStatic{}
// var _ resource.ResourceWithValidateConfig = &protocolsRIPRedistributeStatic{}
// var _ resource.ResourceWithImportState = &protocolsRIPRedistributeStatic{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsRIPRedistributeStatic{}
