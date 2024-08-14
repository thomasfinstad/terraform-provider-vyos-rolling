// Package globalvpnipseclog code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnipseclog

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vpnIPsecLog{}
	_ resource.ResourceWithConfigure = &vpnIPsecLog{}
)

// var _ resource.ResourceWithConfigValidators = &vpnIPsecLog{}
// var _ resource.ResourceWithModifyPlan = &vpnIPsecLog{}
// var _ resource.ResourceWithUpgradeState = &vpnIPsecLog{}
// var _ resource.ResourceWithValidateConfig = &vpnIPsecLog{}
// var _ resource.ResourceWithImportState = &vpnIPsecLog{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnIPsecLog{}