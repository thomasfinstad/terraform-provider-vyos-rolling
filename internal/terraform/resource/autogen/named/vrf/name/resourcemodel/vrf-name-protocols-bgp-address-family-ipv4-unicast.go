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

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpAddressFamilyIPvfourUnicast{}

// VrfNameProtocolsBgpAddressFamilyIPvfourUnicast describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VrfNameProtocolsBgpAddressFamilyIPvfourUnicast struct {
	// LeafNodes

	// TagNodes

	ExistsTagVrfNameProtocolsBgpAddressFamilyIPvfourUnicastAggregateAddress bool `tfsdk:"-" vyos:"aggregate-address,child"`

	ExistsTagVrfNameProtocolsBgpAddressFamilyIPvfourUnicastNetwork bool `tfsdk:"-" vyos:"network,child"`

	// Nodes

	NodeVrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistance *VrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistance `tfsdk:"distance" vyos:"distance,omitempty"`

	NodeVrfNameProtocolsBgpAddressFamilyIPvfourUnicastExport *VrfNameProtocolsBgpAddressFamilyIPvfourUnicastExport `tfsdk:"export" vyos:"export,omitempty"`

	NodeVrfNameProtocolsBgpAddressFamilyIPvfourUnicastImport *VrfNameProtocolsBgpAddressFamilyIPvfourUnicastImport `tfsdk:"import" vyos:"import,omitempty"`

	NodeVrfNameProtocolsBgpAddressFamilyIPvfourUnicastLabel *VrfNameProtocolsBgpAddressFamilyIPvfourUnicastLabel `tfsdk:"label" vyos:"label,omitempty"`

	NodeVrfNameProtocolsBgpAddressFamilyIPvfourUnicastMaximumPaths *VrfNameProtocolsBgpAddressFamilyIPvfourUnicastMaximumPaths `tfsdk:"maximum_paths" vyos:"maximum-paths,omitempty"`

	NodeVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRd *VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRd `tfsdk:"rd" vyos:"rd,omitempty"`

	NodeVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteMap *VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteMap `tfsdk:"route_map" vyos:"route-map,omitempty"`

	NodeVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTarget *VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTarget `tfsdk:"route_target" vyos:"route-target,omitempty"`

	NodeVrfNameProtocolsBgpAddressFamilyIPvfourUnicastNexthop *VrfNameProtocolsBgpAddressFamilyIPvfourUnicastNexthop `tfsdk:"nexthop" vyos:"nexthop,omitempty"`

	NodeVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistribute *VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistribute `tfsdk:"redistribute" vyos:"redistribute,omitempty"`

	NodeVrfNameProtocolsBgpAddressFamilyIPvfourUnicastSID *VrfNameProtocolsBgpAddressFamilyIPvfourUnicastSID `tfsdk:"sid" vyos:"sid,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourUnicast) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"distance": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistance{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Administrative distances for BGP routes

`,
			Description: `Administrative distances for BGP routes

`,
		},

		"export": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicastExport{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Export routes from this address-family

`,
			Description: `Export routes from this address-family

`,
		},

		"import": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicastImport{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Import routes to this address-family

`,
			Description: `Import routes to this address-family

`,
		},

		"label": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicastLabel{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Label value for VRF

`,
			Description: `Label value for VRF

`,
		},

		"maximum_paths": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicastMaximumPaths{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Forward packets over multiple paths

`,
			Description: `Forward packets over multiple paths

`,
		},

		"rd": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRd{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Specify route distinguisher

`,
			Description: `Specify route distinguisher

`,
		},

		"route_map": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteMap{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Route-map to filter route updates to/from this peer

`,
			Description: `Route-map to filter route updates to/from this peer

`,
		},

		"route_target": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTarget{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Specify route target list

`,
			Description: `Specify route target list

`,
		},

		"nexthop": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicastNexthop{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Specify next hop to use for VRF advertised prefixes

`,
			Description: `Specify next hop to use for VRF advertised prefixes

`,
		},

		"redistribute": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistribute{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Redistribute routes from other protocols into BGP

`,
			Description: `Redistribute routes from other protocols into BGP

`,
		},

		"sid": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicastSID{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `SID value for VRF

`,
			Description: `SID value for VRF

`,
		},
	}
}
