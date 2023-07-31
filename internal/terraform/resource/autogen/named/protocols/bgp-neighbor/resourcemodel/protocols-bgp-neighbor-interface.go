// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsBgpNeighborInterface describes the resource data model.
type ProtocolsBgpNeighborInterface struct {
	// LeafNodes
	LeafProtocolsBgpNeighborInterfacePeerGroup       types.String `tfsdk:"peer_group" vyos:"peer-group,omitempty"`
	LeafProtocolsBgpNeighborInterfaceRemoteAs        types.String `tfsdk:"remote_as" vyos:"remote-as,omitempty"`
	LeafProtocolsBgpNeighborInterfaceSourceInterface types.String `tfsdk:"source_interface" vyos:"source-interface,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeProtocolsBgpNeighborInterfaceVsixonly *ProtocolsBgpNeighborInterfaceVsixonly `tfsdk:"v6only" vyos:"v6only,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"peer_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Peer group for this peer

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  Peer-group name  |

`,
		},

		"remote_as": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Neighbor BGP AS number

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-4294967294  |  Neighbor AS number  |
    |  external  |  Any AS different from the local AS  |
    |  internal  |  Neighbor AS number  |

`,
		},

		"source_interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interface used to establish connection

    |  Format  |  Description  |
    |----------|---------------|
    |  interface  |  Interface name  |

`,
		},

		// Nodes

		"v6only": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborInterfaceVsixonly{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Enable BGP with v6 link-local only

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBgpNeighborInterface) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsBgpNeighborInterface) UnmarshalJSON(_ []byte) error {
	return nil
}
