// Package namedinterfacesbondingvifsdhcpvsixoptionspd code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesbondingvifsdhcpvsixoptionspd

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &interfacesBondingVifSDhcpvsixOptionsPd{}
	_ resource.ResourceWithConfigure = &interfacesBondingVifSDhcpvsixOptionsPd{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesBondingVifSDhcpvsixOptionsPd{}
// var _ resource.ResourceWithModifyPlan = &interfacesBondingVifSDhcpvsixOptionsPd{}
// var _ resource.ResourceWithUpgradeState = &interfacesBondingVifSDhcpvsixOptionsPd{}
// var _ resource.ResourceWithValidateConfig = &interfacesBondingVifSDhcpvsixOptionsPd{}
// var _ resource.ResourceWithImportState = &interfacesBondingVifSDhcpvsixOptionsPd{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesBondingVifSDhcpvsixOptionsPd{}