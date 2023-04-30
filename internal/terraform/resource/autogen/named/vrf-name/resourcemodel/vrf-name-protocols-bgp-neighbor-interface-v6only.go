// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// VrfNameProtocolsBgpNeighborInterfaceVsixonly describes the resource data model.
type VrfNameProtocolsBgpNeighborInterfaceVsixonly struct {
	// LeafNodes
	VrfNameProtocolsBgpNeighborInterfaceVsixonlyPeerGroup customtypes.CustomStringValue `tfsdk:"peer_group" json:"peer-group,omitempty"`
	VrfNameProtocolsBgpNeighborInterfaceVsixonlyRemoteAs  customtypes.CustomStringValue `tfsdk:"remote_as" json:"remote-as,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborInterfaceVsixonly) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"peer_group": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Peer group for this peer

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Peer-group name  |

`,
		},

		"remote_as": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Neighbor BGP AS number

|  Format  |  Description  |
|----------|---------------|
|  u32:1-4294967294  |  Neighbor AS number  |
|  external  |  Any AS different from the local AS  |
|  internal  |  Neighbor AS number  |

`,
		},

		// TagNodes

		// Nodes

	}
}
