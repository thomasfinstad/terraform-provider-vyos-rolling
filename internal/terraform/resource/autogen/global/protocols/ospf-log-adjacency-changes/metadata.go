// Package globalprotocolsospflogadjacencychanges code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsospflogadjacencychanges

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r protocolsOspfLogAdjacencyChanges) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protocols_ospf_log_adjacency_changes"
}
