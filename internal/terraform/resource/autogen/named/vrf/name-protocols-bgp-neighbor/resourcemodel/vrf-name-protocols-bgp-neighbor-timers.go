// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsBgpNeighborTimers describes the resource data model.
type VrfNameProtocolsBgpNeighborTimers struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborTimersConnect   types.String `tfsdk:"connect" vyos:"connect,omitempty"`
	LeafVrfNameProtocolsBgpNeighborTimersHoldtime  types.String `tfsdk:"holdtime" vyos:"holdtime,omitempty"`
	LeafVrfNameProtocolsBgpNeighborTimersKeepalive types.Number `tfsdk:"keepalive" vyos:"keepalive,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborTimers) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"connect": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `BGP connect timer for this neighbor

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Connect timer in seconds  |
    |  0  &emsp; |  Disable connect timer  |

`,
		},

		"holdtime": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `BGP hold timer for this neighbor

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Hold timer in seconds  |
    |  0  &emsp; |  Hold timer disabled  |

`,
		},

		"keepalive": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `BGP keepalive interval for this neighbor

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Keepalive interval in seconds  |

`,
		},

		// Nodes

	}
}
