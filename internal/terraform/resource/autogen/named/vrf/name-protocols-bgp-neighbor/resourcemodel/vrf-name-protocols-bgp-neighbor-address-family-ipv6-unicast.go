// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicast describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicast struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastAddpathTxAll         types.Bool   `tfsdk:"addpath_tx_all" vyos:"addpath-tx-all,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastAddpathTxPerAs       types.Bool   `tfsdk:"addpath_tx_per_as" vyos:"addpath-tx-per-as,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastAsOverrIDe           types.Bool   `tfsdk:"as_override" vyos:"as-override,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastMaximumPrefix        types.Number `tfsdk:"maximum_prefix" vyos:"maximum-prefix,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastMaximumPrefixOut     types.Number `tfsdk:"maximum_prefix_out" vyos:"maximum-prefix-out,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastRemovePrivateAs      types.Bool   `tfsdk:"remove_private_as" vyos:"remove-private-as,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastRouteReflectorClient types.Bool   `tfsdk:"route_reflector_client" vyos:"route-reflector-client,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastRouteServerClient    types.Bool   `tfsdk:"route_server_client" vyos:"route-server-client,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastUnsuppressMap        types.String `tfsdk:"unsuppress_map" vyos:"unsuppress-map,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastWeight               types.Number `tfsdk:"weight" vyos:"weight,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastCapability             *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastCapability             `tfsdk:"capability" vyos:"capability,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastNexthopLocal           *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastNexthopLocal           `tfsdk:"nexthop_local" vyos:"nexthop-local,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastPrefixList             *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastPrefixList             `tfsdk:"prefix_list" vyos:"prefix-list,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastConditionallyAdvertise *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastConditionallyAdvertise `tfsdk:"conditionally_advertise" vyos:"conditionally-advertise,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastAllowasIn              *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastAllowasIn              `tfsdk:"allowas_in" vyos:"allowas-in,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastAttributeUnchanged     *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastAttributeUnchanged     `tfsdk:"attribute_unchanged" vyos:"attribute-unchanged,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastDisableSendCommunity   *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastDisableSendCommunity   `tfsdk:"disable_send_community" vyos:"disable-send-community,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastDistributeList         *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastDistributeList         `tfsdk:"distribute_list" vyos:"distribute-list,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastFilterList             *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastFilterList             `tfsdk:"filter_list" vyos:"filter-list,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastNexthopSelf            *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastNexthopSelf            `tfsdk:"nexthop_self" vyos:"nexthop-self,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastRouteMap               *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastRouteMap               `tfsdk:"route_map" vyos:"route-map,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastSoftReconfiguration    *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastSoftReconfiguration    `tfsdk:"soft_reconfiguration" vyos:"soft-reconfiguration,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastDefaultOriginate       *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastDefaultOriginate       `tfsdk:"default_originate" vyos:"default-originate,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicast) ResourceSchemaAttributes() map[string]schema.Attribute {
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
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastCapability{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Advertise capabilities to this neighbor (IPv6)

`,
		},

		"nexthop_local": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastNexthopLocal{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Nexthop attributes

`,
		},

		"prefix_list": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastPrefixList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Prefix-list to filter route updates to/from this peer

`,
		},

		"conditionally_advertise": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastConditionallyAdvertise{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Use route-map to conditionally advertise routes

`,
		},

		"allowas_in": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastAllowasIn{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Accept route that contains the local-as in the as-path

`,
		},

		"attribute_unchanged": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastAttributeUnchanged{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `BGP attributes are sent unchanged

`,
		},

		"disable_send_community": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastDisableSendCommunity{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Disable sending community attributes to this peer

`,
		},

		"distribute_list": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastDistributeList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Access-list to filter route updates to/from this peer-group

`,
		},

		"filter_list": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastFilterList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `as-path-list to filter route updates to/from this peer

`,
		},

		"nexthop_self": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastNexthopSelf{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Disable the next hop calculation for this peer

`,
		},

		"route_map": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastRouteMap{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Route-map to filter route updates to/from this peer

`,
		},

		"soft_reconfiguration": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastSoftReconfiguration{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Soft reconfiguration for peer

`,
		},

		"default_originate": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastDefaultOriginate{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Originate default route to this peer

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicast) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicast) UnmarshalJSON(_ []byte) error {
	return nil
}
