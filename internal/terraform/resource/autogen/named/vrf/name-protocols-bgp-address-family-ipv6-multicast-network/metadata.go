// Package namedvrfnameprotocolsbgpaddressfamilyipvsixmulticastnetwork code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsbgpaddressfamilyipvsixmulticastnetwork

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r vrfNameProtocolsBgpAddressFamilyIPvsixMulticastNetwork) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vrf_name_protocols_bgp_address_family_ipv6_multicast_network"
}
