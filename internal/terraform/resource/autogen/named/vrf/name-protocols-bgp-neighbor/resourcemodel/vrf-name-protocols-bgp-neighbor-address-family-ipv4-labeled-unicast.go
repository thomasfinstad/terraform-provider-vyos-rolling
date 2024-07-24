// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicast{}

// VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicast describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicast struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAddpathTxAll         types.Bool   `tfsdk:"addpath_tx_all" vyos:"addpath-tx-all,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAddpathTxPerAs       types.Bool   `tfsdk:"addpath_tx_per_as" vyos:"addpath-tx-per-as,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAsOverrIDe           types.Bool   `tfsdk:"as_override" vyos:"as-override,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastMaximumPrefix        types.Number `tfsdk:"maximum_prefix" vyos:"maximum-prefix,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastMaximumPrefixOut     types.Number `tfsdk:"maximum_prefix_out" vyos:"maximum-prefix-out,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastRouteReflectorClient types.Bool   `tfsdk:"route_reflector_client" vyos:"route-reflector-client,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastRouteServerClient    types.Bool   `tfsdk:"route_server_client" vyos:"route-server-client,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastUnsuppressMap        types.String `tfsdk:"unsuppress_map" vyos:"unsuppress-map,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastWeight               types.Number `tfsdk:"weight" vyos:"weight,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastCapability             *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastCapability             `tfsdk:"capability" vyos:"capability,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastPrefixList             *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastPrefixList             `tfsdk:"prefix_list" vyos:"prefix-list,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastConditionallyAdvertise *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastConditionallyAdvertise `tfsdk:"conditionally_advertise" vyos:"conditionally-advertise,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAllowasIn              *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAllowasIn              `tfsdk:"allowas_in" vyos:"allowas-in,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAttributeUnchanged     *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAttributeUnchanged     `tfsdk:"attribute_unchanged" vyos:"attribute-unchanged,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDisableSendCommunity   *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDisableSendCommunity   `tfsdk:"disable_send_community" vyos:"disable-send-community,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDistributeList         *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDistributeList         `tfsdk:"distribute_list" vyos:"distribute-list,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastFilterList             *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastFilterList             `tfsdk:"filter_list" vyos:"filter-list,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastNexthopSelf            *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastNexthopSelf            `tfsdk:"nexthop_self" vyos:"nexthop-self,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastRemovePrivateAs        *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastRemovePrivateAs        `tfsdk:"remove_private_as" vyos:"remove-private-as,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastRouteMap               *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastRouteMap               `tfsdk:"route_map" vyos:"route-map,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastSoftReconfiguration    *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastSoftReconfiguration    `tfsdk:"soft_reconfiguration" vyos:"soft-reconfiguration,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDefaultOriginate       *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDefaultOriginate       `tfsdk:"default_originate" vyos:"default-originate,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicast) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"addpath_tx_all": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Use addpath to advertise all paths to a neighbor

`,
			Description: `Use addpath to advertise all paths to a neighbor

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"addpath_tx_per_as": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Use addpath to advertise the bestpath per each neighboring AS

`,
			Description: `Use addpath to advertise the bestpath per each neighboring AS

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"as_override": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Override ASN in outbound updates to configured neighbor local-as

`,
			Description: `Override ASN in outbound updates to configured neighbor local-as

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"maximum_prefix": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum number of prefixes to accept from this peer

    |  Format        |  Description   |
    |----------------|----------------|
    |  1-4294967295  |  Prefix limit  |
`,
			Description: `Maximum number of prefixes to accept from this peer

    |  Format        |  Description   |
    |----------------|----------------|
    |  1-4294967295  |  Prefix limit  |
`,
		},

		"maximum_prefix_out": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum number of prefixes to be sent to this peer

    |  Format        |  Description   |
    |----------------|----------------|
    |  1-4294967295  |  Prefix limit  |
`,
			Description: `Maximum number of prefixes to be sent to this peer

    |  Format        |  Description   |
    |----------------|----------------|
    |  1-4294967295  |  Prefix limit  |
`,
		},

		"route_reflector_client": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Peer is a route reflector client

`,
			Description: `Peer is a route reflector client

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"route_server_client": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Peer is a route server client

`,
			Description: `Peer is a route server client

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"unsuppress_map": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route-map to selectively unsuppress suppressed routes

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Route map name  |
`,
			Description: `Route-map to selectively unsuppress suppressed routes

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Route map name  |
`,
		},

		"weight": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Default weight for routes from this peer

    |  Format   |  Description     |
    |-----------|------------------|
    |  1-65535  |  Default weight  |
`,
			Description: `Default weight for routes from this peer

    |  Format   |  Description     |
    |-----------|------------------|
    |  1-65535  |  Default weight  |
`,
		},

		// Nodes

		"capability": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastCapability{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Advertise capabilities to this neighbor (IPv4)

`,
			Description: `Advertise capabilities to this neighbor (IPv4)

`,
		},

		"prefix_list": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastPrefixList{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv4-Prefix-list to filter route updates to/from this peer

`,
			Description: `IPv4-Prefix-list to filter route updates to/from this peer

`,
		},

		"conditionally_advertise": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastConditionallyAdvertise{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Use route-map to conditionally advertise routes

`,
			Description: `Use route-map to conditionally advertise routes

`,
		},

		"allowas_in": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAllowasIn{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Accept route that contains the local-as in the as-path

`,
			Description: `Accept route that contains the local-as in the as-path

`,
		},

		"attribute_unchanged": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAttributeUnchanged{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `BGP attributes are sent unchanged

`,
			Description: `BGP attributes are sent unchanged

`,
		},

		"disable_send_community": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDisableSendCommunity{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Disable sending community attributes to this peer

`,
			Description: `Disable sending community attributes to this peer

`,
		},

		"distribute_list": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDistributeList{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Access-list to filter route updates to/from this peer-group

`,
			Description: `Access-list to filter route updates to/from this peer-group

`,
		},

		"filter_list": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastFilterList{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `as-path-list to filter route updates to/from this peer

`,
			Description: `as-path-list to filter route updates to/from this peer

`,
		},

		"nexthop_self": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastNexthopSelf{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Disable the next hop calculation for this peer

`,
			Description: `Disable the next hop calculation for this peer

`,
		},

		"remove_private_as": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastRemovePrivateAs{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Remove private AS numbers from AS path in outbound route updates

`,
			Description: `Remove private AS numbers from AS path in outbound route updates

`,
		},

		"route_map": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastRouteMap{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Route-map to filter route updates to/from this peer

`,
			Description: `Route-map to filter route updates to/from this peer

`,
		},

		"soft_reconfiguration": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastSoftReconfiguration{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Soft reconfiguration for peer

`,
			Description: `Soft reconfiguration for peer

`,
		},

		"default_originate": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDefaultOriginate{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Originate default route to this peer

`,
			Description: `Originate default route to this peer

`,
		},
	}
}
