/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (l2vpn-evpn) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &ProtocolsBgpNeighborAddressFamilyLtwovpnEvpn{}

// ProtocolsBgpNeighborAddressFamilyLtwovpnEvpn describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type ProtocolsBgpNeighborAddressFamilyLtwovpnEvpn struct {
	// LeafNodes
	LeafProtocolsBgpNeighborAddressFamilyLtwovpnEvpnRouteReflectorClient types.Bool `tfsdk:"route_reflector_client" vyos:"route-reflector-client,omitempty"`
	LeafProtocolsBgpNeighborAddressFamilyLtwovpnEvpnRouteServerClient    types.Bool `tfsdk:"route_server_client" vyos:"route-server-client,omitempty"`

	// TagNodes

	// Nodes

	NodeProtocolsBgpNeighborAddressFamilyLtwovpnEvpnAllowasIn *ProtocolsBgpNeighborAddressFamilyLtwovpnEvpnAllowasIn `tfsdk:"allowas_in" vyos:"allowas-in,omitempty"`

	NodeProtocolsBgpNeighborAddressFamilyLtwovpnEvpnAttributeUnchanged *ProtocolsBgpNeighborAddressFamilyLtwovpnEvpnAttributeUnchanged `tfsdk:"attribute_unchanged" vyos:"attribute-unchanged,omitempty"`

	NodeProtocolsBgpNeighborAddressFamilyLtwovpnEvpnNexthopSelf *ProtocolsBgpNeighborAddressFamilyLtwovpnEvpnNexthopSelf `tfsdk:"nexthop_self" vyos:"nexthop-self,omitempty"`

	NodeProtocolsBgpNeighborAddressFamilyLtwovpnEvpnRouteMap *ProtocolsBgpNeighborAddressFamilyLtwovpnEvpnRouteMap `tfsdk:"route_map" vyos:"route-map,omitempty"`

	NodeProtocolsBgpNeighborAddressFamilyLtwovpnEvpnSoftReconfiguration *ProtocolsBgpNeighborAddressFamilyLtwovpnEvpnSoftReconfiguration `tfsdk:"soft_reconfiguration" vyos:"soft-reconfiguration,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyLtwovpnEvpn) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"route_reflector_client":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (route-reflector-client) */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (route-server-client) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Peer is a route server client

`,
			Description: `Peer is a route server client

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// TagNodes

		// Nodes

		"allowas_in": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyLtwovpnEvpnAllowasIn{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Accept route that contains the local-as in the as-path

`,
			Description: `Accept route that contains the local-as in the as-path

`,
		},

		"attribute_unchanged": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyLtwovpnEvpnAttributeUnchanged{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `BGP attributes are sent unchanged

`,
			Description: `BGP attributes are sent unchanged

`,
		},

		"nexthop_self": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyLtwovpnEvpnNexthopSelf{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Disable the next hop calculation for this peer

`,
			Description: `Disable the next hop calculation for this peer

`,
		},

		"route_map": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyLtwovpnEvpnRouteMap{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Route-map to filter route updates to/from this peer

`,
			Description: `Route-map to filter route updates to/from this peer

`,
		},

		"soft_reconfiguration": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyLtwovpnEvpnSoftReconfiguration{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Soft reconfiguration for peer

`,
			Description: `Soft reconfiguration for peer

`,
		},
	}
}
