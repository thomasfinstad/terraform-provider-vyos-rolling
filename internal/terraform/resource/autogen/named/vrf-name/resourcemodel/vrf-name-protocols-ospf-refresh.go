// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// VrfNameProtocolsOspfRefresh describes the resource data model.
type VrfNameProtocolsOspfRefresh struct {
	// LeafNodes
	VrfNameProtocolsOspfRefreshTimers customtypes.CustomStringValue `tfsdk:"timers" json:"timers,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VrfNameProtocolsOspfRefresh) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"timers": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Refresh timer

|  Format  |  Description  |
|----------|---------------|
|  u32:10-1800  |  Timer value in seconds  |

`,
		},

		// TagNodes

		// Nodes

	}
}
