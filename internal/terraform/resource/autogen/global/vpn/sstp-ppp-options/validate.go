// Package globalvpnsstppppoptions code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnsstppppoptions

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vpnSstpPppOptions{}
	_ resource.ResourceWithConfigure = &vpnSstpPppOptions{}
)

// var _ resource.ResourceWithConfigValidators = &vpnSstpPppOptions{}
// var _ resource.ResourceWithModifyPlan = &vpnSstpPppOptions{}
// var _ resource.ResourceWithUpgradeState = &vpnSstpPppOptions{}
// var _ resource.ResourceWithValidateConfig = &vpnSstpPppOptions{}
// var _ resource.ResourceWithImportState = &vpnSstpPppOptions{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnSstpPppOptions{}