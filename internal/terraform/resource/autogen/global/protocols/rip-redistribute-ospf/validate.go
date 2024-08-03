// Package globalprotocolsripredistributeospf code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsripredistributeospf

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsRIPRedistributeOspf{}
	_ resource.ResourceWithConfigure = &protocolsRIPRedistributeOspf{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsRIPRedistributeOspf{}
// var _ resource.ResourceWithModifyPlan = &protocolsRIPRedistributeOspf{}
// var _ resource.ResourceWithUpgradeState = &protocolsRIPRedistributeOspf{}
// var _ resource.ResourceWithValidateConfig = &protocolsRIPRedistributeOspf{}
// var _ resource.ResourceWithImportState = &protocolsRIPRedistributeOspf{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsRIPRedistributeOspf{}
