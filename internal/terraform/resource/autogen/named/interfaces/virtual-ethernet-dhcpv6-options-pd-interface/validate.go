// Package namedinterfacesvirtualethernetdhcpvsixoptionspdinterface code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesvirtualethernetdhcpvsixoptionspdinterface

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &interfacesVirtualEthernetDhcpvsixOptionsPdInterface{}
	_ resource.ResourceWithConfigure = &interfacesVirtualEthernetDhcpvsixOptionsPdInterface{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesVirtualEthernetDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithModifyPlan = &interfacesVirtualEthernetDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithUpgradeState = &interfacesVirtualEthernetDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithValidateConfig = &interfacesVirtualEthernetDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithImportState = &interfacesVirtualEthernetDhcpvsixOptionsPdInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesVirtualEthernetDhcpvsixOptionsPdInterface{}