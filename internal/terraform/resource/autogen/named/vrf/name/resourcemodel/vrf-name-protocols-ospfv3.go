// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsOspfvthree describes the resource data model.
type VrfNameProtocolsOspfvthree struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfvthreeRouteMap types.String `tfsdk:"route_map" vyos:"route-map,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagVrfNameProtocolsOspfvthreeArea      bool `tfsdk:"area" vyos:"area,child"`
	ExistsTagVrfNameProtocolsOspfvthreeInterface bool `tfsdk:"interface" vyos:"interface,child"`

	// Nodes
	NodeVrfNameProtocolsOspfvthreeAutoCost            *VrfNameProtocolsOspfvthreeAutoCost            `tfsdk:"auto_cost" vyos:"auto-cost,omitempty"`
	NodeVrfNameProtocolsOspfvthreeDefaultInformation  *VrfNameProtocolsOspfvthreeDefaultInformation  `tfsdk:"default_information" vyos:"default-information,omitempty"`
	NodeVrfNameProtocolsOspfvthreeDistance            *VrfNameProtocolsOspfvthreeDistance            `tfsdk:"distance" vyos:"distance,omitempty"`
	NodeVrfNameProtocolsOspfvthreeLogAdjacencyChanges *VrfNameProtocolsOspfvthreeLogAdjacencyChanges `tfsdk:"log_adjacency_changes" vyos:"log-adjacency-changes,omitempty"`
	NodeVrfNameProtocolsOspfvthreeParameters          *VrfNameProtocolsOspfvthreeParameters          `tfsdk:"parameters" vyos:"parameters,omitempty"`
	NodeVrfNameProtocolsOspfvthreeRedistribute        *VrfNameProtocolsOspfvthreeRedistribute        `tfsdk:"redistribute" vyos:"redistribute,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfvthree) ResourceSchemaAttributes() map[string]schema.Attribute {
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

		"auto_cost": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeAutoCost{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Calculate interface cost according to bandwidth

`,
		},

		"default_information": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeDefaultInformation{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Default route advertisment settings

`,
		},

		"distance": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeDistance{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Administrative distance

`,
		},

		"log_adjacency_changes": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeLogAdjacencyChanges{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Log adjacency state changes

`,
		},

		"parameters": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeParameters{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `OSPFv3 specific parameters

`,
		},

		"redistribute": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeRedistribute{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute information from another routing protocol

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsOspfvthree) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsOspfvthree) UnmarshalJSON(_ []byte) error {
	return nil
}
