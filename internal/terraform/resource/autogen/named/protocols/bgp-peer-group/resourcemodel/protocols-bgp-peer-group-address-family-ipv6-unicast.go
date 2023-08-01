// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicast describes the resource data model.
type ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicast struct {
	// LeafNodes
	LeafProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastAddpathTxAll         types.Bool   `tfsdk:"addpath_tx_all" vyos:"addpath-tx-all,omitempty"`
	LeafProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastAddpathTxPerAs       types.Bool   `tfsdk:"addpath_tx_per_as" vyos:"addpath-tx-per-as,omitempty"`
	LeafProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastAsOverrIDe           types.Bool   `tfsdk:"as_override" vyos:"as-override,omitempty"`
	LeafProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastMaximumPrefix        types.Number `tfsdk:"maximum_prefix" vyos:"maximum-prefix,omitempty"`
	LeafProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastMaximumPrefixOut     types.Number `tfsdk:"maximum_prefix_out" vyos:"maximum-prefix-out,omitempty"`
	LeafProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastRemovePrivateAs      types.Bool   `tfsdk:"remove_private_as" vyos:"remove-private-as,omitempty"`
	LeafProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastRouteReflectorClient types.Bool   `tfsdk:"route_reflector_client" vyos:"route-reflector-client,omitempty"`
	LeafProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastRouteServerClient    types.Bool   `tfsdk:"route_server_client" vyos:"route-server-client,omitempty"`
	LeafProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastUnsuppressMap        types.String `tfsdk:"unsuppress_map" vyos:"unsuppress-map,omitempty"`
	LeafProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastWeight               types.Number `tfsdk:"weight" vyos:"weight,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastCapability             *ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastCapability             `tfsdk:"capability" vyos:"capability,omitempty"`
	NodeProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastNexthopLocal           *ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastNexthopLocal           `tfsdk:"nexthop_local" vyos:"nexthop-local,omitempty"`
	NodeProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastPrefixList             *ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastPrefixList             `tfsdk:"prefix_list" vyos:"prefix-list,omitempty"`
	NodeProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastConditionallyAdvertise *ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastConditionallyAdvertise `tfsdk:"conditionally_advertise" vyos:"conditionally-advertise,omitempty"`
	NodeProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastAllowasIn              *ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastAllowasIn              `tfsdk:"allowas_in" vyos:"allowas-in,omitempty"`
	NodeProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastAttributeUnchanged     *ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastAttributeUnchanged     `tfsdk:"attribute_unchanged" vyos:"attribute-unchanged,omitempty"`
	NodeProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDisableSendCommunity   *ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDisableSendCommunity   `tfsdk:"disable_send_community" vyos:"disable-send-community,omitempty"`
	NodeProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDistributeList         *ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDistributeList         `tfsdk:"distribute_list" vyos:"distribute-list,omitempty"`
	NodeProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastFilterList             *ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastFilterList             `tfsdk:"filter_list" vyos:"filter-list,omitempty"`
	NodeProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastNexthopSelf            *ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastNexthopSelf            `tfsdk:"nexthop_self" vyos:"nexthop-self,omitempty"`
	NodeProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastRouteMap               *ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastRouteMap               `tfsdk:"route_map" vyos:"route-map,omitempty"`
	NodeProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastSoftReconfiguration    *ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastSoftReconfiguration    `tfsdk:"soft_reconfiguration" vyos:"soft-reconfiguration,omitempty"`
	NodeProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDefaultOriginate       *ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDefaultOriginate       `tfsdk:"default_originate" vyos:"default-originate,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicast) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"addpath_tx_all": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Use addpath to advertise all paths to a neighbor

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"addpath_tx_per_as": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Use addpath to advertise the bestpath per each neighboring AS

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"as_override": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Override ASN in outbound updates to configured neighbor local-as

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"maximum_prefix": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum number of prefixes to accept from this peer

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-4294967295  &emsp; |  Prefix limit  |

`,
		},

		"maximum_prefix_out": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum number of prefixes to be sent to this peer

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-4294967295  &emsp; |  Prefix limit  |

`,
		},

		"remove_private_as": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Remove private AS numbers from AS path in outbound route updates

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"route_reflector_client": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Peer is a route reflector client

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"route_server_client": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Peer is a route server client

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"unsuppress_map": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route-map to selectively unsuppress suppressed routes

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Route map name  |

`,
		},

		"weight": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Default weight for routes from this peer

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Default weight  |

`,
		},

		// Nodes

		"capability": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastCapability{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Advertise capabilities to this neighbor (IPv6)

`,
		},

		"nexthop_local": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastNexthopLocal{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Nexthop attributes

`,
		},

		"prefix_list": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastPrefixList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Prefix-list to filter route updates to/from this peer

`,
		},

		"conditionally_advertise": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastConditionallyAdvertise{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Use route-map to conditionally advertise routes

`,
		},

		"allowas_in": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastAllowasIn{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Accept route that contains the local-as in the as-path

`,
		},

		"attribute_unchanged": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastAttributeUnchanged{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `BGP attributes are sent unchanged

`,
		},

		"disable_send_community": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDisableSendCommunity{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Disable sending community attributes to this peer

`,
		},

		"distribute_list": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDistributeList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Access-list to filter route updates to/from this peer-group

`,
		},

		"filter_list": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastFilterList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `as-path-list to filter route updates to/from this peer

`,
		},

		"nexthop_self": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastNexthopSelf{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Disable the next hop calculation for this peer

`,
		},

		"route_map": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastRouteMap{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Route-map to filter route updates to/from this peer

`,
		},

		"soft_reconfiguration": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastSoftReconfiguration{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Soft reconfiguration for peer

`,
		},

		"default_originate": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDefaultOriginate{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Originate default route to this peer

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicast) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicast) UnmarshalJSON(_ []byte) error {
	return nil
}
