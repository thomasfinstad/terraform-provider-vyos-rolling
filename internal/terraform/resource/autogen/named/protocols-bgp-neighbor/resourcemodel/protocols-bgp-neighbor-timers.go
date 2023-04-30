// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ProtocolsBgpNeighborTimers describes the resource data model.
type ProtocolsBgpNeighborTimers struct {
	// LeafNodes
	ProtocolsBgpNeighborTimersConnect   customtypes.CustomStringValue `tfsdk:"connect" json:"connect,omitempty"`
	ProtocolsBgpNeighborTimersHoldtime  customtypes.CustomStringValue `tfsdk:"holdtime" json:"holdtime,omitempty"`
	ProtocolsBgpNeighborTimersKeepalive customtypes.CustomStringValue `tfsdk:"keepalive" json:"keepalive,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ProtocolsBgpNeighborTimers) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"connect": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `BGP connect timer for this neighbor

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Connect timer in seconds  |
|  0  |  Disable connect timer  |

`,
		},

		"holdtime": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `BGP hold timer for this neighbor

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Hold timer in seconds  |
|  0  |  Hold timer disabled  |

`,
		},

		"keepalive": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `BGP keepalive interval for this neighbor

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Keepalive interval in seconds  |

`,
		},

		// TagNodes

		// Nodes

	}
}
