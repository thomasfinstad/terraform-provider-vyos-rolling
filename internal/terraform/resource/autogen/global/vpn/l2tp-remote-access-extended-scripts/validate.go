/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalvpnltwotpremoteaccessextendedscripts code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnltwotpremoteaccessextendedscripts

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vpnLtwotpRemoteAccessExtendedScrIPts{}
	_ resource.ResourceWithConfigure   = &vpnLtwotpRemoteAccessExtendedScrIPts{}
	_ resource.ResourceWithImportState = &vpnLtwotpRemoteAccessExtendedScrIPts{}
)

// var _ resource.ResourceWithConfigValidators = &vpnLtwotpRemoteAccessExtendedScrIPts{}
// var _ resource.ResourceWithModifyPlan = &vpnLtwotpRemoteAccessExtendedScrIPts{}
// var _ resource.ResourceWithUpgradeState = &vpnLtwotpRemoteAccessExtendedScrIPts{}
// var _ resource.ResourceWithValidateConfig = &vpnLtwotpRemoteAccessExtendedScrIPts{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnLtwotpRemoteAccessExtendedScrIPts{}
