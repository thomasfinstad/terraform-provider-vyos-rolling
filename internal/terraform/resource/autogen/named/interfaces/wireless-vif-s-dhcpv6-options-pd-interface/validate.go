// Package namedinterfaceswirelessvifsdhcpvsixoptionspdinterface code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfaceswirelessvifsdhcpvsixoptionspdinterface

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &interfacesWirelessVifSDhcpvsixOptionsPdInterface{}
	_ resource.ResourceWithConfigure = &interfacesWirelessVifSDhcpvsixOptionsPdInterface{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesWirelessVifSDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithModifyPlan = &interfacesWirelessVifSDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithUpgradeState = &interfacesWirelessVifSDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithValidateConfig = &interfacesWirelessVifSDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithImportState = &interfacesWirelessVifSDhcpvsixOptionsPdInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesWirelessVifSDhcpvsixOptionsPdInterface{}
