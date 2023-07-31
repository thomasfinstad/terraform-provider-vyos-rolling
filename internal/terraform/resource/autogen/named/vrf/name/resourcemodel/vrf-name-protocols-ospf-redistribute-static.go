// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsOspfRedistributeStatic describes the resource data model.
type VrfNameProtocolsOspfRedistributeStatic struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfRedistributeStaticMetric     types.Number `tfsdk:"metric" vyos:"metric,omitempty"`
	LeafVrfNameProtocolsOspfRedistributeStaticMetricType types.Number `tfsdk:"metric_type" vyos:"metric-type,omitempty"`
	LeafVrfNameProtocolsOspfRedistributeStaticRouteMap   types.String `tfsdk:"route_map" vyos:"route-map,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfRedistributeStatic) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"metric": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `OSPF default metric

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:0-16777214  |  Default metric  |

`,
		},

		"metric_type": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `OSPF metric type for default routes

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-2  |  Set OSPF External Type 1/2 metrics  |

`,

			// Default:          stringdefault.StaticString(`2`),
			Computed: true,
		},

		"route_map": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Specify route-map name to use

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  Route map name  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsOspfRedistributeStatic) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsOspfRedistributeStatic) UnmarshalJSON(_ []byte) error {
	return nil
}
