// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// VrfNameProtocolsOspfRedistributeBabel describes the resource data model.
type VrfNameProtocolsOspfRedistributeBabel struct {
	// LeafNodes
	VrfNameProtocolsOspfRedistributeBabelMetric     customtypes.CustomStringValue `tfsdk:"metric" json:"metric,omitempty"`
	VrfNameProtocolsOspfRedistributeBabelMetricType customtypes.CustomStringValue `tfsdk:"metric_type" json:"metric-type,omitempty"`
	VrfNameProtocolsOspfRedistributeBabelRouteMap   customtypes.CustomStringValue `tfsdk:"route_map" json:"route-map,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VrfNameProtocolsOspfRedistributeBabel) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"metric": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `OSPF default metric

|  Format  |  Description  |
|----------|---------------|
|  u32:0-16777214  |  Default metric  |

`,
		},

		"metric_type": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `OSPF metric type for default routes

|  Format  |  Description  |
|----------|---------------|
|  u32:1-2  |  Set OSPF External Type 1/2 metrics  |

`,

			// Default:          stringdefault.StaticString(`2`),
			Computed: true,
		},

		"route_map": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Specify route-map name to use

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Route map name  |

`,
		},

		// TagNodes

		// Nodes

	}
}
