// Package namedinterfacesbondingvifsvifcdhcpvsixoptionspdinterface code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesbondingvifsvifcdhcpvsixoptionspdinterface

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &interfacesBondingVifSVifCDhcpvsixOptionsPdInterface{}
	_ resource.ResourceWithConfigure = &interfacesBondingVifSVifCDhcpvsixOptionsPdInterface{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesBondingVifSVifCDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithModifyPlan = &interfacesBondingVifSVifCDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithUpgradeState = &interfacesBondingVifSVifCDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithValidateConfig = &interfacesBondingVifSVifCDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithImportState = &interfacesBondingVifSVifCDhcpvsixOptionsPdInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesBondingVifSVifCDhcpvsixOptionsPdInterface{}