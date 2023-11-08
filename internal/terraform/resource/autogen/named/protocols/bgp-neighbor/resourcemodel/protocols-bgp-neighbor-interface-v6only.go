// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsBgpNeighborInterfaceVsixonly describes the resource data model.
type ProtocolsBgpNeighborInterfaceVsixonly struct {
	// LeafNodes
	LeafProtocolsBgpNeighborInterfaceVsixonlyPeerGroup types.String `tfsdk:"peer_group" vyos:"peer-group,omitempty"`
	LeafProtocolsBgpNeighborInterfaceVsixonlyRemoteAs  types.String `tfsdk:"remote_as" vyos:"remote-as,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborInterfaceVsixonly) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"peer_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Peer group for this peer

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Peer-group name  |

`,
		},

		"remote_as": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Neighbor BGP AS number

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-4294967294  &emsp; |  Neighbor AS number  |
    |  external  &emsp; |  Any AS different from the local AS  |
    |  internal  &emsp; |  Neighbor AS number  |

`,
		},

		// Nodes

	}
}
