// Package globalprotocolsbgpaddressfamilyipvfourunicastroutetargetvpn code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbgpaddressfamilyipvfourunicastroutetargetvpn

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r protocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpn) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protocols_bgp_address_family_ipv4_unicast_route_target_vpn"
}
