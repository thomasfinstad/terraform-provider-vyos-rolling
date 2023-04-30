// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// VrfNameProtocolsOspfNeighbor describes the resource data model.
type VrfNameProtocolsOspfNeighbor struct {
	// LeafNodes
	VrfNameProtocolsOspfNeighborPollInterval customtypes.CustomStringValue `tfsdk:"poll_interval" json:"poll-interval,omitempty"`
	VrfNameProtocolsOspfNeighborPriority     customtypes.CustomStringValue `tfsdk:"priority" json:"priority,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VrfNameProtocolsOspfNeighbor) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"poll_interval": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Dead neighbor polling interval

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Seconds between dead neighbor polling interval  |

`,

			// Default:          stringdefault.StaticString(`60`),
			Computed: true,
		},

		"priority": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Neighbor priority in seconds

|  Format  |  Description  |
|----------|---------------|
|  u32:0-255  |  Neighbor priority  |

`,

			// Default:          stringdefault.StaticString(`0`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
