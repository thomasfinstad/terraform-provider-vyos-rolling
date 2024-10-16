// Package namedloadbalancingwaninterfacehealthtest code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedloadbalancingwaninterfacehealthtest

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &loadBalancingWanInterfaceHealthTest{}
	_ resource.ResourceWithConfigure   = &loadBalancingWanInterfaceHealthTest{}
	_ resource.ResourceWithImportState = &loadBalancingWanInterfaceHealthTest{}
)

// var _ resource.ResourceWithConfigValidators = &loadBalancingWanInterfaceHealthTest{}
// var _ resource.ResourceWithModifyPlan = &loadBalancingWanInterfaceHealthTest{}
// var _ resource.ResourceWithUpgradeState = &loadBalancingWanInterfaceHealthTest{}
// var _ resource.ResourceWithValidateConfig = &loadBalancingWanInterfaceHealthTest{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &loadBalancingWanInterfaceHealthTest{}
