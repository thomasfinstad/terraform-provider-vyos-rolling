// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// FirewallNameRule describes the resource data model.
type FirewallNameRule struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.Number `tfsdk:"rule_id" vyos:",self-id"`

	ParentIDFirewallName types.String `tfsdk:"name" vyos:"name,parent-id"`

	// LeafNodes
	LeafFirewallNameRuleAction              types.String `tfsdk:"action" vyos:"action,omitempty"`
	LeafFirewallNameRuleDescrIPtion         types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafFirewallNameRuleDisable             types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafFirewallNameRuleLog                 types.String `tfsdk:"log" vyos:"log,omitempty"`
	LeafFirewallNameRuleLogLevel            types.String `tfsdk:"log_level" vyos:"log-level,omitempty"`
	LeafFirewallNameRuleProtocol            types.String `tfsdk:"protocol" vyos:"protocol,omitempty"`
	LeafFirewallNameRuleDscp                types.List   `tfsdk:"dscp" vyos:"dscp,omitempty"`
	LeafFirewallNameRuleDscpExclude         types.List   `tfsdk:"dscp_exclude" vyos:"dscp-exclude,omitempty"`
	LeafFirewallNameRulePacketLength        types.List   `tfsdk:"packet_length" vyos:"packet-length,omitempty"`
	LeafFirewallNameRulePacketLengthExclude types.List   `tfsdk:"packet_length_exclude" vyos:"packet-length-exclude,omitempty"`
	LeafFirewallNameRulePacketType          types.String `tfsdk:"packet_type" vyos:"packet-type,omitempty"`
	LeafFirewallNameRuleConnectionMark      types.List   `tfsdk:"connection_mark" vyos:"connection-mark,omitempty"`
	LeafFirewallNameRuleJumpTarget          types.String `tfsdk:"jump_target" vyos:"jump-target,omitempty"`
	LeafFirewallNameRuleQueue               types.Number `tfsdk:"queue" vyos:"queue,omitempty"`
	LeafFirewallNameRuleQueueOptions        types.List   `tfsdk:"queue_options" vyos:"queue-options,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeFirewallNameRuleDestination       *FirewallNameRuleDestination       `tfsdk:"destination" vyos:"destination,omitempty"`
	NodeFirewallNameRuleSource            *FirewallNameRuleSource            `tfsdk:"source" vyos:"source,omitempty"`
	NodeFirewallNameRuleFragment          *FirewallNameRuleFragment          `tfsdk:"fragment" vyos:"fragment,omitempty"`
	NodeFirewallNameRuleInboundInterface  *FirewallNameRuleInboundInterface  `tfsdk:"inbound_interface" vyos:"inbound-interface,omitempty"`
	NodeFirewallNameRuleOutboundInterface *FirewallNameRuleOutboundInterface `tfsdk:"outbound_interface" vyos:"outbound-interface,omitempty"`
	NodeFirewallNameRuleIPsec             *FirewallNameRuleIPsec             `tfsdk:"ipsec" vyos:"ipsec,omitempty"`
	NodeFirewallNameRuleLimit             *FirewallNameRuleLimit             `tfsdk:"limit" vyos:"limit,omitempty"`
	NodeFirewallNameRuleConnectionStatus  *FirewallNameRuleConnectionStatus  `tfsdk:"connection_status" vyos:"connection-status,omitempty"`
	NodeFirewallNameRuleRecent            *FirewallNameRuleRecent            `tfsdk:"recent" vyos:"recent,omitempty"`
	NodeFirewallNameRuleState             *FirewallNameRuleState             `tfsdk:"state" vyos:"state,omitempty"`
	NodeFirewallNameRuleTCP               *FirewallNameRuleTCP               `tfsdk:"tcp" vyos:"tcp,omitempty"`
	NodeFirewallNameRuleTime              *FirewallNameRuleTime              `tfsdk:"time" vyos:"time,omitempty"`
	NodeFirewallNameRuleIcmp              *FirewallNameRuleIcmp              `tfsdk:"icmp" vyos:"icmp,omitempty"`
	NodeFirewallNameRuleTTL               *FirewallNameRuleTTL               `tfsdk:"ttl" vyos:"ttl,omitempty"`
}

// GetID returns the resource ID
func (o FirewallNameRule) GetID() *types.String {
	return &o.ID
}

// SetID configures the resource ID
func (o FirewallNameRule) SetID(id types.String) {
	o.ID = id
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *FirewallNameRule) GetVyosPath() []string {
	return []string{
		"firewall",

		"name",
		o.ParentIDFirewallName.ValueString(),

		"rule",
		o.SelfIdentifier.ValueBigFloat().String(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallNameRule) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"rule_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Firewall rule number (IPv4)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-999999  &emsp; |  Number for this Firewall rule  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		"name_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `IPv4 firewall rule-set name

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"action": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Rule action

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  accept  &emsp; |  Accept matching entries  |
    |  jump  &emsp; |  Jump to another chain  |
    |  reject  &emsp; |  Reject matching entries  |
    |  return  &emsp; |  Return from the current chain and continue at the next rule of the last chain  |
    |  drop  &emsp; |  Drop matching entries  |
    |  queue  &emsp; |  Enqueue packet to userspace  |

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

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Option to disable firewall rule

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"log": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Option to log packets matching rule

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  enable  &emsp; |  Enable log  |
    |  disable  &emsp; |  Disable log  |

