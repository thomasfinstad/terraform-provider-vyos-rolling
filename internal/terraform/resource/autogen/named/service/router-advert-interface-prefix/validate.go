// Package namedservicerouteradvertinterfaceprefix code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicerouteradvertinterfaceprefix

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceRouterAdvertInterfacePrefix{}
	_ resource.ResourceWithConfigure = &serviceRouterAdvertInterfacePrefix{}
)

// var _ resource.ResourceWithConfigValidators = &serviceRouterAdvertInterfacePrefix{}
// var _ resource.ResourceWithModifyPlan = &serviceRouterAdvertInterfacePrefix{}
// var _ resource.ResourceWithUpgradeState = &serviceRouterAdvertInterfacePrefix{}
// var _ resource.ResourceWithValidateConfig = &serviceRouterAdvertInterfacePrefix{}
// var _ resource.ResourceWithImportState = &serviceRouterAdvertInterfacePrefix{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceRouterAdvertInterfacePrefix{}
