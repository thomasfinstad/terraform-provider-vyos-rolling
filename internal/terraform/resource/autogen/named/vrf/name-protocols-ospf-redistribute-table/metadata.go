// Package namedvrfnameprotocolsospfredistributetable code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsospfredistributetable

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r vrfNameProtocolsOspfRedistributeTable) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vrf_name_protocols_ospf_redistribute_table"
}
