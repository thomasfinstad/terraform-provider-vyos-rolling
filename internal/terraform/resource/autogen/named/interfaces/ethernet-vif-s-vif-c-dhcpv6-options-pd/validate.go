// Package namedinterfacesethernetvifsvifcdhcpvsixoptionspd code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesethernetvifsvifcdhcpvsixoptionspd

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &interfacesEthernetVifSVifCDhcpvsixOptionsPd{}
	_ resource.ResourceWithConfigure   = &interfacesEthernetVifSVifCDhcpvsixOptionsPd{}
	_ resource.ResourceWithImportState = &interfacesEthernetVifSVifCDhcpvsixOptionsPd{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesEthernetVifSVifCDhcpvsixOptionsPd{}
// var _ resource.ResourceWithModifyPlan = &interfacesEthernetVifSVifCDhcpvsixOptionsPd{}
// var _ resource.ResourceWithUpgradeState = &interfacesEthernetVifSVifCDhcpvsixOptionsPd{}
// var _ resource.ResourceWithValidateConfig = &interfacesEthernetVifSVifCDhcpvsixOptionsPd{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesEthernetVifSVifCDhcpvsixOptionsPd{}
