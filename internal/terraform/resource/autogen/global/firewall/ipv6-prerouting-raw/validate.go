// Package globalfirewallipvsixpreroutingraw code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalfirewallipvsixpreroutingraw

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &firewallIPvsixPreroutingRaw{}
	_ resource.ResourceWithConfigure = &firewallIPvsixPreroutingRaw{}
)

// var _ resource.ResourceWithConfigValidators = &firewallIPvsixPreroutingRaw{}
// var _ resource.ResourceWithModifyPlan = &firewallIPvsixPreroutingRaw{}
// var _ resource.ResourceWithUpgradeState = &firewallIPvsixPreroutingRaw{}
// var _ resource.ResourceWithValidateConfig = &firewallIPvsixPreroutingRaw{}
// var _ resource.ResourceWithImportState = &firewallIPvsixPreroutingRaw{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallIPvsixPreroutingRaw{}
