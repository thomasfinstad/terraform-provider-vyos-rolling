// Package namedinterfaceswirelessvifsvifcdhcpvsixoptionspd code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfaceswirelessvifsvifcdhcpvsixoptionspd

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &interfacesWirelessVifSVifCDhcpvsixOptionsPd{}
	_ resource.ResourceWithConfigure = &interfacesWirelessVifSVifCDhcpvsixOptionsPd{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesWirelessVifSVifCDhcpvsixOptionsPd{}
// var _ resource.ResourceWithModifyPlan = &interfacesWirelessVifSVifCDhcpvsixOptionsPd{}
// var _ resource.ResourceWithUpgradeState = &interfacesWirelessVifSVifCDhcpvsixOptionsPd{}
// var _ resource.ResourceWithValidateConfig = &interfacesWirelessVifSVifCDhcpvsixOptionsPd{}
// var _ resource.ResourceWithImportState = &interfacesWirelessVifSVifCDhcpvsixOptionsPd{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesWirelessVifSVifCDhcpvsixOptionsPd{}
