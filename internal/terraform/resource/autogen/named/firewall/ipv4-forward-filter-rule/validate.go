// Package namedfirewallipvfourforwardfilterrule code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallipvfourforwardfilterrule

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &firewallIPvfourForwardFilterRule{}
	_ resource.ResourceWithConfigure = &firewallIPvfourForwardFilterRule{}
)

// var _ resource.ResourceWithConfigValidators = &firewallIPvfourForwardFilterRule{}
// var _ resource.ResourceWithModifyPlan = &firewallIPvfourForwardFilterRule{}
// var _ resource.ResourceWithUpgradeState = &firewallIPvfourForwardFilterRule{}
// var _ resource.ResourceWithValidateConfig = &firewallIPvfourForwardFilterRule{}
// var _ resource.ResourceWithImportState = &firewallIPvfourForwardFilterRule{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallIPvfourForwardFilterRule{}
