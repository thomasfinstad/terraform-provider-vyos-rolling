/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalvpnltwotpremoteaccessipsecsettingsauthenticationxfivezeronine code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnltwotpremoteaccessipsecsettingsauthenticationxfivezeronine

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (x509) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vpnLtwotpRemoteAccessIPsecSettingsAuthenticationXfivezeronine{}
	_ resource.ResourceWithConfigure   = &vpnLtwotpRemoteAccessIPsecSettingsAuthenticationXfivezeronine{}
	_ resource.ResourceWithImportState = &vpnLtwotpRemoteAccessIPsecSettingsAuthenticationXfivezeronine{}
)

// var _ resource.ResourceWithConfigValidators = &vpnLtwotpRemoteAccessIPsecSettingsAuthenticationXfivezeronine{}
// var _ resource.ResourceWithModifyPlan = &vpnLtwotpRemoteAccessIPsecSettingsAuthenticationXfivezeronine{}
// var _ resource.ResourceWithUpgradeState = &vpnLtwotpRemoteAccessIPsecSettingsAuthenticationXfivezeronine{}
// var _ resource.ResourceWithValidateConfig = &vpnLtwotpRemoteAccessIPsecSettingsAuthenticationXfivezeronine{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnLtwotpRemoteAccessIPsecSettingsAuthenticationXfivezeronine{}
