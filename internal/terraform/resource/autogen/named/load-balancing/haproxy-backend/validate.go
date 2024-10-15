// Package namedloadbalancinghaproxybackend code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedloadbalancinghaproxybackend

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &loadBalancingHaproxyBackend{}
	_ resource.ResourceWithConfigure = &loadBalancingHaproxyBackend{}
)

// var _ resource.ResourceWithConfigValidators = &loadBalancingHaproxyBackend{}
// var _ resource.ResourceWithModifyPlan = &loadBalancingHaproxyBackend{}
// var _ resource.ResourceWithUpgradeState = &loadBalancingHaproxyBackend{}
// var _ resource.ResourceWithValidateConfig = &loadBalancingHaproxyBackend{}
// var _ resource.ResourceWithImportState = &loadBalancingHaproxyBackend{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &loadBalancingHaproxyBackend{}