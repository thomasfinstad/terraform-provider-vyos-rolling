// Package namedprotocolsospfaccesslist code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsospfaccesslist

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r protocolsOspfAccessList) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protocols_ospf_access_list"
}
