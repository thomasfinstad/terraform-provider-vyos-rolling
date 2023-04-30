// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsOspfAreaVirtualLinkAuthenticationMdfive describes the resource data model.
type ProtocolsOspfAreaVirtualLinkAuthenticationMdfive struct {
	// LeafNodes

	// TagNodes
	ProtocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyID types.Map `tfsdk:"key_id" json:"key-id,omitempty"`

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ProtocolsOspfAreaVirtualLinkAuthenticationMdfive) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		"key_id": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ProtocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyID{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `MD5 key id

|  Format  |  Description  |
|----------|---------------|
|  u32:1-255  |  MD5 key id  |

`,
		},

		// Nodes

	}
}
