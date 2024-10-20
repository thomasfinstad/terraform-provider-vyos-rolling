/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalloadbalancingwanstickyconnections code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalloadbalancingwanstickyconnections

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &loadBalancingWanStickyConnections{}
	_ resource.ResourceWithConfigure   = &loadBalancingWanStickyConnections{}
	_ resource.ResourceWithImportState = &loadBalancingWanStickyConnections{}
)

// var _ resource.ResourceWithConfigValidators = &loadBalancingWanStickyConnections{}
// var _ resource.ResourceWithModifyPlan = &loadBalancingWanStickyConnections{}
// var _ resource.ResourceWithUpgradeState = &loadBalancingWanStickyConnections{}
// var _ resource.ResourceWithValidateConfig = &loadBalancingWanStickyConnections{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &loadBalancingWanStickyConnections{}
