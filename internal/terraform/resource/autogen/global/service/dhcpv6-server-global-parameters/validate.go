// Package globalservicedhcpvsixserverglobalparameters code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicedhcpvsixserverglobalparameters

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceDhcpvsixServerGlobalParameters{}
	_ resource.ResourceWithConfigure   = &serviceDhcpvsixServerGlobalParameters{}
	_ resource.ResourceWithImportState = &serviceDhcpvsixServerGlobalParameters{}
)

// var _ resource.ResourceWithConfigValidators = &serviceDhcpvsixServerGlobalParameters{}
// var _ resource.ResourceWithModifyPlan = &serviceDhcpvsixServerGlobalParameters{}
// var _ resource.ResourceWithUpgradeState = &serviceDhcpvsixServerGlobalParameters{}
// var _ resource.ResourceWithValidateConfig = &serviceDhcpvsixServerGlobalParameters{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceDhcpvsixServerGlobalParameters{}
