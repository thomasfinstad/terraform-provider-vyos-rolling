// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ServicePppoeServerPadoDelay describes the resource data model.
type ServicePppoeServerPadoDelay struct {
	// LeafNodes
	ServicePppoeServerPadoDelaySessions customtypes.CustomStringValue `tfsdk:"sessions" json:"sessions,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ServicePppoeServerPadoDelay) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"sessions": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Number of sessions

|  Format  |  Description  |
|----------|---------------|
|  u32:1-999999  |  Number of sessions  |

`,
		},

		// TagNodes

		// Nodes

	}
}
