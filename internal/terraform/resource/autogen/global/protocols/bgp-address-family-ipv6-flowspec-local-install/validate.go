// Package globalprotocolsbgpaddressfamilyipvsixflowspeclocalinstall code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbgpaddressfamilyipvsixflowspeclocalinstall

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsBgpAddressFamilyIPvsixFlowspecLocalInstall{}
	_ resource.ResourceWithConfigure   = &protocolsBgpAddressFamilyIPvsixFlowspecLocalInstall{}
	_ resource.ResourceWithImportState = &protocolsBgpAddressFamilyIPvsixFlowspecLocalInstall{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBgpAddressFamilyIPvsixFlowspecLocalInstall{}
// var _ resource.ResourceWithModifyPlan = &protocolsBgpAddressFamilyIPvsixFlowspecLocalInstall{}
// var _ resource.ResourceWithUpgradeState = &protocolsBgpAddressFamilyIPvsixFlowspecLocalInstall{}
// var _ resource.ResourceWithValidateConfig = &protocolsBgpAddressFamilyIPvsixFlowspecLocalInstall{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBgpAddressFamilyIPvsixFlowspecLocalInstall{}
