// Package namedvpnsstpclientippool code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnsstpclientippool

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vpnSstpClientIPPool{}
	_ resource.ResourceWithConfigure = &vpnSstpClientIPPool{}
)

// var _ resource.ResourceWithConfigValidators = &vpnSstpClientIPPool{}
// var _ resource.ResourceWithModifyPlan = &vpnSstpClientIPPool{}
// var _ resource.ResourceWithUpgradeState = &vpnSstpClientIPPool{}
// var _ resource.ResourceWithValidateConfig = &vpnSstpClientIPPool{}
// var _ resource.ResourceWithImportState = &vpnSstpClientIPPool{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnSstpClientIPPool{}
