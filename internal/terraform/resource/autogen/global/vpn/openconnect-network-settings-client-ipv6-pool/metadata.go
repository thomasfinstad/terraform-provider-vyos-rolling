// Package globalvpnopenconnectnetworksettingsclientipvsixpool code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnopenconnectnetworksettingsclientipvsixpool

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r vpnOpenconnectNetworkSettingsClientIPvsixPool) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_openconnect_network_settings_client_ipv6_pool"
}
