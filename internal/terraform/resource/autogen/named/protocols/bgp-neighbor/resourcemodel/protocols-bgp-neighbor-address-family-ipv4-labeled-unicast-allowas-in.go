// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAllowasIn describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAllowasIn struct {
	// LeafNodes
	LeafProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAllowasInNumber types.Number `tfsdk:"number" vyos:"number,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAllowasIn) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"number": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Number of occurrences of AS number

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-10  &emsp; |  Number of times AS is allowed in path  |

`,
		},

		// Nodes

	}
}
