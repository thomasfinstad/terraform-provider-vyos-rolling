// Package globalprotocolsospf code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsospf

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsOspf{}
	_ resource.ResourceWithConfigure = &protocolsOspf{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOspf{}
// var _ resource.ResourceWithModifyPlan = &protocolsOspf{}
// var _ resource.ResourceWithUpgradeState = &protocolsOspf{}
// var _ resource.ResourceWithValidateConfig = &protocolsOspf{}
// var _ resource.ResourceWithImportState = &protocolsOspf{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOspf{}