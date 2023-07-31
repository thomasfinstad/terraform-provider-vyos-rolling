// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsRIPngDistributeListInterfacePrefixList describes the resource data model.
type ProtocolsRIPngDistributeListInterfacePrefixList struct {
	// LeafNodes
	LeafProtocolsRIPngDistributeListInterfacePrefixListIn  types.String `tfsdk:"in" vyos:"in,omitempty"`
	LeafProtocolsRIPngDistributeListInterfacePrefixListOut types.String `tfsdk:"out" vyos:"out,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsRIPngDistributeListInterfacePrefixList) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"in": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Prefix-list to apply to input packets

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  Prefix-list to apply to input packets  |

`,
		},

		"out": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Prefix-list to apply to output packets

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  Prefix-list to apply to output packets  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsRIPngDistributeListInterfacePrefixList) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsRIPngDistributeListInterfacePrefixList) UnmarshalJSON(_ []byte) error {
	return nil
}
