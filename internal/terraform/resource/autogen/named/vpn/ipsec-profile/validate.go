// Package namedvpnipsecprofile code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnipsecprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vpnIPsecProfile{}
	_ resource.ResourceWithConfigure   = &vpnIPsecProfile{}
	_ resource.ResourceWithImportState = &vpnIPsecProfile{}
)

// var _ resource.ResourceWithConfigValidators = &vpnIPsecProfile{}
// var _ resource.ResourceWithModifyPlan = &vpnIPsecProfile{}
// var _ resource.ResourceWithUpgradeState = &vpnIPsecProfile{}
// var _ resource.ResourceWithValidateConfig = &vpnIPsecProfile{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnIPsecProfile{}
