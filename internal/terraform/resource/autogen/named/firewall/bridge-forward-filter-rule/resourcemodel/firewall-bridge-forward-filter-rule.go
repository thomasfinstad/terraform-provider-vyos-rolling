// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/numberplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &FirewallBrIDgeForwardFilterRule{}

// FirewallBrIDgeForwardFilterRule describes the resource data model.
type FirewallBrIDgeForwardFilterRule struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafFirewallBrIDgeForwardFilterRuleDescrIPtion         types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafFirewallBrIDgeForwardFilterRuleDisable             types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafFirewallBrIDgeForwardFilterRuleDscp                types.List   `tfsdk:"dscp" vyos:"dscp,omitempty"`
	LeafFirewallBrIDgeForwardFilterRuleDscpExclude         types.List   `tfsdk:"dscp_exclude" vyos:"dscp-exclude,omitempty"`
	LeafFirewallBrIDgeForwardFilterRuleMark                types.String `tfsdk:"mark" vyos:"mark,omitempty"`
	LeafFirewallBrIDgeForwardFilterRuleLog                 types.Bool   `tfsdk:"log" vyos:"log,omitempty"`
	LeafFirewallBrIDgeForwardFilterRuleEthernetType        types.String `tfsdk:"ethernet_type" vyos:"ethernet-type,omitempty"`
	LeafFirewallBrIDgeForwardFilterRuleQueue               types.Number `tfsdk:"queue" vyos:"queue,omitempty"`
	LeafFirewallBrIDgeForwardFilterRuleQueueOptions        types.List   `tfsdk:"queue_options" vyos:"queue-options,omitempty"`
	LeafFirewallBrIDgeForwardFilterRulePacketLength        types.List   `tfsdk:"packet_length" vyos:"packet-length,omitempty"`
	LeafFirewallBrIDgeForwardFilterRulePacketLengthExclude types.List   `tfsdk:"packet_length_exclude" vyos:"packet-length-exclude,omitempty"`
	LeafFirewallBrIDgeForwardFilterRulePacketType          types.String `tfsdk:"packet_type" vyos:"packet-type,omitempty"`
	LeafFirewallBrIDgeForwardFilterRuleProtocol            types.String `tfsdk:"protocol" vyos:"protocol,omitempty"`
	LeafFirewallBrIDgeForwardFilterRuleJumpTarget          types.String `tfsdk:"jump_target" vyos:"jump-target,omitempty"`
	LeafFirewallBrIDgeForwardFilterRuleAction              types.String `tfsdk:"action" vyos:"action,omitempty"`
	LeafFirewallBrIDgeForwardFilterRuleConnectionMark      types.List   `tfsdk:"connection_mark" vyos:"connection-mark,omitempty"`
	LeafFirewallBrIDgeForwardFilterRuleState               types.List   `tfsdk:"state" vyos:"state,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeFirewallBrIDgeForwardFilterRuleFragment          *FirewallBrIDgeForwardFilterRuleFragment          `tfsdk:"fragment" vyos:"fragment,omitempty"`
	NodeFirewallBrIDgeForwardFilterRuleHopLimit          *FirewallBrIDgeForwardFilterRuleHopLimit          `tfsdk:"hop_limit" vyos:"hop-limit,omitempty"`
	NodeFirewallBrIDgeForwardFilterRuleIcmp              *FirewallBrIDgeForwardFilterRuleIcmp              `tfsdk:"icmp" vyos:"icmp,omitempty"`
	NodeFirewallBrIDgeForwardFilterRuleIcmpvsix          *FirewallBrIDgeForwardFilterRuleIcmpvsix          `tfsdk:"icmpv6" vyos:"icmpv6,omitempty"`
	NodeFirewallBrIDgeForwardFilterRuleLimit             *FirewallBrIDgeForwardFilterRuleLimit             `tfsdk:"limit" vyos:"limit,omitempty"`
	NodeFirewallBrIDgeForwardFilterRuleLogOptions        *FirewallBrIDgeForwardFilterRuleLogOptions        `tfsdk:"log_options" vyos:"log-options,omitempty"`
	NodeFirewallBrIDgeForwardFilterRuleIPsec             *FirewallBrIDgeForwardFilterRuleIPsec             `tfsdk:"ipsec" vyos:"ipsec,omitempty"`
	NodeFirewallBrIDgeForwardFilterRuleVlan              *FirewallBrIDgeForwardFilterRuleVlan              `tfsdk:"vlan" vyos:"vlan,omitempty"`
	NodeFirewallBrIDgeForwardFilterRuleTCP               *FirewallBrIDgeForwardFilterRuleTCP               `tfsdk:"tcp" vyos:"tcp,omitempty"`
	NodeFirewallBrIDgeForwardFilterRuleTime              *FirewallBrIDgeForwardFilterRuleTime              `tfsdk:"time" vyos:"time,omitempty"`
	NodeFirewallBrIDgeForwardFilterRuleTTL               *FirewallBrIDgeForwardFilterRuleTTL               `tfsdk:"ttl" vyos:"ttl,omitempty"`
	NodeFirewallBrIDgeForwardFilterRuleDestination       *FirewallBrIDgeForwardFilterRuleDestination       `tfsdk:"destination" vyos:"destination,omitempty"`
	NodeFirewallBrIDgeForwardFilterRuleSource            *FirewallBrIDgeForwardFilterRuleSource            `tfsdk:"source" vyos:"source,omitempty"`
	NodeFirewallBrIDgeForwardFilterRuleConnectionStatus  *FirewallBrIDgeForwardFilterRuleConnectionStatus  `tfsdk:"connection_status" vyos:"connection-status,omitempty"`
	NodeFirewallBrIDgeForwardFilterRuleInboundInterface  *FirewallBrIDgeForwardFilterRuleInboundInterface  `tfsdk:"inbound_interface" vyos:"inbound-interface,omitempty"`
	NodeFirewallBrIDgeForwardFilterRuleOutboundInterface *FirewallBrIDgeForwardFilterRuleOutboundInterface `tfsdk:"outbound_interface" vyos:"outbound-interface,omitempty"`
	NodeFirewallBrIDgeForwardFilterRuleSet               *FirewallBrIDgeForwardFilterRuleSet               `tfsdk:"set" vyos:"set,omitempty"`
}

// SetID configures the resource ID
func (o *FirewallBrIDgeForwardFilterRule) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *FirewallBrIDgeForwardFilterRule) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *FirewallBrIDgeForwardFilterRule) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *FirewallBrIDgeForwardFilterRule) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"rule",
		o.SelfIdentifier.Attributes()["rule"].(types.Number).ValueBigFloat().String(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *FirewallBrIDgeForwardFilterRule) GetVyosParentPath() []string {
	return []string{
		"firewall",

		"bridge",

		"forward",

		"filter",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *FirewallBrIDgeForwardFilterRule) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallBrIDgeForwardFilterRule) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"rule": schema.NumberAttribute{
					Required: true,
					MarkdownDescription: `Bridge Firewall forward filter rule number

    |  Format    |  Description                    |
    |------------|---------------------------------|
    |  1-999999  |  Number for this firewall rule  |
