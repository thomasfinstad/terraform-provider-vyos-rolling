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

var _ helpers.VyosResourceDataModel = &ProtocolsBgpAddressFamilyIPvfourUnicast{}

// ProtocolsBgpAddressFamilyIPvfourUnicast describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type ProtocolsBgpAddressFamilyIPvfourUnicast struct {
	// LeafNodes

	// TagNodes

	ExistsTagProtocolsBgpAddressFamilyIPvfourUnicastAggregateAddress bool `tfsdk:"-" vyos:"aggregate-address,child"`

	ExistsTagProtocolsBgpAddressFamilyIPvfourUnicastNetwork bool `tfsdk:"-" vyos:"network,child"`

	// Nodes

	ExistsNodeProtocolsBgpAddressFamilyIPvfourUnicastDistance bool `tfsdk:"-" vyos:"distance,child"`

	ExistsNodeProtocolsBgpAddressFamilyIPvfourUnicastExport bool `tfsdk:"-" vyos:"export,child"`

	ExistsNodeProtocolsBgpAddressFamilyIPvfourUnicastImport bool `tfsdk:"-" vyos:"import,child"`

	NodeProtocolsBgpAddressFamilyIPvfourUnicastLabel *ProtocolsBgpAddressFamilyIPvfourUnicastLabel `tfsdk:"label" vyos:"label,omitempty"`

	ExistsNodeProtocolsBgpAddressFamilyIPvfourUnicastMaximumPaths bool `tfsdk:"-" vyos:"maximum-paths,child"`

	NodeProtocolsBgpAddressFamilyIPvfourUnicastRd *ProtocolsBgpAddressFamilyIPvfourUnicastRd `tfsdk:"rd" vyos:"rd,omitempty"`

	NodeProtocolsBgpAddressFamilyIPvfourUnicastRouteMap *ProtocolsBgpAddressFamilyIPvfourUnicastRouteMap `tfsdk:"route_map" vyos:"route-map,omitempty"`

	NodeProtocolsBgpAddressFamilyIPvfourUnicastRouteTarget *ProtocolsBgpAddressFamilyIPvfourUnicastRouteTarget `tfsdk:"route_target" vyos:"route-target,omitempty"`

	NodeProtocolsBgpAddressFamilyIPvfourUnicastNexthop *ProtocolsBgpAddressFamilyIPvfourUnicastNexthop `tfsdk:"nexthop" vyos:"nexthop,omitempty"`

	ExistsNodeProtocolsBgpAddressFamilyIPvfourUnicastRedistribute bool `tfsdk:"-" vyos:"redistribute,child"`

	NodeProtocolsBgpAddressFamilyIPvfourUnicastSID *ProtocolsBgpAddressFamilyIPvfourUnicastSID `tfsdk:"sid" vyos:"sid,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpAddressFamilyIPvfourUnicast) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"label": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpAddressFamilyIPvfourUnicastLabel{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Label value for VRF

`,
			Description: `Label value for VRF

`,
		},

		"rd": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpAddressFamilyIPvfourUnicastRd{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Specify route distinguisher

`,
			Description: `Specify route distinguisher

`,
		},

		"route_map": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpAddressFamilyIPvfourUnicastRouteMap{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Route-map to filter route updates to/from this peer

`,
			Description: `Route-map to filter route updates to/from this peer

`,
		},

		"route_target": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpAddressFamilyIPvfourUnicastRouteTarget{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Specify route target list

`,
			Description: `Specify route target list

`,
		},

		"nexthop": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpAddressFamilyIPvfourUnicastNexthop{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Specify next hop to use for VRF advertised prefixes

`,
			Description: `Specify next hop to use for VRF advertised prefixes

`,
		},

		"sid": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpAddressFamilyIPvfourUnicastSID{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `SID value for VRF

`,
			Description: `SID value for VRF

`,
		},
	}
}
