// Package namedinterfacesbondingvifsdhcpvsixoptionspdinterface code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesbondingvifsdhcpvsixoptionspdinterface

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &interfacesBondingVifSDhcpvsixOptionsPdInterface{}
	_ resource.ResourceWithConfigure   = &interfacesBondingVifSDhcpvsixOptionsPdInterface{}
	_ resource.ResourceWithImportState = &interfacesBondingVifSDhcpvsixOptionsPdInterface{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesBondingVifSDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithModifyPlan = &interfacesBondingVifSDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithUpgradeState = &interfacesBondingVifSDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithValidateConfig = &interfacesBondingVifSDhcpvsixOptionsPdInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesBondingVifSDhcpvsixOptionsPdInterface{}
