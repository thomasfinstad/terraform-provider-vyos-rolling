// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/numberplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &FirewallIPvsixInputFilterRule{}

// FirewallIPvsixInputFilterRule describes the resource data model.
type FirewallIPvsixInputFilterRule struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Number `tfsdk:"rule_id" vyos:"-,self-id"`

	// LeafNodes
	LeafFirewallIPvsixInputFilterRuleAction              types.String `tfsdk:"action" vyos:"action,omitempty"`
	LeafFirewallIPvsixInputFilterRuleDescrIPtion         types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafFirewallIPvsixInputFilterRuleDscp                types.List   `tfsdk:"dscp" vyos:"dscp,omitempty"`
	LeafFirewallIPvsixInputFilterRuleDscpExclude         types.List   `tfsdk:"dscp_exclude" vyos:"dscp-exclude,omitempty"`
	LeafFirewallIPvsixInputFilterRulePacketLength        types.List   `tfsdk:"packet_length" vyos:"packet-length,omitempty"`
	LeafFirewallIPvsixInputFilterRulePacketLengthExclude types.List   `tfsdk:"packet_length_exclude" vyos:"packet-length-exclude,omitempty"`
	LeafFirewallIPvsixInputFilterRulePacketType          types.String `tfsdk:"packet_type" vyos:"packet-type,omitempty"`
	LeafFirewallIPvsixInputFilterRuleMark                types.String `tfsdk:"mark" vyos:"mark,omitempty"`
	LeafFirewallIPvsixInputFilterRuleConnectionMark      types.List   `tfsdk:"connection_mark" vyos:"connection-mark,omitempty"`
	LeafFirewallIPvsixInputFilterRuleConntrackHelper     types.List   `tfsdk:"conntrack_helper" vyos:"conntrack-helper,omitempty"`
	LeafFirewallIPvsixInputFilterRuleQueue               types.Number `tfsdk:"queue" vyos:"queue,omitempty"`
	LeafFirewallIPvsixInputFilterRuleQueueOptions        types.List   `tfsdk:"queue_options" vyos:"queue-options,omitempty"`
	LeafFirewallIPvsixInputFilterRuleDisable             types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafFirewallIPvsixInputFilterRuleLog                 types.Bool   `tfsdk:"log" vyos:"log,omitempty"`
	LeafFirewallIPvsixInputFilterRuleProtocol            types.String `tfsdk:"protocol" vyos:"protocol,omitempty"`
	LeafFirewallIPvsixInputFilterRuleState               types.List   `tfsdk:"state" vyos:"state,omitempty"`
	LeafFirewallIPvsixInputFilterRuleJumpTarget          types.String `tfsdk:"jump_target" vyos:"jump-target,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeFirewallIPvsixInputFilterRuleFragment          *FirewallIPvsixInputFilterRuleFragment          `tfsdk:"fragment" vyos:"fragment,omitempty"`
	NodeFirewallIPvsixInputFilterRuleLimit             *FirewallIPvsixInputFilterRuleLimit             `tfsdk:"limit" vyos:"limit,omitempty"`
	NodeFirewallIPvsixInputFilterRuleLogOptions        *FirewallIPvsixInputFilterRuleLogOptions        `tfsdk:"log_options" vyos:"log-options,omitempty"`
	NodeFirewallIPvsixInputFilterRuleConnectionStatus  *FirewallIPvsixInputFilterRuleConnectionStatus  `tfsdk:"connection_status" vyos:"connection-status,omitempty"`
	NodeFirewallIPvsixInputFilterRuleRecent            *FirewallIPvsixInputFilterRuleRecent            `tfsdk:"recent" vyos:"recent,omitempty"`
	NodeFirewallIPvsixInputFilterRuleSynproxy          *FirewallIPvsixInputFilterRuleSynproxy          `tfsdk:"synproxy" vyos:"synproxy,omitempty"`
	NodeFirewallIPvsixInputFilterRuleTCP               *FirewallIPvsixInputFilterRuleTCP               `tfsdk:"tcp" vyos:"tcp,omitempty"`
	NodeFirewallIPvsixInputFilterRuleTime              *FirewallIPvsixInputFilterRuleTime              `tfsdk:"time" vyos:"time,omitempty"`
	NodeFirewallIPvsixInputFilterRuleHopLimit          *FirewallIPvsixInputFilterRuleHopLimit          `tfsdk:"hop_limit" vyos:"hop-limit,omitempty"`
	NodeFirewallIPvsixInputFilterRuleAddAddressToGroup *FirewallIPvsixInputFilterRuleAddAddressToGroup `tfsdk:"add_address_to_group" vyos:"add-address-to-group,omitempty"`
	NodeFirewallIPvsixInputFilterRuleDestination       *FirewallIPvsixInputFilterRuleDestination       `tfsdk:"destination" vyos:"destination,omitempty"`
	NodeFirewallIPvsixInputFilterRuleIcmpvsix          *FirewallIPvsixInputFilterRuleIcmpvsix          `tfsdk:"icmpv6" vyos:"icmpv6,omitempty"`
	NodeFirewallIPvsixInputFilterRuleSource            *FirewallIPvsixInputFilterRuleSource            `tfsdk:"source" vyos:"source,omitempty"`
	NodeFirewallIPvsixInputFilterRuleInboundInterface  *FirewallIPvsixInputFilterRuleInboundInterface  `tfsdk:"inbound_interface" vyos:"inbound-interface,omitempty"`
	NodeFirewallIPvsixInputFilterRuleIPsec             *FirewallIPvsixInputFilterRuleIPsec             `tfsdk:"ipsec" vyos:"ipsec,omitempty"`
}

// SetID configures the resource ID
func (o *FirewallIPvsixInputFilterRule) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *FirewallIPvsixInputFilterRule) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"firewall",

		"ipv6",

		"input",

		"filter",

		"rule",
		o.SelfIdentifier.ValueBigFloat().String(),
	}
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *FirewallIPvsixInputFilterRule) GetVyosParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixInputFilterRule) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"rule_id": schema.NumberAttribute{
			Required: true,
			MarkdownDescription: `IPv6 Firewall input filter rule number

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-999999  &emsp; |  Number for this firewall rule  |

