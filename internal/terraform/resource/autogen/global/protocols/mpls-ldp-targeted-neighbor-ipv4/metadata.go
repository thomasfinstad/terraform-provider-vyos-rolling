// Package globalprotocolsmplsldptargetedneighboripvfour code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsmplsldptargetedneighboripvfour

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r protocolsMplsLdpTargetedNeighborIPvfour) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protocols_mpls_ldp_targeted_neighbor_ipv4"
}
