// Package namedvpnopenconnectaccountingradiusserver code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnopenconnectaccountingradiusserver

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vpnOpenconnectAccountingRadiusServer{}
	_ resource.ResourceWithConfigure = &vpnOpenconnectAccountingRadiusServer{}
)

// var _ resource.ResourceWithConfigValidators = &vpnOpenconnectAccountingRadiusServer{}
// var _ resource.ResourceWithModifyPlan = &vpnOpenconnectAccountingRadiusServer{}
// var _ resource.ResourceWithUpgradeState = &vpnOpenconnectAccountingRadiusServer{}
// var _ resource.ResourceWithValidateConfig = &vpnOpenconnectAccountingRadiusServer{}
// var _ resource.ResourceWithImportState = &vpnOpenconnectAccountingRadiusServer{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnOpenconnectAccountingRadiusServer{}
