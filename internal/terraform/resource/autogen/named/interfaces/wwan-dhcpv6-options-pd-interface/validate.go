// Package namedinterfaceswwandhcpvsixoptionspdinterface code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfaceswwandhcpvsixoptionspdinterface

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &interfacesWwanDhcpvsixOptionsPdInterface{}
	_ resource.ResourceWithConfigure   = &interfacesWwanDhcpvsixOptionsPdInterface{}
	_ resource.ResourceWithImportState = &interfacesWwanDhcpvsixOptionsPdInterface{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesWwanDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithModifyPlan = &interfacesWwanDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithUpgradeState = &interfacesWwanDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithValidateConfig = &interfacesWwanDhcpvsixOptionsPdInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesWwanDhcpvsixOptionsPdInterface{}
