// Package globalvpnltwotpremoteaccesslns code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnltwotpremoteaccesslns

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vpnLtwotpRemoteAccessLns{}
	_ resource.ResourceWithConfigure   = &vpnLtwotpRemoteAccessLns{}
	_ resource.ResourceWithImportState = &vpnLtwotpRemoteAccessLns{}
)

// var _ resource.ResourceWithConfigValidators = &vpnLtwotpRemoteAccessLns{}
// var _ resource.ResourceWithModifyPlan = &vpnLtwotpRemoteAccessLns{}
// var _ resource.ResourceWithUpgradeState = &vpnLtwotpRemoteAccessLns{}
// var _ resource.ResourceWithValidateConfig = &vpnLtwotpRemoteAccessLns{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnLtwotpRemoteAccessLns{}
