// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeRIP describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeRIP struct {
	// LeafNodes
	VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeRIPMetric   customtypes.CustomStringValue `tfsdk:"metric" json:"metric,omitempty"`
	VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeRIPRouteMap customtypes.CustomStringValue `tfsdk:"route_map" json:"route-map,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeRIP) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"metric": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Metric for redistributed routes

|  Format  |  Description  |
|----------|---------------|
|  u32:1-4294967295  |  Metric for redistributed routes  |

`,
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
