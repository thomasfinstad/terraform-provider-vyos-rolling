// Package namedfirewallbridgepreroutingfilterrule code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallbridgepreroutingfilterrule

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &firewallBrIDgePreroutingFilterRule{}
	_ resource.ResourceWithConfigure = &firewallBrIDgePreroutingFilterRule{}
)

// var _ resource.ResourceWithConfigValidators = &firewallBrIDgePreroutingFilterRule{}
// var _ resource.ResourceWithModifyPlan = &firewallBrIDgePreroutingFilterRule{}
// var _ resource.ResourceWithUpgradeState = &firewallBrIDgePreroutingFilterRule{}
// var _ resource.ResourceWithValidateConfig = &firewallBrIDgePreroutingFilterRule{}
// var _ resource.ResourceWithImportState = &firewallBrIDgePreroutingFilterRule{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallBrIDgePreroutingFilterRule{}
