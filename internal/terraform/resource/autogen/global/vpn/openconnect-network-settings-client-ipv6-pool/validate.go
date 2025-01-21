/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalvpnopenconnectnetworksettingsclientipvsixpool code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnopenconnectnetworksettingsclientipvsixpool

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (client-ipv6-pool) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vpnOpenconnectNetworkSettingsClientIPvsixPool{}
	_ resource.ResourceWithConfigure   = &vpnOpenconnectNetworkSettingsClientIPvsixPool{}
	_ resource.ResourceWithImportState = &vpnOpenconnectNetworkSettingsClientIPvsixPool{}
)

// var _ resource.ResourceWithConfigValidators = &vpnOpenconnectNetworkSettingsClientIPvsixPool{}
// var _ resource.ResourceWithModifyPlan = &vpnOpenconnectNetworkSettingsClientIPvsixPool{}
// var _ resource.ResourceWithUpgradeState = &vpnOpenconnectNetworkSettingsClientIPvsixPool{}
// var _ resource.ResourceWithValidateConfig = &vpnOpenconnectNetworkSettingsClientIPvsixPool{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnOpenconnectNetworkSettingsClientIPvsixPool{}
