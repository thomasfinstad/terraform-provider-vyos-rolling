/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalvpnopenconnectnetworksettingsclientipsettings code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnopenconnectnetworksettingsclientipsettings

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vpnOpenconnectNetworkSettingsClientIPSettings{}
	_ resource.ResourceWithConfigure   = &vpnOpenconnectNetworkSettingsClientIPSettings{}
	_ resource.ResourceWithImportState = &vpnOpenconnectNetworkSettingsClientIPSettings{}
)

// var _ resource.ResourceWithConfigValidators = &vpnOpenconnectNetworkSettingsClientIPSettings{}
// var _ resource.ResourceWithModifyPlan = &vpnOpenconnectNetworkSettingsClientIPSettings{}
// var _ resource.ResourceWithUpgradeState = &vpnOpenconnectNetworkSettingsClientIPSettings{}
// var _ resource.ResourceWithValidateConfig = &vpnOpenconnectNetworkSettingsClientIPSettings{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnOpenconnectNetworkSettingsClientIPSettings{}
