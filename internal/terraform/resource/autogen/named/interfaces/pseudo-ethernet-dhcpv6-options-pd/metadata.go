// Package namedinterfacespseudoethernetdhcpvsixoptionspd code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacespseudoethernetdhcpvsixoptionspd

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r interfacesPseudoEthernetDhcpvsixOptionsPd) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interfaces_pseudo_ethernet_dhcpv6_options_pd"
}
