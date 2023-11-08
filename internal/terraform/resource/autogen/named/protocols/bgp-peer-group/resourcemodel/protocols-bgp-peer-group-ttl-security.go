// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsBgpPeerGroupTTLSecURIty describes the resource data model.
type ProtocolsBgpPeerGroupTTLSecURIty struct {
	// LeafNodes
	LeafProtocolsBgpPeerGroupTTLSecURItyHops types.Number `tfsdk:"hops" vyos:"hops,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpPeerGroupTTLSecURIty) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"hops": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Number of the maximum number of hops to the BGP peer

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-254  &emsp; |  Number of hops  |

`,
		},

		// Nodes

	}
}
