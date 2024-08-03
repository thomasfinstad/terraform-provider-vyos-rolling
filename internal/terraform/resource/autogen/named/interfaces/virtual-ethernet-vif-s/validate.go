// Package namedinterfacesvirtualethernetvifs code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesvirtualethernetvifs

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &interfacesVirtualEthernetVifS{}
	_ resource.ResourceWithConfigure = &interfacesVirtualEthernetVifS{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesVirtualEthernetVifS{}
// var _ resource.ResourceWithModifyPlan = &interfacesVirtualEthernetVifS{}
// var _ resource.ResourceWithUpgradeState = &interfacesVirtualEthernetVifS{}
// var _ resource.ResourceWithValidateConfig = &interfacesVirtualEthernetVifS{}
// var _ resource.ResourceWithImportState = &interfacesVirtualEthernetVifS{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesVirtualEthernetVifS{}
