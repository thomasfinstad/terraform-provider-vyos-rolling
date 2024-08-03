// Package globalloadbalancingwanstickyconnections code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalloadbalancingwanstickyconnections

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &loadBalancingWanStickyConnections{}
	_ resource.ResourceWithConfigure = &loadBalancingWanStickyConnections{}
)

// var _ resource.ResourceWithConfigValidators = &loadBalancingWanStickyConnections{}
// var _ resource.ResourceWithModifyPlan = &loadBalancingWanStickyConnections{}
// var _ resource.ResourceWithUpgradeState = &loadBalancingWanStickyConnections{}
// var _ resource.ResourceWithValidateConfig = &loadBalancingWanStickyConnections{}
// var _ resource.ResourceWithImportState = &loadBalancingWanStickyConnections{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &loadBalancingWanStickyConnections{}
