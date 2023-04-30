// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ProtocolsBgpNeighborAddressFamilyIPvsixUnicastAllowasIn describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvsixUnicastAllowasIn struct {
	// LeafNodes
	ProtocolsBgpNeighborAddressFamilyIPvsixUnicastAllowasInNumber customtypes.CustomStringValue `tfsdk:"number" json:"number,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvsixUnicastAllowasIn) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"number": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Number of occurrences of AS number

|  Format  |  Description  |
|----------|---------------|
|  u32:1-10  |  Number of times AS is allowed in path  |

`,
		},

		// TagNodes

		// Nodes

	}
}
