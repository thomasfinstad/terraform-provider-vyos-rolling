// Package globalprotocolsospfredistributeisis code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsospfredistributeisis

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsOspfRedistributeIsis{}
	_ resource.ResourceWithConfigure = &protocolsOspfRedistributeIsis{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOspfRedistributeIsis{}
// var _ resource.ResourceWithModifyPlan = &protocolsOspfRedistributeIsis{}
// var _ resource.ResourceWithUpgradeState = &protocolsOspfRedistributeIsis{}
// var _ resource.ResourceWithValidateConfig = &protocolsOspfRedistributeIsis{}
// var _ resource.ResourceWithImportState = &protocolsOspfRedistributeIsis{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOspfRedistributeIsis{}
