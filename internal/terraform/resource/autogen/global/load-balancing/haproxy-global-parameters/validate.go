// Package globalloadbalancinghaproxyglobalparameters code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalloadbalancinghaproxyglobalparameters

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &loadBalancingHaproxyGlobalParameters{}
	_ resource.ResourceWithConfigure = &loadBalancingHaproxyGlobalParameters{}
)

// var _ resource.ResourceWithConfigValidators = &loadBalancingHaproxyGlobalParameters{}
// var _ resource.ResourceWithModifyPlan = &loadBalancingHaproxyGlobalParameters{}
// var _ resource.ResourceWithUpgradeState = &loadBalancingHaproxyGlobalParameters{}
// var _ resource.ResourceWithValidateConfig = &loadBalancingHaproxyGlobalParameters{}
// var _ resource.ResourceWithImportState = &loadBalancingHaproxyGlobalParameters{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &loadBalancingHaproxyGlobalParameters{}
