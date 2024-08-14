// Package namedinterfacesethernetvifdhcpvsixoptionspd code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesethernetvifdhcpvsixoptionspd

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &interfacesEthernetVifDhcpvsixOptionsPd{}
	_ resource.ResourceWithConfigure = &interfacesEthernetVifDhcpvsixOptionsPd{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesEthernetVifDhcpvsixOptionsPd{}
// var _ resource.ResourceWithModifyPlan = &interfacesEthernetVifDhcpvsixOptionsPd{}
// var _ resource.ResourceWithUpgradeState = &interfacesEthernetVifDhcpvsixOptionsPd{}
// var _ resource.ResourceWithValidateConfig = &interfacesEthernetVifDhcpvsixOptionsPd{}
// var _ resource.ResourceWithImportState = &interfacesEthernetVifDhcpvsixOptionsPd{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesEthernetVifDhcpvsixOptionsPd{}