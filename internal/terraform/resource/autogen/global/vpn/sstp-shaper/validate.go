// Package globalvpnsstpshaper code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnsstpshaper

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vpnSstpShaper{}
	_ resource.ResourceWithConfigure = &vpnSstpShaper{}
)

// var _ resource.ResourceWithConfigValidators = &vpnSstpShaper{}
// var _ resource.ResourceWithModifyPlan = &vpnSstpShaper{}
// var _ resource.ResourceWithUpgradeState = &vpnSstpShaper{}
// var _ resource.ResourceWithValidateConfig = &vpnSstpShaper{}
// var _ resource.ResourceWithImportState = &vpnSstpShaper{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnSstpShaper{}
