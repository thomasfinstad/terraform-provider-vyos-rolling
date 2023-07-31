// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsRIPngDistributeListInterfaceAccessList describes the resource data model.
type ProtocolsRIPngDistributeListInterfaceAccessList struct {
	// LeafNodes
	LeafProtocolsRIPngDistributeListInterfaceAccessListIn  types.Number `tfsdk:"in" vyos:"in,omitempty"`
	LeafProtocolsRIPngDistributeListInterfaceAccessListOut types.Number `tfsdk:"out" vyos:"out,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsRIPngDistributeListInterfaceAccessList) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"in": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Access list to apply to input packets

    |  Format  |  Description  |
    |----------|---------------|
    |  u32  |  Access list to apply to input packets  |

`,
		},

		"out": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Access list to apply to output packets

    |  Format  |  Description  |
    |----------|---------------|
    |  u32  |  Access list to apply to output packets  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsRIPngDistributeListInterfaceAccessList) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsRIPngDistributeListInterfaceAccessList) UnmarshalJSON(_ []byte) error {
	return nil
}
