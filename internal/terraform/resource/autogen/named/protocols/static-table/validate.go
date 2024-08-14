// Package namedprotocolsstatictable code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsstatictable

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsStaticTable{}
	_ resource.ResourceWithConfigure = &protocolsStaticTable{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsStaticTable{}
// var _ resource.ResourceWithModifyPlan = &protocolsStaticTable{}
// var _ resource.ResourceWithUpgradeState = &protocolsStaticTable{}
// var _ resource.ResourceWithValidateConfig = &protocolsStaticTable{}
// var _ resource.ResourceWithImportState = &protocolsStaticTable{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsStaticTable{}