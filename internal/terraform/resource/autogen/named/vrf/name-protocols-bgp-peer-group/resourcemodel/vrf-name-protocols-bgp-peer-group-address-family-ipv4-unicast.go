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

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicast{}

// VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicast describes the resource data model.
type VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicast struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAddpathTxAll         types.Bool   `tfsdk:"addpath_tx_all" vyos:"addpath-tx-all,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAddpathTxPerAs       types.Bool   `tfsdk:"addpath_tx_per_as" vyos:"addpath-tx-per-as,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAsOverrIDe           types.Bool   `tfsdk:"as_override" vyos:"as-override,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastMaximumPrefix        types.Number `tfsdk:"maximum_prefix" vyos:"maximum-prefix,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastMaximumPrefixOut     types.Number `tfsdk:"maximum_prefix_out" vyos:"maximum-prefix-out,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteReflectorClient types.Bool   `tfsdk:"route_reflector_client" vyos:"route-reflector-client,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteServerClient    types.Bool   `tfsdk:"route_server_client" vyos:"route-server-client,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastUnsuppressMap        types.String `tfsdk:"unsuppress_map" vyos:"unsuppress-map,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastWeight               types.Number `tfsdk:"weight" vyos:"weight,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapability             *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapability             `tfsdk:"capability" vyos:"capability,omitempty"`
	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastPrefixList             *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastPrefixList             `tfsdk:"prefix_list" vyos:"prefix-list,omitempty"`
	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastConditionallyAdvertise *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastConditionallyAdvertise `tfsdk:"conditionally_advertise" vyos:"conditionally-advertise,omitempty"`
	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAllowasIn              *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAllowasIn              `tfsdk:"allowas_in" vyos:"allowas-in,omitempty"`
	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchanged     *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchanged     `tfsdk:"attribute_unchanged" vyos:"attribute-unchanged,omitempty"`
	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDisableSendCommunity   *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDisableSendCommunity   `tfsdk:"disable_send_community" vyos:"disable-send-community,omitempty"`
	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDistributeList         *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDistributeList         `tfsdk:"distribute_list" vyos:"distribute-list,omitempty"`
	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastFilterList             *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastFilterList             `tfsdk:"filter_list" vyos:"filter-list,omitempty"`
	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastNexthopSelf            *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastNexthopSelf            `tfsdk:"nexthop_self" vyos:"nexthop-self,omitempty"`
	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRemovePrivateAs        *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRemovePrivateAs        `tfsdk:"remove_private_as" vyos:"remove-private-as,omitempty"`
	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteMap               *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteMap               `tfsdk:"route_map" vyos:"route-map,omitempty"`
	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastSoftReconfiguration    *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastSoftReconfiguration    `tfsdk:"soft_reconfiguration" vyos:"soft-reconfiguration,omitempty"`
	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDefaultOriginate       *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDefaultOriginate       `tfsdk:"default_originate" vyos:"default-originate,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicast) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
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
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapability{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Advertise capabilities to this neighbor (IPv4)

`,
			Description: `Advertise capabilities to this neighbor (IPv4)

`,
		},

		"prefix_list": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastPrefixList{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv4-Prefix-list to filter route updates to/from this peer

`,
			Description: `IPv4-Prefix-list to filter route updates to/from this peer

`,
		},

		"conditionally_advertise": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastConditionallyAdvertise{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Use route-map to conditionally advertise routes

`,
			Description: `Use route-map to conditionally advertise routes

`,
		},

		"allowas_in": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAllowasIn{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Accept route that contains the local-as in the as-path

`,
			Description: `Accept route that contains the local-as in the as-path

`,
		},

		"attribute_unchanged": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchanged{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `BGP attributes are sent unchanged

`,
			Description: `BGP attributes are sent unchanged

`,
		},

		"disable_send_community": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDisableSendCommunity{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Disable sending community attributes to this peer

`,
			Description: `Disable sending community attributes to this peer

`,
		},

		"distribute_list": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDistributeList{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Access-list to filter route updates to/from this peer-group

`,
			Description: `Access-list to filter route updates to/from this peer-group

`,
		},

		"filter_list": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastFilterList{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `as-path-list to filter route updates to/from this peer

`,
			Description: `as-path-list to filter route updates to/from this peer

`,
		},

		"nexthop_self": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastNexthopSelf{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Disable the next hop calculation for this peer

`,
			Description: `Disable the next hop calculation for this peer

`,
		},

		"remove_private_as": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRemovePrivateAs{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Remove private AS numbers from AS path in outbound route updates

`,
			Description: `Remove private AS numbers from AS path in outbound route updates

`,
		},

		"route_map": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteMap{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Route-map to filter route updates to/from this peer

`,
			Description: `Route-map to filter route updates to/from this peer

`,
		},

		"soft_reconfiguration": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastSoftReconfiguration{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Soft reconfiguration for peer

`,
			Description: `Soft reconfiguration for peer

`,
		},

		"default_originate": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDefaultOriginate{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Originate default route to this peer

`,
			Description: `Originate default route to this peer

`,
		},
	}
}
