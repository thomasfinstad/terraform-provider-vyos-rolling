/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalvpnltwotpremoteaccesspppoptions code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnltwotpremoteaccesspppoptions

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (ppp-options) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vpnLtwotpRemoteAccessPppOptions{}
	_ resource.ResourceWithConfigure   = &vpnLtwotpRemoteAccessPppOptions{}
	_ resource.ResourceWithImportState = &vpnLtwotpRemoteAccessPppOptions{}
)

// var _ resource.ResourceWithConfigValidators = &vpnLtwotpRemoteAccessPppOptions{}
// var _ resource.ResourceWithModifyPlan = &vpnLtwotpRemoteAccessPppOptions{}
// var _ resource.ResourceWithUpgradeState = &vpnLtwotpRemoteAccessPppOptions{}
// var _ resource.ResourceWithValidateConfig = &vpnLtwotpRemoteAccessPppOptions{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnLtwotpRemoteAccessPppOptions{}
