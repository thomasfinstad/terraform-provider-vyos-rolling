// Package globalvpnsstpsnmp code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnsstpsnmp

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vpnSstpSnmp{}
	_ resource.ResourceWithConfigure = &vpnSstpSnmp{}
)

// var _ resource.ResourceWithConfigValidators = &vpnSstpSnmp{}
// var _ resource.ResourceWithModifyPlan = &vpnSstpSnmp{}
// var _ resource.ResourceWithUpgradeState = &vpnSstpSnmp{}
// var _ resource.ResourceWithValidateConfig = &vpnSstpSnmp{}
// var _ resource.ResourceWithImportState = &vpnSstpSnmp{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnSstpSnmp{}
