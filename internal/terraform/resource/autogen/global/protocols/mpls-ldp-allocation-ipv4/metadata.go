// Package globalprotocolsmplsldpallocationipvfour code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsmplsldpallocationipvfour

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r protocolsMplsLdpAllocationIPvfour) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protocols_mpls_ldp_allocation_ipv4"
}
