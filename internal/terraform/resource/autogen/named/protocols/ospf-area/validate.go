// Package namedprotocolsospfarea code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsospfarea

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsOspfArea{}
	_ resource.ResourceWithConfigure = &protocolsOspfArea{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOspfArea{}
// var _ resource.ResourceWithModifyPlan = &protocolsOspfArea{}
// var _ resource.ResourceWithUpgradeState = &protocolsOspfArea{}
// var _ resource.ResourceWithValidateConfig = &protocolsOspfArea{}
// var _ resource.ResourceWithImportState = &protocolsOspfArea{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOspfArea{}