// Package namedinterfaceswirelessvifsdhcpvsixoptionspd code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfaceswirelessvifsdhcpvsixoptionspd

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &interfacesWirelessVifSDhcpvsixOptionsPd{}
	_ resource.ResourceWithConfigure   = &interfacesWirelessVifSDhcpvsixOptionsPd{}
	_ resource.ResourceWithImportState = &interfacesWirelessVifSDhcpvsixOptionsPd{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesWirelessVifSDhcpvsixOptionsPd{}
// var _ resource.ResourceWithModifyPlan = &interfacesWirelessVifSDhcpvsixOptionsPd{}
// var _ resource.ResourceWithUpgradeState = &interfacesWirelessVifSDhcpvsixOptionsPd{}
// var _ resource.ResourceWithValidateConfig = &interfacesWirelessVifSDhcpvsixOptionsPd{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesWirelessVifSDhcpvsixOptionsPd{}
