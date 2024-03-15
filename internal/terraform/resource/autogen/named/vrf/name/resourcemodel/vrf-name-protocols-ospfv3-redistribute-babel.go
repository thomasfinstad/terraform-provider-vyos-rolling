// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsOspfvthreeRedistributeBabel describes the resource data model.
type VrfNameProtocolsOspfvthreeRedistributeBabel struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfvthreeRedistributeBabelMetric     types.Number `tfsdk:"metric" vyos:"metric,omitempty"`
	LeafVrfNameProtocolsOspfvthreeRedistributeBabelMetricType types.Number `tfsdk:"metric_type" vyos:"metric-type,omitempty"`
	LeafVrfNameProtocolsOspfvthreeRedistributeBabelRouteMap   types.String `tfsdk:"route_map" vyos:"route-map,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfvthreeRedistributeBabel) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"metric": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `OSPF default metric

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-16777214  &emsp; |  Default metric  |

`,
		},

		"metric_type": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `OSPF metric type for default routes

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-2  &emsp; |  Set OSPF External Type 1/2 metrics  |

`,

			// Default:          stringdefault.StaticString(`2`),
			Computed: true,
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
