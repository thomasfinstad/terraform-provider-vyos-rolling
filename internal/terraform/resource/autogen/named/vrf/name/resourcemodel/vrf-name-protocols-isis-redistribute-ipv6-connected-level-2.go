// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsIsisRedistributeIPvsixConnectedLevelTwo describes the resource data model.
type VrfNameProtocolsIsisRedistributeIPvsixConnectedLevelTwo struct {
	// LeafNodes
	LeafVrfNameProtocolsIsisRedistributeIPvsixConnectedLevelTwoMetric   types.Number `tfsdk:"metric" vyos:"metric,omitempty"`
	LeafVrfNameProtocolsIsisRedistributeIPvsixConnectedLevelTwoRouteMap types.String `tfsdk:"route_map" vyos:"route-map,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisRedistributeIPvsixConnectedLevelTwo) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"metric": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set default metric for circuit

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-16777215  &emsp; |  Default metric value  |

`,
		},

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
func (o *VrfNameProtocolsIsisRedistributeIPvsixConnectedLevelTwo) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsIsisRedistributeIPvsixConnectedLevelTwo) UnmarshalJSON(_ []byte) error {
	return nil
}
