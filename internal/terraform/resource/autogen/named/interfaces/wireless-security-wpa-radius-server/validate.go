// Package namedinterfaceswirelesssecuritywparadiusserver code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfaceswirelesssecuritywparadiusserver

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &interfacesWirelessSecURItyWpaRadiusServer{}
	_ resource.ResourceWithConfigure   = &interfacesWirelessSecURItyWpaRadiusServer{}
	_ resource.ResourceWithImportState = &interfacesWirelessSecURItyWpaRadiusServer{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesWirelessSecURItyWpaRadiusServer{}
// var _ resource.ResourceWithModifyPlan = &interfacesWirelessSecURItyWpaRadiusServer{}
// var _ resource.ResourceWithUpgradeState = &interfacesWirelessSecURItyWpaRadiusServer{}
// var _ resource.ResourceWithValidateConfig = &interfacesWirelessSecURItyWpaRadiusServer{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesWirelessSecURItyWpaRadiusServer{}
