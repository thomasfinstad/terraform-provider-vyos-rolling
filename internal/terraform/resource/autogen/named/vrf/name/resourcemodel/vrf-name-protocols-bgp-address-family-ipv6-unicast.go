// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// VrfNameProtocolsBgpAddressFamilyIPvsixUnicast describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvsixUnicast struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagVrfNameProtocolsBgpAddressFamilyIPvsixUnicastAggregateAddress bool `tfsdk:"aggregate_address" vyos:"aggregate-address,child"`
	ExistsTagVrfNameProtocolsBgpAddressFamilyIPvsixUnicastNetwork          bool `tfsdk:"network" vyos:"network,child"`

	// Nodes
	NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastDistance     *VrfNameProtocolsBgpAddressFamilyIPvsixUnicastDistance     `tfsdk:"distance" vyos:"distance,omitempty"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastExport       *VrfNameProtocolsBgpAddressFamilyIPvsixUnicastExport       `tfsdk:"export" vyos:"export,omitempty"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastImport       *VrfNameProtocolsBgpAddressFamilyIPvsixUnicastImport       `tfsdk:"import" vyos:"import,omitempty"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastLabel        *VrfNameProtocolsBgpAddressFamilyIPvsixUnicastLabel        `tfsdk:"label" vyos:"label,omitempty"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastMaximumPaths *VrfNameProtocolsBgpAddressFamilyIPvsixUnicastMaximumPaths `tfsdk:"maximum_paths" vyos:"maximum-paths,omitempty"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRd           *VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRd           `tfsdk:"rd" vyos:"rd,omitempty"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRouteMap     *VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRouteMap     `tfsdk:"route_map" vyos:"route-map,omitempty"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRouteTarget  *VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRouteTarget  `tfsdk:"route_target" vyos:"route-target,omitempty"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastNexthop      *VrfNameProtocolsBgpAddressFamilyIPvsixUnicastNexthop      `tfsdk:"nexthop" vyos:"nexthop,omitempty"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistribute *VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistribute `tfsdk:"redistribute" vyos:"redistribute,omitempty"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastSID          *VrfNameProtocolsBgpAddressFamilyIPvsixUnicastSID          `tfsdk:"sid" vyos:"sid,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvsixUnicast) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"distance": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvsixUnicastDistance{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Administrative distances for BGP routes

`,
		},

		"export": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvsixUnicastExport{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Export routes from this address-family

`,
		},

		"import": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvsixUnicastImport{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Import routes to this address-family

`,
		},

		"label": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvsixUnicastLabel{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Label value for VRF

`,
		},

		"maximum_paths": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvsixUnicastMaximumPaths{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Forward packets over multiple paths

`,
		},

		"rd": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRd{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Specify route distinguisher

`,
		},

		"route_map": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRouteMap{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Route-map to filter route updates to/from this peer

`,
		},

		"route_target": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRouteTarget{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Specify route target list

`,
		},

		"nexthop": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvsixUnicastNexthop{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Specify next hop to use for VRF advertised prefixes

`,
		},

		"redistribute": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRedistribute{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute routes from other protocols into BGP

`,
		},

		"sid": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvsixUnicastSID{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `SID value for VRF

`,
		},
	}
}
