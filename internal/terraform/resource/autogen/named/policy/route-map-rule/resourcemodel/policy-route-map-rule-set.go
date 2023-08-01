// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// PolicyRouteMapRuleSet describes the resource data model.
type PolicyRouteMapRuleSet struct {
	// LeafNodes
	LeafPolicyRouteMapRuleSetAtomicAggregate types.Bool   `tfsdk:"atomic_aggregate" vyos:"atomic-aggregate,omitempty"`
	LeafPolicyRouteMapRuleSetDistance        types.Number `tfsdk:"distance" vyos:"distance,omitempty"`
	LeafPolicyRouteMapRuleSetIPNextHop       types.String `tfsdk:"ip_next_hop" vyos:"ip-next-hop,omitempty"`
	LeafPolicyRouteMapRuleSetLocalPreference types.Number `tfsdk:"local_preference" vyos:"local-preference,omitempty"`
	LeafPolicyRouteMapRuleSetMetric          types.String `tfsdk:"metric" vyos:"metric,omitempty"`
	LeafPolicyRouteMapRuleSetMetricType      types.String `tfsdk:"metric_type" vyos:"metric-type,omitempty"`
	LeafPolicyRouteMapRuleSetOrigin          types.String `tfsdk:"origin" vyos:"origin,omitempty"`
	LeafPolicyRouteMapRuleSetOriginatorID    types.String `tfsdk:"originator_id" vyos:"originator-id,omitempty"`
	LeafPolicyRouteMapRuleSetSrc             types.String `tfsdk:"src" vyos:"src,omitempty"`
	LeafPolicyRouteMapRuleSetTable           types.Number `tfsdk:"table" vyos:"table,omitempty"`
	LeafPolicyRouteMapRuleSetTag             types.Number `tfsdk:"tag" vyos:"tag,omitempty"`
	LeafPolicyRouteMapRuleSetWeight          types.Number `tfsdk:"weight" vyos:"weight,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodePolicyRouteMapRuleSetAggregator       *PolicyRouteMapRuleSetAggregator       `tfsdk:"aggregator" vyos:"aggregator,omitempty"`
	NodePolicyRouteMapRuleSetAsPath           *PolicyRouteMapRuleSetAsPath           `tfsdk:"as_path" vyos:"as-path,omitempty"`
	NodePolicyRouteMapRuleSetCommunity        *PolicyRouteMapRuleSetCommunity        `tfsdk:"community" vyos:"community,omitempty"`
	NodePolicyRouteMapRuleSetLargeCommunity   *PolicyRouteMapRuleSetLargeCommunity   `tfsdk:"large_community" vyos:"large-community,omitempty"`
	NodePolicyRouteMapRuleSetExtcommunity     *PolicyRouteMapRuleSetExtcommunity     `tfsdk:"extcommunity" vyos:"extcommunity,omitempty"`
	NodePolicyRouteMapRuleSetEvpn             *PolicyRouteMapRuleSetEvpn             `tfsdk:"evpn" vyos:"evpn,omitempty"`
	NodePolicyRouteMapRuleSetIPvsixNextHop    *PolicyRouteMapRuleSetIPvsixNextHop    `tfsdk:"ipv6_next_hop" vyos:"ipv6-next-hop,omitempty"`
	NodePolicyRouteMapRuleSetLthreevpnNexthop *PolicyRouteMapRuleSetLthreevpnNexthop `tfsdk:"l3vpn_nexthop" vyos:"l3vpn-nexthop,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteMapRuleSet) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"atomic_aggregate": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `BGP atomic aggregate attribute

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"distance": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Locally significant administrative distance

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-255  &emsp; |  Distance value  |

`,
		},

		"ip_next_hop": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Nexthop IP address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IP address  |
    |  unchanged  &emsp; |  Set the BGP nexthop address as unchanged  |
    |  peer-address  &emsp; |  Set the BGP nexthop address to the address of the peer  |

`,
		},

		"local_preference": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `BGP local preference attribute

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-4294967295  &emsp; |  Local preference value  |

`,
		},

		"metric": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Destination routing protocol metric

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  <+/-metric>  &emsp; |  Add or subtract metric  |
    |  number: 0-4294967295  &emsp; |  Metric value  |

`,
		},

		"metric_type": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Open Shortest Path First (OSPF) external metric-type

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  type-1  &emsp; |  OSPF external type 1 metric  |
    |  type-2  &emsp; |  OSPF external type 2 metric  |

`,
		},

		"origin": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Border Gateway Protocl (BGP) origin code

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  igp  &emsp; |  Interior gateway protocol origin  |
    |  egp  &emsp; |  Exterior gateway protocol origin  |
    |  incomplete  &emsp; |  Incomplete origin  |

`,
		},

		"originator_id": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `BGP originator ID attribute

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  Orignator IP address  |

`,
		},

		"src": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Source address for route

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IPv4 address  |
    |  ipv6  &emsp; |  IPv6 address  |

`,
		},

		"table": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set prefixes to table

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-200  &emsp; |  Table value  |

`,
		},

		"tag": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Tag value for routing protocol

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Tag value  |

`,
		},

		"weight": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `BGP weight attribute

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-4294967295  &emsp; |  BGP weight  |

`,
		},

		// Nodes

		"aggregator": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleSetAggregator{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `BGP aggregator attribute

`,
		},

		"as_path": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleSetAsPath{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Transform BGP AS_PATH attribute

`,
		},

		"community": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleSetCommunity{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `BGP community attribute

`,
		},

		"large_community": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleSetLargeCommunity{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `BGP large community attribute

`,
		},

		"extcommunity": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleSetExtcommunity{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `BGP extended community attribute

`,
		},

		"evpn": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleSetEvpn{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Ethernet Virtual Private Network

`,
		},

		"ipv6_next_hop": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleSetIPvsixNextHop{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Nexthop IPv6 address

`,
		},

		"l3vpn_nexthop": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleSetLthreevpnNexthop{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Next hop Information

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PolicyRouteMapRuleSet) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *PolicyRouteMapRuleSet) UnmarshalJSON(_ []byte) error {
	return nil
}
