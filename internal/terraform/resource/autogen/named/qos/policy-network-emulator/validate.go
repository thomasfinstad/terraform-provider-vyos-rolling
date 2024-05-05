// Package namedqospolicynetworkemulator code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedqospolicynetworkemulator

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &qosPolicyNetworkEmulator{}
	_ resource.ResourceWithConfigure = &qosPolicyNetworkEmulator{}
)

// var _ resource.ResourceWithConfigValidators = &qosPolicyNetworkEmulator{}
// var _ resource.ResourceWithModifyPlan = &qosPolicyNetworkEmulator{}
// var _ resource.ResourceWithUpgradeState = &qosPolicyNetworkEmulator{}
// var _ resource.ResourceWithValidateConfig = &qosPolicyNetworkEmulator{}
// var _ resource.ResourceWithImportState = &qosPolicyNetworkEmulator{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &qosPolicyNetworkEmulator{}
