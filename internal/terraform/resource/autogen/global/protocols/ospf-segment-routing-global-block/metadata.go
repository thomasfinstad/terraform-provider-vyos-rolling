// Package globalprotocolsospfsegmentroutingglobalblock code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsospfsegmentroutingglobalblock

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r protocolsOspfSegmentRoutingGlobalBlock) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protocols_ospf_segment_routing_global_block"
}
