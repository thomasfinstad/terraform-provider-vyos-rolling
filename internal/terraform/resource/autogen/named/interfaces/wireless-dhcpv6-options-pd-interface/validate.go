// Package namedinterfaceswirelessdhcpvsixoptionspdinterface code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfaceswirelessdhcpvsixoptionspdinterface

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &interfacesWirelessDhcpvsixOptionsPdInterface{}
	_ resource.ResourceWithConfigure = &interfacesWirelessDhcpvsixOptionsPdInterface{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesWirelessDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithModifyPlan = &interfacesWirelessDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithUpgradeState = &interfacesWirelessDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithValidateConfig = &interfacesWirelessDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithImportState = &interfacesWirelessDhcpvsixOptionsPdInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesWirelessDhcpvsixOptionsPdInterface{}
