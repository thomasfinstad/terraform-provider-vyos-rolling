// Package globalfirewallbridgeoutputfilter code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalfirewallbridgeoutputfilter

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &firewallBrIDgeOutputFilter{}
	_ resource.ResourceWithConfigure = &firewallBrIDgeOutputFilter{}
)

// var _ resource.ResourceWithConfigValidators = &firewallBrIDgeOutputFilter{}
// var _ resource.ResourceWithModifyPlan = &firewallBrIDgeOutputFilter{}
// var _ resource.ResourceWithUpgradeState = &firewallBrIDgeOutputFilter{}
// var _ resource.ResourceWithValidateConfig = &firewallBrIDgeOutputFilter{}
// var _ resource.ResourceWithImportState = &firewallBrIDgeOutputFilter{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallBrIDgeOutputFilter{}