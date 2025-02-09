/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalvpnltwotpremoteaccessipsecsettings code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnltwotpremoteaccessipsecsettings

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (ipsec-settings) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vpnLtwotpRemoteAccessIPsecSettings{}
	_ resource.ResourceWithConfigure   = &vpnLtwotpRemoteAccessIPsecSettings{}
	_ resource.ResourceWithImportState = &vpnLtwotpRemoteAccessIPsecSettings{}
)

// var _ resource.ResourceWithConfigValidators = &vpnLtwotpRemoteAccessIPsecSettings{}
// var _ resource.ResourceWithModifyPlan = &vpnLtwotpRemoteAccessIPsecSettings{}
// var _ resource.ResourceWithUpgradeState = &vpnLtwotpRemoteAccessIPsecSettings{}
// var _ resource.ResourceWithValidateConfig = &vpnLtwotpRemoteAccessIPsecSettings{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnLtwotpRemoteAccessIPsecSettings{}
