// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsOspfvthreeDistanceOspfvthree describes the resource data model.
type VrfNameProtocolsOspfvthreeDistanceOspfvthree struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfvthreeDistanceOspfvthreeExternal  types.Number `tfsdk:"external" vyos:"external,omitempty"`
	LeafVrfNameProtocolsOspfvthreeDistanceOspfvthreeInterArea types.Number `tfsdk:"inter_area" vyos:"inter-area,omitempty"`
	LeafVrfNameProtocolsOspfvthreeDistanceOspfvthreeIntraArea types.Number `tfsdk:"intra_area" vyos:"intra-area,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfvthreeDistanceOspfvthree) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"external": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Distance for external routes

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-255  |  Distance for external routes  |

`,
		},

		"inter_area": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Distance for inter-area routes

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-255  |  Distance for inter-area routes  |

`,
		},

		"intra_area": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Distance for intra-area routes

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-255  |  Distance for intra-area routes  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsOspfvthreeDistanceOspfvthree) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsOspfvthreeDistanceOspfvthree) UnmarshalJSON(_ []byte) error {
	return nil
}
