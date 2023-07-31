// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsBgpNeighborAddressFamilyIPvsixMulticastDistributeList describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvsixMulticastDistributeList struct {
	// LeafNodes
	LeafProtocolsBgpNeighborAddressFamilyIPvsixMulticastDistributeListExport types.Number `tfsdk:"export" vyos:"export,omitempty"`
	LeafProtocolsBgpNeighborAddressFamilyIPvsixMulticastDistributeListImport types.Number `tfsdk:"import" vyos:"import,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvsixMulticastDistributeList) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"export": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Access-list to filter outgoing route updates to this peer-group

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-65535  |  Access-list to filter outgoing route updates to this peer-group  |

`,
		},

		"import": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Access-list to filter incoming route updates from this peer-group

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-65535  |  Access-list to filter incoming route updates from this peer-group  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBgpNeighborAddressFamilyIPvsixMulticastDistributeList) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsBgpNeighborAddressFamilyIPvsixMulticastDistributeList) UnmarshalJSON(_ []byte) error {
	return nil
}
