// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsBgpNeighborAddressFamilyIPvfourMulticastDefaultOriginate describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvfourMulticastDefaultOriginate struct {
	// LeafNodes
	LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastDefaultOriginateRouteMap types.String `tfsdk:"route_map" vyos:"route-map,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvfourMulticastDefaultOriginate) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"route_map": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Specify route-map name to use

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Route map name  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBgpNeighborAddressFamilyIPvfourMulticastDefaultOriginate) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsBgpNeighborAddressFamilyIPvfourMulticastDefaultOriginate) UnmarshalJSON(_ []byte) error {
	return nil
}
