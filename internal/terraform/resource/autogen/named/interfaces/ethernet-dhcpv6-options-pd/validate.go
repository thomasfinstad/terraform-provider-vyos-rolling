// Package namedinterfacesethernetdhcpvsixoptionspd code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesethernetdhcpvsixoptionspd

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &interfacesEthernetDhcpvsixOptionsPd{}
	_ resource.ResourceWithConfigure   = &interfacesEthernetDhcpvsixOptionsPd{}
	_ resource.ResourceWithImportState = &interfacesEthernetDhcpvsixOptionsPd{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesEthernetDhcpvsixOptionsPd{}
// var _ resource.ResourceWithModifyPlan = &interfacesEthernetDhcpvsixOptionsPd{}
// var _ resource.ResourceWithUpgradeState = &interfacesEthernetDhcpvsixOptionsPd{}
// var _ resource.ResourceWithValidateConfig = &interfacesEthernetDhcpvsixOptionsPd{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesEthernetDhcpvsixOptionsPd{}
