// Package namedinterfacesethernetvifsdhcpvsixoptionspdinterface code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesethernetvifsdhcpvsixoptionspdinterface

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &interfacesEthernetVifSDhcpvsixOptionsPdInterface{}
	_ resource.ResourceWithConfigure = &interfacesEthernetVifSDhcpvsixOptionsPdInterface{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesEthernetVifSDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithModifyPlan = &interfacesEthernetVifSDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithUpgradeState = &interfacesEthernetVifSDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithValidateConfig = &interfacesEthernetVifSDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithImportState = &interfacesEthernetVifSDhcpvsixOptionsPdInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesEthernetVifSDhcpvsixOptionsPdInterface{}
