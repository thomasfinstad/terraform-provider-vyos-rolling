// Package namedinterfaceswwandhcpvsixoptionspd code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfaceswwandhcpvsixoptionspd

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r interfacesWwanDhcpvsixOptionsPd) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interfaces_wwan_dhcpv6_options_pd"
}