`,
					Description: `Bridge Firewall forward filter rule number

    |  Format    |  Description                    |
    |------------|---------------------------------|
    |  1-999999  |  Number for this firewall rule  |
`,
					PlanModifiers: []planmodifier.Number{
						numberplanmodifier.RequiresReplace(),
					},
				},
			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Description  |
`,
			Description: `Description

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Description  |
`,
		},

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
			Description: `Disable instance

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"dscp": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `DSCP value

    |  Format       |  Description          |
    |---------------|-----------------------|
    |  0-63         |  DSCP value to match  |
    |  <start-end>  |  DSCP range to match  |
`,
			Description: `DSCP value

    |  Format       |  Description          |
    |---------------|-----------------------|
    |  0-63         |  DSCP value to match  |
    |  <start-end>  |  DSCP range to match  |
`,
		},

		"dscp_exclude": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `DSCP value not to match

    |  Format       |  Description              |
    |---------------|---------------------------|
    |  0-63         |  DSCP value not to match  |
    |  <start-end>  |  DSCP range not to match  |
`,
			Description: `DSCP value not to match

    |  Format       |  Description              |
    |---------------|---------------------------|
    |  0-63         |  DSCP value not to match  |
    |  <start-end>  |  DSCP range not to match  |
`,
		},

		"mark": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Firewall mark

    |  Format         |  Description                            |
    |-----------------|-----------------------------------------|
    |  0-2147483647   |  Firewall mark to match                 |
    |  !0-2147483647  |  Inverted Firewall mark to match        |
    |  <start-end>    |  Firewall mark range to match           |
    |  !<start-end>   |  Firewall mark inverted range to match  |
`,
			Description: `Firewall mark

    |  Format         |  Description                            |
    |-----------------|-----------------------------------------|
    |  0-2147483647   |  Firewall mark to match                 |
    |  !0-2147483647  |  Inverted Firewall mark to match        |
    |  <start-end>    |  Firewall mark range to match           |
    |  !<start-end>   |  Firewall mark inverted range to match  |
`,
		},

		"log": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Log packets hitting this rule

`,
			Description: `Log packets hitting this rule

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"ethernet_type": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Ethernet type

    |  Format   |  Description                  |
    |-----------|-------------------------------|
    |  802.1q   |  Customer VLAN tag type       |
    |  802.1ad  |  Service VLAN tag type        |
    |  arp      |  Adress Resolution Protocol   |
    |  _ipv4    |  Internet Protocol version 4  |
    |  _ipv6    |  Internet Protocol version 6  |
`,
			Description: `Ethernet type

    |  Format   |  Description                  |
    |-----------|-------------------------------|
    |  802.1q   |  Customer VLAN tag type       |
    |  802.1ad  |  Service VLAN tag type        |
    |  arp      |  Adress Resolution Protocol   |
    |  _ipv4    |  Internet Protocol version 4  |
    |  _ipv6    |  Internet Protocol version 6  |
`,
		},

		"queue": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Queue target to use. Action queue must be defined to use this setting

    |  Format   |  Description   |
    |-----------|----------------|
    |  0-65535  |  Queue target  |
`,
			Description: `Queue target to use. Action queue must be defined to use this setting

    |  Format   |  Description   |
    |-----------|----------------|
    |  0-65535  |  Queue target  |
`,
		},

		"queue_options": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Options used for queue target. Action queue must be defined to use this setting

    |  Format  |  Description                                                      |
    |----------|-------------------------------------------------------------------|
    |  bypass  |  Let packets go through if userspace application cannot back off  |
    |  fanout  |  Distribute packets between several queues                        |
`,
			Description: `Options used for queue target. Action queue must be defined to use this setting

    |  Format  |  Description                                                      |
    |----------|-------------------------------------------------------------------|
    |  bypass  |  Let packets go through if userspace application cannot back off  |
    |  fanout  |  Distribute packets between several queues                        |
`,
		},

		"packet_length": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Payload size in bytes, including header and data to match

    |  Format       |  Description                   |
    |---------------|--------------------------------|
    |  1-65535      |  Packet length to match        |
    |  <start-end>  |  Packet length range to match  |
`,
			Description: `Payload size in bytes, including header and data to match

    |  Format       |  Description                   |
    |---------------|--------------------------------|
    |  1-65535      |  Packet length to match        |
    |  <start-end>  |  Packet length range to match  |
`,
		},

		"packet_length_exclude": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Payload size in bytes, including header and data not to match

    |  Format       |  Description                       |
    |---------------|------------------------------------|
    |  1-65535      |  Packet length not to match        |
    |  <start-end>  |  Packet length range not to match  |
`,
			Description: `Payload size in bytes, including header and data not to match

    |  Format       |  Description                       |
    |---------------|------------------------------------|
    |  1-65535      |  Packet length not to match        |
    |  <start-end>  |  Packet length range not to match  |
`,
		},

		"packet_type": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Packet type

    |  Format     |  Description                                      |
    |-------------|---------------------------------------------------|
    |  broadcast  |  Match broadcast packet type                      |
    |  host       |  Match host packet type, addressed to local host  |
    |  multicast  |  Match multicast packet type                      |
    |  other      |  Match packet addressed to another host           |
`,
			Description: `Packet type

    |  Format     |  Description                                      |
    |-------------|---------------------------------------------------|
    |  broadcast  |  Match broadcast packet type                      |
    |  host       |  Match host packet type, addressed to local host  |
    |  multicast  |  Match multicast packet type                      |
    |  other      |  Match packet addressed to another host           |
`,
		},

		"protocol": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Protocol to match (protocol name, number, or "all")

    |  Format       |  Description         |
    |---------------|----------------------|
    |  all          |  All IP protocols    |
    |  tcp_udp      |  Both TCP and UDP    |
    |  0-255        |  IP protocol number  |
    |  <protocol>   |  IP protocol name    |
    |  !<protocol>  |  IP protocol name    |
`,
			Description: `Protocol to match (protocol name, number, or "all")

    |  Format       |  Description         |
    |---------------|----------------------|
    |  all          |  All IP protocols    |
    |  tcp_udp      |  Both TCP and UDP    |
    |  0-255        |  IP protocol number  |
    |  <protocol>   |  IP protocol name    |
    |  !<protocol>  |  IP protocol name    |
`,
		},

		"jump_target": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set jump target. Action jump must be defined to use this setting

`,
			Description: `Set jump target. Action jump must be defined to use this setting

`,
		},

		"action": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Rule action

    |  Format    |  Description                                                                    |
    |------------|---------------------------------------------------------------------------------|
    |  accept    |  Accept matching entries                                                        |
    |  continue  |  Continue parsing next rule                                                     |
    |  jump      |  Jump to another chain                                                          |
    |  return    |  Return from the current chain and continue at the next rule of the last chain  |
    |  drop      |  Drop matching entries                                                          |
    |  queue     |  Enqueue packet to userspace                                                    |
`,
			Description: `Rule action

    |  Format    |  Description                                                                    |
    |------------|---------------------------------------------------------------------------------|
    |  accept    |  Accept matching entries                                                        |
    |  continue  |  Continue parsing next rule                                                     |
    |  jump      |  Jump to another chain                                                          |
    |  return    |  Return from the current chain and continue at the next rule of the last chain  |
    |  drop      |  Drop matching entries                                                          |
    |  queue     |  Enqueue packet to userspace                                                    |
`,
		},

		"connection_mark": schema.ListAttribute{
			ElementType: types.NumberType,
			Optional:    true,
			MarkdownDescription: `Connection mark

    |  Format        |  Description               |
    |----------------|----------------------------|
    |  0-2147483647  |  Connection-mark to match  |
`,
			Description: `Connection mark

    |  Format        |  Description               |
    |----------------|----------------------------|
    |  0-2147483647  |  Connection-mark to match  |
`,
		},

		"state": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Session state

    |  Format       |  Description        |
    |---------------|---------------------|
    |  established  |  Established state  |
    |  invalid      |  Invalid state      |
    |  new          |  New state          |
    |  related      |  Related state      |
`,
			Description: `Session state

    |  Format       |  Description        |
    |---------------|---------------------|
    |  established  |  Established state  |
    |  invalid      |  Invalid state      |
    |  new          |  New state          |
    |  related      |  Related state      |
`,
		},

		// Nodes

		"fragment": schema.SingleNestedAttribute{
			Attributes: FirewallBrIDgeForwardFilterRuleFragment{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IP fragment match

`,
			Description: `IP fragment match

`,
		},

		"hop_limit": schema.SingleNestedAttribute{
			Attributes: FirewallBrIDgeForwardFilterRuleHopLimit{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Hop limit

`,
			Description: `Hop limit

`,
		},

		"icmp": schema.SingleNestedAttribute{
			Attributes: FirewallBrIDgeForwardFilterRuleIcmp{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `ICMP type and code information

`,
			Description: `ICMP type and code information

`,
		},

		"icmpv6": schema.SingleNestedAttribute{
			Attributes: FirewallBrIDgeForwardFilterRuleIcmpvsix{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `ICMPv6 type and code information

`,
			Description: `ICMPv6 type and code information

`,
		},

		"limit": schema.SingleNestedAttribute{
			Attributes: FirewallBrIDgeForwardFilterRuleLimit{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Rate limit using a token bucket filter

`,
			Description: `Rate limit using a token bucket filter

`,
		},

		"log_options": schema.SingleNestedAttribute{
			Attributes: FirewallBrIDgeForwardFilterRuleLogOptions{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Log options

`,
			Description: `Log options

`,
		},

		"ipsec": schema.SingleNestedAttribute{
			Attributes: FirewallBrIDgeForwardFilterRuleIPsec{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPsec encapsulated packets

`,
			Description: `IPsec encapsulated packets

`,
		},

		"vlan": schema.SingleNestedAttribute{
			Attributes: FirewallBrIDgeForwardFilterRuleVlan{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `VLAN parameters

`,
			Description: `VLAN parameters

`,
		},

		"tcp": schema.SingleNestedAttribute{
			Attributes: FirewallBrIDgeForwardFilterRuleTCP{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `TCP options to match

`,
			Description: `TCP options to match

`,
		},

		"time": schema.SingleNestedAttribute{
			Attributes: FirewallBrIDgeForwardFilterRuleTime{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Time to match rule

`,
			Description: `Time to match rule

`,
		},

		"ttl": schema.SingleNestedAttribute{
			Attributes: FirewallBrIDgeForwardFilterRuleTTL{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Time to live limit

`,
			Description: `Time to live limit

`,
		},

		"destination": schema.SingleNestedAttribute{
			Attributes: FirewallBrIDgeForwardFilterRuleDestination{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Destination parameters

`,
			Description: `Destination parameters

`,
		},

		"source": schema.SingleNestedAttribute{
			Attributes: FirewallBrIDgeForwardFilterRuleSource{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Source parameters

`,
			Description: `Source parameters

`,
		},

		"connection_status": schema.SingleNestedAttribute{
			Attributes: FirewallBrIDgeForwardFilterRuleConnectionStatus{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Connection status

`,
			Description: `Connection status

`,
		},

		"inbound_interface": schema.SingleNestedAttribute{
			Attributes: FirewallBrIDgeForwardFilterRuleInboundInterface{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Match inbound-interface

`,
			Description: `Match inbound-interface

`,
		},

		"outbound_interface": schema.SingleNestedAttribute{
			Attributes: FirewallBrIDgeForwardFilterRuleOutboundInterface{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Match outbound-interface

`,
			Description: `Match outbound-interface

`,
		},

		"set": schema.SingleNestedAttribute{
			Attributes: FirewallBrIDgeForwardFilterRuleSet{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Packet modifications

`,
			Description: `Packet modifications

`,
		},
	}
}
