// Package globalvpnltwotpremoteaccessshaper code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnltwotpremoteaccessshaper

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vpnLtwotpRemoteAccessShaper{}
	_ resource.ResourceWithConfigure = &vpnLtwotpRemoteAccessShaper{}
)

// var _ resource.ResourceWithConfigValidators = &vpnLtwotpRemoteAccessShaper{}
// var _ resource.ResourceWithModifyPlan = &vpnLtwotpRemoteAccessShaper{}
// var _ resource.ResourceWithUpgradeState = &vpnLtwotpRemoteAccessShaper{}
// var _ resource.ResourceWithValidateConfig = &vpnLtwotpRemoteAccessShaper{}
// var _ resource.ResourceWithImportState = &vpnLtwotpRemoteAccessShaper{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnLtwotpRemoteAccessShaper{}
