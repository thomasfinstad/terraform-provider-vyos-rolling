// Package globalprotocolsisisredistributeipvfourriplevelone code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsisisredistributeipvfourriplevelone

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r protocolsIsisRedistributeIPvfourRIPLevelOne) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protocols_isis_redistribute_ipv4_rip_level_1"
}
