/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &ProtocolsBgpAddressFamilyIPvsixUnicast{}

// ProtocolsBgpAddressFamilyIPvsixUnicast describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type ProtocolsBgpAddressFamilyIPvsixUnicast struct {
	// LeafNodes

	// TagNodes

	ExistsTagProtocolsBgpAddressFamilyIPvsixUnicastAggregateAddress bool `tfsdk:"-" vyos:"aggregate-address,child"`

	ExistsTagProtocolsBgpAddressFamilyIPvsixUnicastNetwork bool `tfsdk:"-" vyos:"network,child"`

	// Nodes

	ExistsNodeProtocolsBgpAddressFamilyIPvsixUnicastDistance bool `tfsdk:"-" vyos:"distance,child"`

	ExistsNodeProtocolsBgpAddressFamilyIPvsixUnicastExport bool `tfsdk:"-" vyos:"export,child"`

	ExistsNodeProtocolsBgpAddressFamilyIPvsixUnicastImport bool `tfsdk:"-" vyos:"import,child"`

	NodeProtocolsBgpAddressFamilyIPvsixUnicastLabel *ProtocolsBgpAddressFamilyIPvsixUnicastLabel `tfsdk:"label" vyos:"label,omitempty"`

	ExistsNodeProtocolsBgpAddressFamilyIPvsixUnicastMaximumPaths bool `tfsdk:"-" vyos:"maximum-paths,child"`

	NodeProtocolsBgpAddressFamilyIPvsixUnicastRd *ProtocolsBgpAddressFamilyIPvsixUnicastRd `tfsdk:"rd" vyos:"rd,omitempty"`

	NodeProtocolsBgpAddressFamilyIPvsixUnicastRouteMap *ProtocolsBgpAddressFamilyIPvsixUnicastRouteMap `tfsdk:"route_map" vyos:"route-map,omitempty"`

	NodeProtocolsBgpAddressFamilyIPvsixUnicastRouteTarget *ProtocolsBgpAddressFamilyIPvsixUnicastRouteTarget `tfsdk:"route_target" vyos:"route-target,omitempty"`

	NodeProtocolsBgpAddressFamilyIPvsixUnicastNexthop *ProtocolsBgpAddressFamilyIPvsixUnicastNexthop `tfsdk:"nexthop" vyos:"nexthop,omitempty"`

	ExistsNodeProtocolsBgpAddressFamilyIPvsixUnicastRedistribute bool `tfsdk:"-" vyos:"redistribute,child"`

	NodeProtocolsBgpAddressFamilyIPvsixUnicastSID *ProtocolsBgpAddressFamilyIPvsixUnicastSID `tfsdk:"sid" vyos:"sid,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpAddressFamilyIPvsixUnicast) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"label": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpAddressFamilyIPvsixUnicastLabel{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Label value for VRF

`,
			Description: `Label value for VRF

`,
		},

		"rd": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpAddressFamilyIPvsixUnicastRd{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Specify route distinguisher

`,
			Description: `Specify route distinguisher

`,
		},

		"route_map": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpAddressFamilyIPvsixUnicastRouteMap{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Route-map to filter route updates to/from this peer

`,
			Description: `Route-map to filter route updates to/from this peer

`,
		},

		"route_target": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpAddressFamilyIPvsixUnicastRouteTarget{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Specify route target list

`,
			Description: `Specify route target list

`,
		},

		"nexthop": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpAddressFamilyIPvsixUnicastNexthop{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Specify next hop to use for VRF advertised prefixes

`,
			Description: `Specify next hop to use for VRF advertised prefixes

`,
		},

		"sid": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpAddressFamilyIPvsixUnicastSID{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `SID value for VRF

`,
			Description: `SID value for VRF

`,
		},
	}
}