`,
		},

		"log_level": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set log-level. Log must be enable.

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  emerg  &emsp; |  Emerg log level  |
    |  alert  &emsp; |  Alert log level  |
    |  crit  &emsp; |  Critical log level  |
    |  err  &emsp; |  Error log level  |
    |  warn  &emsp; |  Warning log level  |
    |  notice  &emsp; |  Notice log level  |
    |  info  &emsp; |  Info log level  |
    |  debug  &emsp; |  Debug log level  |

`,
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

		"connection_mark": schema.ListAttribute{
			ElementType: types.NumberType,
			Optional:    true,
			MarkdownDescription: `Connection mark

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-2147483647  &emsp; |  Connection-mark to match  |

`,
		},

		"jump_target": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set jump target. Action jump must be defined to use this setting

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

		// Nodes

		"destination": schema.SingleNestedAttribute{
			Attributes: FirewallNameRuleDestination{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Destination parameters

`,
		},

		"source": schema.SingleNestedAttribute{
			Attributes: FirewallNameRuleSource{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Source parameters

`,
		},

		"fragment": schema.SingleNestedAttribute{
			Attributes: FirewallNameRuleFragment{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IP fragment match

`,
		},

		"inbound_interface": schema.SingleNestedAttribute{
			Attributes: FirewallNameRuleInboundInterface{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match inbound-interface

`,
		},

		"outbound_interface": schema.SingleNestedAttribute{
			Attributes: FirewallNameRuleOutboundInterface{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match outbound-interface

`,
		},

		"ipsec": schema.SingleNestedAttribute{
			Attributes: FirewallNameRuleIPsec{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Inbound IPsec packets

`,
		},

		"limit": schema.SingleNestedAttribute{
			Attributes: FirewallNameRuleLimit{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Rate limit using a token bucket filter

`,
		},

		"connection_status": schema.SingleNestedAttribute{
			Attributes: FirewallNameRuleConnectionStatus{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Connection status

`,
		},

		"recent": schema.SingleNestedAttribute{
			Attributes: FirewallNameRuleRecent{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Parameters for matching recently seen sources

`,
		},

		"state": schema.SingleNestedAttribute{
			Attributes: FirewallNameRuleState{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Session state

`,
		},

		"tcp": schema.SingleNestedAttribute{
			Attributes: FirewallNameRuleTCP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `TCP flags to match

`,
		},

		"time": schema.SingleNestedAttribute{
			Attributes: FirewallNameRuleTime{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Time to match rule

`,
		},

		"icmp": schema.SingleNestedAttribute{
			Attributes: FirewallNameRuleIcmp{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `ICMP type and code information

`,
		},

		"ttl": schema.SingleNestedAttribute{
			Attributes: FirewallNameRuleTTL{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Time to live limit

`,
		},
	}
}
