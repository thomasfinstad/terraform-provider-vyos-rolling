// Package globalprotocolsbgpaddressfamilyipvfourunicastredistributeisis code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbgpaddressfamilyipvfourunicastredistributeisis

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r protocolsBgpAddressFamilyIPvfourUnicastRedistributeIsis) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protocols_bgp_address_family_ipv4_unicast_redistribute_isis"
}
