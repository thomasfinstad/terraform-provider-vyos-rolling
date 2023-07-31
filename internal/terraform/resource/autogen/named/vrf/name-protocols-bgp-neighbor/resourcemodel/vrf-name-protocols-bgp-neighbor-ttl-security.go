// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsBgpNeighborTTLSecURIty describes the resource data model.
type VrfNameProtocolsBgpNeighborTTLSecURIty struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborTTLSecURItyHops types.Number `tfsdk:"hops" vyos:"hops,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborTTLSecURIty) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"hops": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Number of the maximum number of hops to the BGP peer

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-254  |  Number of hops  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsBgpNeighborTTLSecURIty) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsBgpNeighborTTLSecURIty) UnmarshalJSON(_ []byte) error {
	return nil
}
