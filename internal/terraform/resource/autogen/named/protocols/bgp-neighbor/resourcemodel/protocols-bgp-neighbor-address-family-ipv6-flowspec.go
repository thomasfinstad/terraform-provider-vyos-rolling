/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &ProtocolsBgpNeighborAddressFamilyIPvsixFlowspec{}

// ProtocolsBgpNeighborAddressFamilyIPvsixFlowspec describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvsixFlowspec struct {
	// LeafNodes
	LeafProtocolsBgpNeighborAddressFamilyIPvsixFlowspecRouteReflectorClient types.Bool `tfsdk:"route_reflector_client" vyos:"route-reflector-client,omitempty"`
	LeafProtocolsBgpNeighborAddressFamilyIPvsixFlowspecRouteServerClient    types.Bool `tfsdk:"route_server_client" vyos:"route-server-client,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeProtocolsBgpNeighborAddressFamilyIPvsixFlowspecPrefixList          *ProtocolsBgpNeighborAddressFamilyIPvsixFlowspecPrefixList          `tfsdk:"prefix_list" vyos:"prefix-list,omitempty"`
	NodeProtocolsBgpNeighborAddressFamilyIPvsixFlowspecFilterList          *ProtocolsBgpNeighborAddressFamilyIPvsixFlowspecFilterList          `tfsdk:"filter_list" vyos:"filter-list,omitempty"`
	NodeProtocolsBgpNeighborAddressFamilyIPvsixFlowspecRouteMap            *ProtocolsBgpNeighborAddressFamilyIPvsixFlowspecRouteMap            `tfsdk:"route_map" vyos:"route-map,omitempty"`
	NodeProtocolsBgpNeighborAddressFamilyIPvsixFlowspecSoftReconfiguration *ProtocolsBgpNeighborAddressFamilyIPvsixFlowspecSoftReconfiguration `tfsdk:"soft_reconfiguration" vyos:"soft-reconfiguration,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvsixFlowspec) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"route_reflector_client":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Peer is a route reflector client

`,
			Description: `Peer is a route reflector client

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"route_server_client":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Peer is a route server client

`,
			Description: `Peer is a route server client

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

		"prefix_list": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvsixFlowspecPrefixList{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Prefix-list to filter route updates to/from this peer

`,
			Description: `Prefix-list to filter route updates to/from this peer

`,
		},

		"filter_list": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvsixFlowspecFilterList{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `as-path-list to filter route updates to/from this peer

`,
			Description: `as-path-list to filter route updates to/from this peer

`,
		},

		"route_map": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvsixFlowspecRouteMap{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Route-map to filter route updates to/from this peer

`,
			Description: `Route-map to filter route updates to/from this peer

`,
		},

		"soft_reconfiguration": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvsixFlowspecSoftReconfiguration{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Soft reconfiguration for peer

`,
			Description: `Soft reconfiguration for peer

`,
		},
	}
}
