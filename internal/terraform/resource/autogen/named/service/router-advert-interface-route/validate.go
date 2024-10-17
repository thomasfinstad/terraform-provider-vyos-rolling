// Package namedservicerouteradvertinterfaceroute code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicerouteradvertinterfaceroute

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceRouterAdvertInterfaceRoute{}
	_ resource.ResourceWithConfigure   = &serviceRouterAdvertInterfaceRoute{}
	_ resource.ResourceWithImportState = &serviceRouterAdvertInterfaceRoute{}
)

// var _ resource.ResourceWithConfigValidators = &serviceRouterAdvertInterfaceRoute{}
// var _ resource.ResourceWithModifyPlan = &serviceRouterAdvertInterfaceRoute{}
// var _ resource.ResourceWithUpgradeState = &serviceRouterAdvertInterfaceRoute{}
// var _ resource.ResourceWithValidateConfig = &serviceRouterAdvertInterfaceRoute{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceRouterAdvertInterfaceRoute{}
