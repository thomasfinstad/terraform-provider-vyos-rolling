// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/numberplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &FirewallIPvfourNameRule{}

// FirewallIPvfourNameRule describes the resource data model.
type FirewallIPvfourNameRule struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Number `tfsdk:"rule_id" vyos:"-,self-id"`

	ParentIDFirewallIPvfourName types.String `tfsdk:"name_id" vyos:"name,parent-id"`

	// LeafNodes
	LeafFirewallIPvfourNameRuleAction              types.String `tfsdk:"action" vyos:"action,omitempty"`
	LeafFirewallIPvfourNameRuleDescrIPtion         types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafFirewallIPvfourNameRuleDscp                types.List   `tfsdk:"dscp" vyos:"dscp,omitempty"`
	LeafFirewallIPvfourNameRuleDscpExclude         types.List   `tfsdk:"dscp_exclude" vyos:"dscp-exclude,omitempty"`
	LeafFirewallIPvfourNameRulePacketLength        types.List   `tfsdk:"packet_length" vyos:"packet-length,omitempty"`
	LeafFirewallIPvfourNameRulePacketLengthExclude types.List   `tfsdk:"packet_length_exclude" vyos:"packet-length-exclude,omitempty"`
	LeafFirewallIPvfourNameRulePacketType          types.String `tfsdk:"packet_type" vyos:"packet-type,omitempty"`
	LeafFirewallIPvfourNameRuleMark                types.String `tfsdk:"mark" vyos:"mark,omitempty"`
	LeafFirewallIPvfourNameRuleConnectionMark      types.List   `tfsdk:"connection_mark" vyos:"connection-mark,omitempty"`
	LeafFirewallIPvfourNameRuleConntrackHelper     types.List   `tfsdk:"conntrack_helper" vyos:"conntrack-helper,omitempty"`
	LeafFirewallIPvfourNameRuleQueue               types.Number `tfsdk:"queue" vyos:"queue,omitempty"`
	LeafFirewallIPvfourNameRuleQueueOptions        types.List   `tfsdk:"queue_options" vyos:"queue-options,omitempty"`
	LeafFirewallIPvfourNameRuleDisable             types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafFirewallIPvfourNameRuleLog                 types.Bool   `tfsdk:"log" vyos:"log,omitempty"`
	LeafFirewallIPvfourNameRuleProtocol            types.String `tfsdk:"protocol" vyos:"protocol,omitempty"`
	LeafFirewallIPvfourNameRuleState               types.List   `tfsdk:"state" vyos:"state,omitempty"`
	LeafFirewallIPvfourNameRuleJumpTarget          types.String `tfsdk:"jump_target" vyos:"jump-target,omitempty"`
	LeafFirewallIPvfourNameRuleOffloadTarget       types.String `tfsdk:"offload_target" vyos:"offload-target,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeFirewallIPvfourNameRuleFragment          *FirewallIPvfourNameRuleFragment          `tfsdk:"fragment" vyos:"fragment,omitempty"`
	NodeFirewallIPvfourNameRuleLimit             *FirewallIPvfourNameRuleLimit             `tfsdk:"limit" vyos:"limit,omitempty"`
	NodeFirewallIPvfourNameRuleLogOptions        *FirewallIPvfourNameRuleLogOptions        `tfsdk:"log_options" vyos:"log-options,omitempty"`
	NodeFirewallIPvfourNameRuleConnectionStatus  *FirewallIPvfourNameRuleConnectionStatus  `tfsdk:"connection_status" vyos:"connection-status,omitempty"`
	NodeFirewallIPvfourNameRuleRecent            *FirewallIPvfourNameRuleRecent            `tfsdk:"recent" vyos:"recent,omitempty"`
	NodeFirewallIPvfourNameRuleSynproxy          *FirewallIPvfourNameRuleSynproxy          `tfsdk:"synproxy" vyos:"synproxy,omitempty"`
	NodeFirewallIPvfourNameRuleTCP               *FirewallIPvfourNameRuleTCP               `tfsdk:"tcp" vyos:"tcp,omitempty"`
	NodeFirewallIPvfourNameRuleTime              *FirewallIPvfourNameRuleTime              `tfsdk:"time" vyos:"time,omitempty"`
	NodeFirewallIPvfourNameRuleTTL               *FirewallIPvfourNameRuleTTL               `tfsdk:"ttl" vyos:"ttl,omitempty"`
	NodeFirewallIPvfourNameRuleAddAddressToGroup *FirewallIPvfourNameRuleAddAddressToGroup `tfsdk:"add_address_to_group" vyos:"add-address-to-group,omitempty"`
	NodeFirewallIPvfourNameRuleDestination       *FirewallIPvfourNameRuleDestination       `tfsdk:"destination" vyos:"destination,omitempty"`
	NodeFirewallIPvfourNameRuleIcmp              *FirewallIPvfourNameRuleIcmp              `tfsdk:"icmp" vyos:"icmp,omitempty"`
	NodeFirewallIPvfourNameRuleSource            *FirewallIPvfourNameRuleSource            `tfsdk:"source" vyos:"source,omitempty"`
	NodeFirewallIPvfourNameRuleInboundInterface  *FirewallIPvfourNameRuleInboundInterface  `tfsdk:"inbound_interface" vyos:"inbound-interface,omitempty"`
	NodeFirewallIPvfourNameRuleIPsec             *FirewallIPvfourNameRuleIPsec             `tfsdk:"ipsec" vyos:"ipsec,omitempty"`
	NodeFirewallIPvfourNameRuleOutboundInterface *FirewallIPvfourNameRuleOutboundInterface `tfsdk:"outbound_interface" vyos:"outbound-interface,omitempty"`
}

// SetID configures the resource ID
func (o *FirewallIPvfourNameRule) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *FirewallIPvfourNameRule) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"firewall",

		"ipv4",

		"name",
		o.ParentIDFirewallIPvfourName.ValueString(),

		"rule",
		o.SelfIdentifier.ValueBigFloat().String(),
	}
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *FirewallIPvfourNameRule) GetVyosParentPath() []string {
	return []string{
		"firewall",

		"ipv4",

		"name",
		o.ParentIDFirewallIPvfourName.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvfourNameRule) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"rule_id": schema.NumberAttribute{
			Required: true,
			MarkdownDescription: `IPv4 Firewall custom rule number

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-999999  &emsp; |  Number for this firewall rule  |

`,
			PlanModifiers: []planmodifier.Number{
				numberplanmodifier.RequiresReplace(),
			},
		},

		"name_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `IPv4 custom firewall

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			}, Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in name_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illigal character in  name_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
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

		"offload_target": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set flowtable offload target. Action offload must be defined to use this setting

`,
		},

		// Nodes

		"fragment": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourNameRuleFragment{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IP fragment match

`,
		},

		"limit": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourNameRuleLimit{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Rate limit using a token bucket filter

`,
		},

		"log_options": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourNameRuleLogOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Log options

`,
		},

		"connection_status": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourNameRuleConnectionStatus{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Connection status

`,
		},

		"recent": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourNameRuleRecent{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Parameters for matching recently seen sources

`,
		},

		"synproxy": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourNameRuleSynproxy{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Synproxy options

`,
		},

		"tcp": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourNameRuleTCP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `TCP options to match

`,
		},

		"time": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourNameRuleTime{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Time to match rule

`,
		},

		"ttl": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourNameRuleTTL{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Time to live limit

`,
		},

		"add_address_to_group": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourNameRuleAddAddressToGroup{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Add ip address to dynamic address-group

`,
		},

		"destination": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourNameRuleDestination{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Destination parameters

`,
		},

		"icmp": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourNameRuleIcmp{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `ICMP type and code information

`,
		},

		"source": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourNameRuleSource{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Source parameters

`,
		},

		"inbound_interface": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourNameRuleInboundInterface{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match inbound-interface

`,
		},

		"ipsec": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourNameRuleIPsec{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Inbound IPsec packets

`,
		},

		"outbound_interface": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourNameRuleOutboundInterface{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match outbound-interface

`,
		},
	}
}
