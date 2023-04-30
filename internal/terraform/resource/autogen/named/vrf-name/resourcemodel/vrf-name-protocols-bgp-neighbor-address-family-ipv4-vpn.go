// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpn describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpn struct {
	// LeafNodes
	VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAddpathTxAll         customtypes.CustomStringValue `tfsdk:"addpath_tx_all" json:"addpath-tx-all,omitempty"`
	VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAddpathTxPerAs       customtypes.CustomStringValue `tfsdk:"addpath_tx_per_as" json:"addpath-tx-per-as,omitempty"`
	VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAsOverrIDe           customtypes.CustomStringValue `tfsdk:"as_override" json:"as-override,omitempty"`
	VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnMaximumPrefix        customtypes.CustomStringValue `tfsdk:"maximum_prefix" json:"maximum-prefix,omitempty"`
	VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnMaximumPrefixOut     customtypes.CustomStringValue `tfsdk:"maximum_prefix_out" json:"maximum-prefix-out,omitempty"`
	VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnRemovePrivateAs      customtypes.CustomStringValue `tfsdk:"remove_private_as" json:"remove-private-as,omitempty"`
	VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteReflectorClient customtypes.CustomStringValue `tfsdk:"route_reflector_client" json:"route-reflector-client,omitempty"`
	VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteServerClient    customtypes.CustomStringValue `tfsdk:"route_server_client" json:"route-server-client,omitempty"`
	VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnUnsuppressMap        customtypes.CustomStringValue `tfsdk:"unsuppress_map" json:"unsuppress-map,omitempty"`
	VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnWeight               customtypes.CustomStringValue `tfsdk:"weight" json:"weight,omitempty"`

	// TagNodes

	// Nodes
	VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnPrefixList             types.Object `tfsdk:"prefix_list" json:"prefix-list,omitempty"`
	VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertise types.Object `tfsdk:"conditionally_advertise" json:"conditionally-advertise,omitempty"`
	VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAllowasIn              types.Object `tfsdk:"allowas_in" json:"allowas-in,omitempty"`
	VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchanged     types.Object `tfsdk:"attribute_unchanged" json:"attribute-unchanged,omitempty"`
	VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunity   types.Object `tfsdk:"disable_send_community" json:"disable-send-community,omitempty"`
	VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnDistributeList         types.Object `tfsdk:"distribute_list" json:"distribute-list,omitempty"`
	VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnFilterList             types.Object `tfsdk:"filter_list" json:"filter-list,omitempty"`
	VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnNexthopSelf            types.Object `tfsdk:"nexthop_self" json:"nexthop-self,omitempty"`
	VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteMap               types.Object `tfsdk:"route_map" json:"route-map,omitempty"`
	VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnSoftReconfiguration    types.Object `tfsdk:"soft_reconfiguration" json:"soft-reconfiguration,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpn) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"addpath_tx_all": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Use addpath to advertise all paths to a neighbor

`,
		},

		"addpath_tx_per_as": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Use addpath to advertise the bestpath per each neighboring AS

`,
		},

		"as_override": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Override ASN in outbound updates to configured neighbor local-as

`,
		},

		"maximum_prefix": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Maximum number of prefixes to accept from this peer

|  Format  |  Description  |
|----------|---------------|
|  u32:1-4294967295  |  Prefix limit  |

`,
		},

		"maximum_prefix_out": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Maximum number of prefixes to be sent to this peer

|  Format  |  Description  |
|----------|---------------|
|  u32:1-4294967295  |  Prefix limit  |

`,
		},

		"remove_private_as": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Remove private AS numbers from AS path in outbound route updates

`,
		},

		"route_reflector_client": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Peer is a route reflector client

`,
		},

		"route_server_client": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Peer is a route server client

`,
		},

		"unsuppress_map": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Route-map to selectively unsuppress suppressed routes

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Route map name  |

`,
		},

		"weight": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Default weight for routes from this peer

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Default weight  |

`,
		},

		// TagNodes

		// Nodes

		"prefix_list": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnPrefixList{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4-Prefix-list to filter route updates to/from this peer

`,
		},

		"conditionally_advertise": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertise{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Use route-map to conditionally advertise routes

`,
		},

		"allowas_in": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAllowasIn{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Accept route that contains the local-as in the as-path

`,
		},

		"attribute_unchanged": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchanged{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `BGP attributes are sent unchanged

`,
		},

		"disable_send_community": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunity{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Disable sending community attributes to this peer

`,
		},

		"distribute_list": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnDistributeList{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Access-list to filter route updates to/from this peer-group

`,
		},

		"filter_list": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnFilterList{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `as-path-list to filter route updates to/from this peer

`,
		},

		"nexthop_self": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnNexthopSelf{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Disable the next hop calculation for this peer

`,
		},

		"route_map": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteMap{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Route-map to filter route updates to/from this peer

`,
		},

		"soft_reconfiguration": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnSoftReconfiguration{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Soft reconfiguration for peer

`,
		},
	}
}