`,
			PlanModifiers: []planmodifier.Number{
				numberplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"action": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Rule action

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  accept  &emsp; |  Accept matching entries  |
    |  continue  &emsp; |  Continue parsing next rule  |
    |  jump  &emsp; |  Jump to another chain  |
    |  reject  &emsp; |  Reject matching entries  |
    |  return  &emsp; |  Return from the current chain and continue at the next rule of the last chain  |
    |  drop  &emsp; |  Drop matching entries  |
    |  queue  &emsp; |  Enqueue packet to userspace  |
    |  offload  &emsp; |  Offload packet via flowtable  |
    |  synproxy  &emsp; |  Synproxy connections  |

`,
		},

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Description  |

`,
		},

		"dscp": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `DSCP value

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-63  &emsp; |  DSCP value to match  |
    |  <start-end>  &emsp; |  DSCP range to match  |

`,
		},

		"dscp_exclude": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `DSCP value not to match

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-63  &emsp; |  DSCP value not to match  |
    |  <start-end>  &emsp; |  DSCP range not to match  |

`,
		},

		"packet_length": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Payload size in bytes, including header and data to match

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Packet length to match  |
    |  <start-end>  &emsp; |  Packet length range to match  |

`,
		},

		"packet_length_exclude": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Payload size in bytes, including header and data not to match

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Packet length not to match  |
    |  <start-end>  &emsp; |  Packet length range not to match  |

`,
		},

		"packet_type": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Packet type

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  broadcast  &emsp; |  Match broadcast packet type  |
    |  host  &emsp; |  Match host packet type, addressed to local host  |
    |  multicast  &emsp; |  Match multicast packet type  |
    |  other  &emsp; |  Match packet addressed to another host  |

`,
		},

		"mark": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Firewall mark

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-2147483647  &emsp; |  Firewall mark to match  |
    |  !number: 0-2147483647  &emsp; |  Inverted Firewall mark to match  |
    |  <start-end>  &emsp; |  Firewall mark range to match  |
    |  !<start-end>  &emsp; |  Firewall mark inverted range to match  |

`,
		},

		"connection_mark": schema.ListAttribute{
			ElementType: types.NumberType,
			Optional:    true,
			MarkdownDescription: `Connection mark

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-2147483647  &emsp; |  Connection-mark to match  |

`,
		},

		"conntrack_helper": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Match related traffic from conntrack helpers

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ftp  &emsp; |  Related traffic from FTP helper  |
    |  h323  &emsp; |  Related traffic from H.323 helper  |
    |  pptp  &emsp; |  Related traffic from PPTP helper  |
    |  nfs  &emsp; |  Related traffic from NFS helper  |
    |  sip  &emsp; |  Related traffic from SIP helper  |
    |  tftp  &emsp; |  Related traffic from TFTP helper  |
    |  sqlnet  &emsp; |  Related traffic from SQLNet helper  |

`,
		},

		"queue": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Queue target to use. Action queue must be defined to use this setting

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-65535  &emsp; |  Queue target  |

`,
		},

		"queue_options": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Options used for queue target. Action queue must be defined to use this setting

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  bypass  &emsp; |  Let packets go through if userspace application cannot back off  |
    |  fanout  &emsp; |  Distribute packets between several queues  |

`,
		},

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"log": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Log packets hitting this rule

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"protocol": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Protocol to match (protocol name, number, or "all")

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  all  &emsp; |  All IP protocols  |
    |  tcp_udp  &emsp; |  Both TCP and UDP  |
    |  number: 0-255  &emsp; |  IP protocol number  |
    |  <protocol>  &emsp; |  IP protocol name  |
    |  !<protocol>  &emsp; |  IP protocol name  |

`,
		},

		"state": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Session state

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  established  &emsp; |  Established state  |
    |  invalid  &emsp; |  Invalid state  |
    |  new  &emsp; |  New state  |
    |  related  &emsp; |  Related state  |

`,
		},

		"jump_target": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set jump target. Action jump must be defined to use this setting

`,
		},

		// Nodes

		"fragment": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixInputFilterRuleFragment{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IP fragment match

`,
		},

		"limit": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixInputFilterRuleLimit{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Rate limit using a token bucket filter

`,
		},

		"log_options": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixInputFilterRuleLogOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Log options

`,
		},

		"connection_status": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixInputFilterRuleConnectionStatus{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Connection status

`,
		},

		"recent": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixInputFilterRuleRecent{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Parameters for matching recently seen sources

`,
		},

		"synproxy": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixInputFilterRuleSynproxy{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Synproxy options

`,
		},

		"tcp": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixInputFilterRuleTCP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `TCP options to match

`,
		},

		"time": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixInputFilterRuleTime{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Time to match rule

`,
		},

		"hop_limit": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixInputFilterRuleHopLimit{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Hop limit

`,
		},

		"add_address_to_group": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixInputFilterRuleAddAddressToGroup{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Add ipv6 address to dynamic ipv6-address-group

`,
		},

		"destination": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixInputFilterRuleDestination{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Destination parameters

`,
		},

		"icmpv6": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixInputFilterRuleIcmpvsix{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `ICMPv6 type and code information

`,
		},

		"source": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixInputFilterRuleSource{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Source parameters

`,
		},

		"inbound_interface": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixInputFilterRuleInboundInterface{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match inbound-interface

`,
		},

		"ipsec": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixInputFilterRuleIPsec{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Inbound IPsec packets

`,
		},
	}
}
