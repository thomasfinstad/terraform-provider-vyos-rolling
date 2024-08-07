// Package namedfirewallbridgeinputfilterrule code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallbridgeinputfilterrule

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &firewallBrIDgeInputFilterRule{}
	_ resource.ResourceWithConfigure = &firewallBrIDgeInputFilterRule{}
)

// var _ resource.ResourceWithConfigValidators = &firewallBrIDgeInputFilterRule{}
// var _ resource.ResourceWithModifyPlan = &firewallBrIDgeInputFilterRule{}
// var _ resource.ResourceWithUpgradeState = &firewallBrIDgeInputFilterRule{}
// var _ resource.ResourceWithValidateConfig = &firewallBrIDgeInputFilterRule{}
// var _ resource.ResourceWithImportState = &firewallBrIDgeInputFilterRule{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallBrIDgeInputFilterRule{}
