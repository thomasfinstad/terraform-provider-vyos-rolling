// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// PolicyRoutesixRule describes the resource data model.
type PolicyRoutesixRule struct {
	// LeafNodes
	PolicyRoutesixRuleAction              customtypes.CustomStringValue `tfsdk:"action" json:"action,omitempty"`
	PolicyRoutesixRuleDescrIPtion         customtypes.CustomStringValue `tfsdk:"description" json:"description,omitempty"`
	PolicyRoutesixRuleDisable             customtypes.CustomStringValue `tfsdk:"disable" json:"disable,omitempty"`
	PolicyRoutesixRuleLog                 customtypes.CustomStringValue `tfsdk:"log" json:"log,omitempty"`
	PolicyRoutesixRuleProtocol            customtypes.CustomStringValue `tfsdk:"protocol" json:"protocol,omitempty"`
	PolicyRoutesixRuleDscp                customtypes.CustomStringValue `tfsdk:"dscp" json:"dscp,omitempty"`
	PolicyRoutesixRuleDscpExclude         customtypes.CustomStringValue `tfsdk:"dscp_exclude" json:"dscp-exclude,omitempty"`
	PolicyRoutesixRulePacketLength        customtypes.CustomStringValue `tfsdk:"packet_length" json:"packet-length,omitempty"`
	PolicyRoutesixRulePacketLengthExclude customtypes.CustomStringValue `tfsdk:"packet_length_exclude" json:"packet-length-exclude,omitempty"`
	PolicyRoutesixRulePacketType          customtypes.CustomStringValue `tfsdk:"packet_type" json:"packet-type,omitempty"`
	PolicyRoutesixRuleConnectionMark      customtypes.CustomStringValue `tfsdk:"connection_mark" json:"connection-mark,omitempty"`

	// TagNodes

	// Nodes
	PolicyRoutesixRuleDestination types.Object `tfsdk:"destination" json:"destination,omitempty"`
	PolicyRoutesixRuleSource      types.Object `tfsdk:"source" json:"source,omitempty"`
	PolicyRoutesixRuleFragment    types.Object `tfsdk:"fragment" json:"fragment,omitempty"`
	PolicyRoutesixRuleIPsec       types.Object `tfsdk:"ipsec" json:"ipsec,omitempty"`
	PolicyRoutesixRuleLimit       types.Object `tfsdk:"limit" json:"limit,omitempty"`
	PolicyRoutesixRuleRecent      types.Object `tfsdk:"recent" json:"recent,omitempty"`
	PolicyRoutesixRuleSet         types.Object `tfsdk:"set" json:"set,omitempty"`
	PolicyRoutesixRuleState       types.Object `tfsdk:"state" json:"state,omitempty"`
	PolicyRoutesixRuleTCP         types.Object `tfsdk:"tcp" json:"tcp,omitempty"`
	PolicyRoutesixRuleTime        types.Object `tfsdk:"time" json:"time,omitempty"`
	PolicyRoutesixRuleIcmpvsix    types.Object `tfsdk:"icmpv6" json:"icmpv6,omitempty"`
	PolicyRoutesixRuleHopLimit    types.Object `tfsdk:"hop_limit" json:"hop-limit,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o PolicyRoutesixRule) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"action": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Rule action

|  Format  |  Description  |
|----------|---------------|
|  accept  |  Accept matching entries  |
|  reject  |  Reject matching entries  |
|  return  |  Return from the current chain and continue at the next rule of the last chain  |
|  drop  |  Drop matching entries  |

`,
		},

		"description": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |

`,
		},

		"disable": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Option to disable firewall rule

`,
		},

		"log": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Option to log packets matching rule

|  Format  |  Description  |
|----------|---------------|
|  enable  |  Enable log  |
|  disable  |  Disable log  |

`,
		},

		"protocol": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Protocol to match (protocol name, number, or "all")

|  Format  |  Description  |
|----------|---------------|
|  all  |  All IP protocols  |
|  tcp_udp  |  Both TCP and UDP  |
|  0-255  |  IP protocol number  |
|  !<protocol>  |  IP protocol number  |

`,

			// Default:          stringdefault.StaticString(`all`),
			Computed: true,
		},

		"dscp": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `DSCP value

|  Format  |  Description  |
|----------|---------------|
|  u32:0-63  |  DSCP value to match  |
|  <start-end>  |  DSCP range to match  |

`,
		},

		"dscp_exclude": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `DSCP value not to match

|  Format  |  Description  |
|----------|---------------|
|  u32:0-63  |  DSCP value not to match  |
|  <start-end>  |  DSCP range not to match  |

`,
		},

		"packet_length": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Payload size in bytes, including header and data to match

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Packet length to match  |
|  <start-end>  |  Packet length range to match  |

`,
		},

		"packet_length_exclude": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Payload size in bytes, including header and data not to match

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Packet length not to match  |
|  <start-end>  |  Packet length range not to match  |

`,
		},

		"packet_type": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Packet type

|  Format  |  Description  |
|----------|---------------|
|  broadcast  |  Match broadcast packet type  |
|  host  |  Match host packet type, addressed to local host  |
|  multicast  |  Match multicast packet type  |
|  other  |  Match packet addressed to another host  |

`,
		},

		"connection_mark": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Connection mark

|  Format  |  Description  |
|----------|---------------|
|  u32:0-2147483647  |  Connection-mark to match  |

`,
		},

		// TagNodes

		// Nodes

		"destination": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleDestination{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Destination parameters

`,
		},

		"source": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleSource{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Source parameters

`,
		},

		"fragment": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleFragment{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `IP fragment match

`,
		},

		"ipsec": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleIPsec{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Inbound IPsec packets

`,
		},

		"limit": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleLimit{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Rate limit using a token bucket filter

`,
		},

		"recent": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleRecent{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Parameters for matching recently seen sources

`,
		},

		"set": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleSet{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Packet modifications

`,
		},

		"state": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleState{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Session state

`,
		},

		"tcp": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleTCP{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `TCP flags to match

`,
		},

		"time": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleTime{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Time to match rule

`,
		},

		"icmpv6": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleIcmpvsix{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `ICMPv6 type and code information

`,
		},

		"hop_limit": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleHopLimit{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Hop limit

`,
		},
	}
}
