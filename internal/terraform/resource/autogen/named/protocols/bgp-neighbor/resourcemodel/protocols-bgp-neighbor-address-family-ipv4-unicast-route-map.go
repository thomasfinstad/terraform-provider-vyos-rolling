// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsBgpNeighborAddressFamilyIPvfourUnicastRouteMap describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvfourUnicastRouteMap struct {
	// LeafNodes
	LeafProtocolsBgpNeighborAddressFamilyIPvfourUnicastRouteMapExport types.String `tfsdk:"export" vyos:"export,omitempty"`
	LeafProtocolsBgpNeighborAddressFamilyIPvfourUnicastRouteMapImport types.String `tfsdk:"import" vyos:"import,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvfourUnicastRouteMap) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"export": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route-map to filter outgoing route updates

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  Route map name  |

`,
		},

		"import": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route-map to filter incoming route updates

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  Route map name  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBgpNeighborAddressFamilyIPvfourUnicastRouteMap) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsBgpNeighborAddressFamilyIPvfourUnicastRouteMap) UnmarshalJSON(_ []byte) error {
	return nil
}
