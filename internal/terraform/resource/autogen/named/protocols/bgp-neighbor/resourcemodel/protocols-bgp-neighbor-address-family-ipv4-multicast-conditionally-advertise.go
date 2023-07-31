// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsBgpNeighborAddressFamilyIPvfourMulticastConditionallyAdvertise describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvfourMulticastConditionallyAdvertise struct {
	// LeafNodes
	LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastConditionallyAdvertiseAdvertiseMap types.String `tfsdk:"advertise_map" vyos:"advertise-map,omitempty"`
	LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastConditionallyAdvertiseExistMap     types.String `tfsdk:"exist_map" vyos:"exist-map,omitempty"`
	LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastConditionallyAdvertiseNonExistMap  types.String `tfsdk:"non_exist_map" vyos:"non-exist-map,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvfourMulticastConditionallyAdvertise) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"advertise_map": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route-map to conditionally advertise routes

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  Route map name  |

`,
		},

		"exist_map": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Advertise routes only if prefixes in exist-map are installed in BGP table

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  Route map name  |

`,
		},

		"non_exist_map": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Advertise routes only if prefixes in non-exist-map are not installed in BGP table

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  Route map name  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBgpNeighborAddressFamilyIPvfourMulticastConditionallyAdvertise) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsBgpNeighborAddressFamilyIPvfourMulticastConditionallyAdvertise) UnmarshalJSON(_ []byte) error {
	return nil
}
