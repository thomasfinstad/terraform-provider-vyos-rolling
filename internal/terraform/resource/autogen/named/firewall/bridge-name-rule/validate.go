// Package namedfirewallbridgenamerule code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallbridgenamerule

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &firewallBrIDgeNameRule{}
	_ resource.ResourceWithConfigure = &firewallBrIDgeNameRule{}
)

// var _ resource.ResourceWithConfigValidators = &firewallBrIDgeNameRule{}
// var _ resource.ResourceWithModifyPlan = &firewallBrIDgeNameRule{}
// var _ resource.ResourceWithUpgradeState = &firewallBrIDgeNameRule{}
// var _ resource.ResourceWithValidateConfig = &firewallBrIDgeNameRule{}
// var _ resource.ResourceWithImportState = &firewallBrIDgeNameRule{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallBrIDgeNameRule{}
