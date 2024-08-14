// Package globalprotocolsospfdistanceospf code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsospfdistanceospf

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsOspfDistanceOspf{}
	_ resource.ResourceWithConfigure = &protocolsOspfDistanceOspf{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOspfDistanceOspf{}
// var _ resource.ResourceWithModifyPlan = &protocolsOspfDistanceOspf{}
// var _ resource.ResourceWithUpgradeState = &protocolsOspfDistanceOspf{}
// var _ resource.ResourceWithValidateConfig = &protocolsOspfDistanceOspf{}
// var _ resource.ResourceWithImportState = &protocolsOspfDistanceOspf{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOspfDistanceOspf{}