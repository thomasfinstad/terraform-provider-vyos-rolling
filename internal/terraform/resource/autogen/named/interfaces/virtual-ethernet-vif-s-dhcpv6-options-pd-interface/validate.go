// Package namedinterfacesvirtualethernetvifsdhcpvsixoptionspdinterface code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesvirtualethernetvifsdhcpvsixoptionspdinterface

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &interfacesVirtualEthernetVifSDhcpvsixOptionsPdInterface{}
	_ resource.ResourceWithConfigure = &interfacesVirtualEthernetVifSDhcpvsixOptionsPdInterface{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesVirtualEthernetVifSDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithModifyPlan = &interfacesVirtualEthernetVifSDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithUpgradeState = &interfacesVirtualEthernetVifSDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithValidateConfig = &interfacesVirtualEthernetVifSDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithImportState = &interfacesVirtualEthernetVifSDhcpvsixOptionsPdInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesVirtualEthernetVifSDhcpvsixOptionsPdInterface{}
