// Package globalvpnopenconnectnetworksettingsclientipsettings code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnopenconnectnetworksettingsclientipsettings

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r vpnOpenconnectNetworkSettingsClientIPSettings) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_openconnect_network_settings_client_ip_settings"
}
