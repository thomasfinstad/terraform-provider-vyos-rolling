// Package globalvpnsstpssl code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnsstpssl

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vpnSstpSsl{}
	_ resource.ResourceWithConfigure = &vpnSstpSsl{}
)

// var _ resource.ResourceWithConfigValidators = &vpnSstpSsl{}
// var _ resource.ResourceWithModifyPlan = &vpnSstpSsl{}
// var _ resource.ResourceWithUpgradeState = &vpnSstpSsl{}
// var _ resource.ResourceWithValidateConfig = &vpnSstpSsl{}
// var _ resource.ResourceWithImportState = &vpnSstpSsl{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnSstpSsl{}
