// Package namedprotocolsstatictableroutesixnexthop code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsstatictableroutesixnexthop

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r protocolsStaticTableRoutesixNextHop) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protocols_static_table_route6_next_hop"
}