// Package namedinterfacesbondingdhcpvsixoptionspd code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesbondingdhcpvsixoptionspd

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &interfacesBondingDhcpvsixOptionsPd{}
	_ resource.ResourceWithConfigure = &interfacesBondingDhcpvsixOptionsPd{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesBondingDhcpvsixOptionsPd{}
// var _ resource.ResourceWithModifyPlan = &interfacesBondingDhcpvsixOptionsPd{}
// var _ resource.ResourceWithUpgradeState = &interfacesBondingDhcpvsixOptionsPd{}
// var _ resource.ResourceWithValidateConfig = &interfacesBondingDhcpvsixOptionsPd{}
// var _ resource.ResourceWithImportState = &interfacesBondingDhcpvsixOptionsPd{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesBondingDhcpvsixOptionsPd{}
