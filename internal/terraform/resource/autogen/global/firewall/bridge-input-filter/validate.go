// Package globalfirewallbridgeinputfilter code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalfirewallbridgeinputfilter

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &firewallBrIDgeInputFilter{}
	_ resource.ResourceWithConfigure   = &firewallBrIDgeInputFilter{}
	_ resource.ResourceWithImportState = &firewallBrIDgeInputFilter{}
)

// var _ resource.ResourceWithConfigValidators = &firewallBrIDgeInputFilter{}
// var _ resource.ResourceWithModifyPlan = &firewallBrIDgeInputFilter{}
// var _ resource.ResourceWithUpgradeState = &firewallBrIDgeInputFilter{}
// var _ resource.ResourceWithValidateConfig = &firewallBrIDgeInputFilter{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallBrIDgeInputFilter{}
