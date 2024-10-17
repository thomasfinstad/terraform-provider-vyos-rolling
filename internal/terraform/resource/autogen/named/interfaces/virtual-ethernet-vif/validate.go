// Package namedinterfacesvirtualethernetvif code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesvirtualethernetvif

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &interfacesVirtualEthernetVif{}
	_ resource.ResourceWithConfigure   = &interfacesVirtualEthernetVif{}
	_ resource.ResourceWithImportState = &interfacesVirtualEthernetVif{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesVirtualEthernetVif{}
// var _ resource.ResourceWithModifyPlan = &interfacesVirtualEthernetVif{}
// var _ resource.ResourceWithUpgradeState = &interfacesVirtualEthernetVif{}
// var _ resource.ResourceWithValidateConfig = &interfacesVirtualEthernetVif{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesVirtualEthernetVif{}
