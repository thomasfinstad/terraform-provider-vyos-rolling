// Package namedprotocolsstatictableroute code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsstatictableroute

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsStaticTableRoute{}
	_ resource.ResourceWithConfigure = &protocolsStaticTableRoute{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsStaticTableRoute{}
// var _ resource.ResourceWithModifyPlan = &protocolsStaticTableRoute{}
// var _ resource.ResourceWithUpgradeState = &protocolsStaticTableRoute{}
// var _ resource.ResourceWithValidateConfig = &protocolsStaticTableRoute{}
// var _ resource.ResourceWithImportState = &protocolsStaticTableRoute{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsStaticTableRoute{}
