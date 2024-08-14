// Package globalvpnipsec code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnipsec

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vpnIPsec{}
	_ resource.ResourceWithConfigure = &vpnIPsec{}
)

// var _ resource.ResourceWithConfigValidators = &vpnIPsec{}
// var _ resource.ResourceWithModifyPlan = &vpnIPsec{}
// var _ resource.ResourceWithUpgradeState = &vpnIPsec{}
// var _ resource.ResourceWithValidateConfig = &vpnIPsec{}
// var _ resource.ResourceWithImportState = &vpnIPsec{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnIPsec{}