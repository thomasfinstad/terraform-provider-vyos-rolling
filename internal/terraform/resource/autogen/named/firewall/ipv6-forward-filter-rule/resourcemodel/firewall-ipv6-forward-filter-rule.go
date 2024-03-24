// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
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

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &FirewallIPvsixForwardFilterRule{}

// FirewallIPvsixForwardFilterRule describes the resource data model.
type FirewallIPvsixForwardFilterRule struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Number `tfsdk:"rule_id" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafFirewallIPvsixForwardFilterRuleAction              types.String `tfsdk:"action" vyos:"action,omitempty"`
	LeafFirewallIPvsixForwardFilterRuleDescrIPtion         types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafFirewallIPvsixForwardFilterRuleDscp                types.List   `tfsdk:"dscp" vyos:"dscp,omitempty"`
	LeafFirewallIPvsixForwardFilterRuleDscpExclude         types.List   `tfsdk:"dscp_exclude" vyos:"dscp-exclude,omitempty"`
	LeafFirewallIPvsixForwardFilterRulePacketLength        types.List   `tfsdk:"packet_length" vyos:"packet-length,omitempty"`
	LeafFirewallIPvsixForwardFilterRulePacketLengthExclude types.List   `tfsdk:"packet_length_exclude" vyos:"packet-length-exclude,omitempty"`
	LeafFirewallIPvsixForwardFilterRulePacketType          types.String `tfsdk:"packet_type" vyos:"packet-type,omitempty"`
	LeafFirewallIPvsixForwardFilterRuleMark                types.String `tfsdk:"mark" vyos:"mark,omitempty"`
	LeafFirewallIPvsixForwardFilterRuleConnectionMark      types.List   `tfsdk:"connection_mark" vyos:"connection-mark,omitempty"`
	LeafFirewallIPvsixForwardFilterRuleConntrackHelper     types.List   `tfsdk:"conntrack_helper" vyos:"conntrack-helper,omitempty"`
	LeafFirewallIPvsixForwardFilterRuleQueue               types.Number `tfsdk:"queue" vyos:"queue,omitempty"`
	LeafFirewallIPvsixForwardFilterRuleQueueOptions        types.List   `tfsdk:"queue_options" vyos:"queue-options,omitempty"`
	LeafFirewallIPvsixForwardFilterRuleDisable             types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafFirewallIPvsixForwardFilterRuleLog                 types.Bool   `tfsdk:"log" vyos:"log,omitempty"`
	LeafFirewallIPvsixForwardFilterRuleProtocol            types.String `tfsdk:"protocol" vyos:"protocol,omitempty"`
	LeafFirewallIPvsixForwardFilterRuleState               types.List   `tfsdk:"state" vyos:"state,omitempty"`
	LeafFirewallIPvsixForwardFilterRuleJumpTarget          types.String `tfsdk:"jump_target" vyos:"jump-target,omitempty"`
	LeafFirewallIPvsixForwardFilterRuleOffloadTarget       types.String `tfsdk:"offload_target" vyos:"offload-target,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeFirewallIPvsixForwardFilterRuleFragment          *FirewallIPvsixForwardFilterRuleFragment          `tfsdk:"fragment" vyos:"fragment,omitempty"`
	NodeFirewallIPvsixForwardFilterRuleLimit             *FirewallIPvsixForwardFilterRuleLimit             `tfsdk:"limit" vyos:"limit,omitempty"`
	NodeFirewallIPvsixForwardFilterRuleLogOptions        *FirewallIPvsixForwardFilterRuleLogOptions        `tfsdk:"log_options" vyos:"log-options,omitempty"`
	NodeFirewallIPvsixForwardFilterRuleConnectionStatus  *FirewallIPvsixForwardFilterRuleConnectionStatus  `tfsdk:"connection_status" vyos:"connection-status,omitempty"`
	NodeFirewallIPvsixForwardFilterRuleRecent            *FirewallIPvsixForwardFilterRuleRecent            `tfsdk:"recent" vyos:"recent,omitempty"`
	NodeFirewallIPvsixForwardFilterRuleSynproxy          *FirewallIPvsixForwardFilterRuleSynproxy          `tfsdk:"synproxy" vyos:"synproxy,omitempty"`
	NodeFirewallIPvsixForwardFilterRuleTCP               *FirewallIPvsixForwardFilterRuleTCP               `tfsdk:"tcp" vyos:"tcp,omitempty"`
	NodeFirewallIPvsixForwardFilterRuleTime              *FirewallIPvsixForwardFilterRuleTime              `tfsdk:"time" vyos:"time,omitempty"`
	NodeFirewallIPvsixForwardFilterRuleHopLimit          *FirewallIPvsixForwardFilterRuleHopLimit          `tfsdk:"hop_limit" vyos:"hop-limit,omitempty"`
	NodeFirewallIPvsixForwardFilterRuleAddAddressToGroup *FirewallIPvsixForwardFilterRuleAddAddressToGroup `tfsdk:"add_address_to_group" vyos:"add-address-to-group,omitempty"`
	NodeFirewallIPvsixForwardFilterRuleDestination       *FirewallIPvsixForwardFilterRuleDestination       `tfsdk:"destination" vyos:"destination,omitempty"`
	NodeFirewallIPvsixForwardFilterRuleIcmpvsix          *FirewallIPvsixForwardFilterRuleIcmpvsix          `tfsdk:"icmpv6" vyos:"icmpv6,omitempty"`
	NodeFirewallIPvsixForwardFilterRuleSource            *FirewallIPvsixForwardFilterRuleSource            `tfsdk:"source" vyos:"source,omitempty"`
	NodeFirewallIPvsixForwardFilterRuleInboundInterface  *FirewallIPvsixForwardFilterRuleInboundInterface  `tfsdk:"inbound_interface" vyos:"inbound-interface,omitempty"`
	NodeFirewallIPvsixForwardFilterRuleIPsec             *FirewallIPvsixForwardFilterRuleIPsec             `tfsdk:"ipsec" vyos:"ipsec,omitempty"`
	NodeFirewallIPvsixForwardFilterRuleOutboundInterface *FirewallIPvsixForwardFilterRuleOutboundInterface `tfsdk:"outbound_interface" vyos:"outbound-interface,omitempty"`
}

// SetID configures the resource ID
func (o *FirewallIPvsixForwardFilterRule) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *FirewallIPvsixForwardFilterRule) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *FirewallIPvsixForwardFilterRule) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *FirewallIPvsixForwardFilterRule) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"rule",
		o.SelfIdentifier.ValueBigFloat().String(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *FirewallIPvsixForwardFilterRule) GetVyosParentPath() []string {
	return []string{
		"firewall",

		"ipv6",

		"forward",

		"filter",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *FirewallIPvsixForwardFilterRule) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixForwardFilterRule) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"rule_id": schema.NumberAttribute{
			Required: true,
			MarkdownDescription: `IPv6 Firewall forward filter rule number

    |  Format    |  Description                    |
    |------------|---------------------------------|
    |  1-999999  |  Number for this firewall rule  |
`,
			Description: `IPv6 Firewall forward filter rule number

    |  Format    |  Description                    |
    |------------|---------------------------------|
    |  1-999999  |  Number for this firewall rule  |
`,
			PlanModifiers: []planmodifier.Number{
				numberplanmodifier.RequiresReplace(),
			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"action": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Rule action

    |  Format    |  Description                                                                    |
    |------------|---------------------------------------------------------------------------------|
    |  accept    |  Accept matching entries                                                        |
    |  continue  |  Continue parsing next rule                                                     |
    |  jump      |  Jump to another chain                                                          |
    |  reject    |  Reject matching entries                                                        |
    |  return    |  Return from the current chain and continue at the next rule of the last chain  |
    |  drop      |  Drop matching entries                                                          |
    |  queue     |  Enqueue packet to userspace                                                    |
    |  offload   |  Offload packet via flowtable                                                   |
    |  synproxy  |  Synproxy connections                                                           |
`,
			Description: `Rule action

    |  Format    |  Description                                                                    |
    |------------|---------------------------------------------------------------------------------|
    |  accept    |  Accept matching entries                                                        |
    |  continue  |  Continue parsing next rule                                                     |
    |  jump      |  Jump to another chain                                                          |
    |  reject    |  Reject matching entries                                                        |
    |  return    |  Return from the current chain and continue at the next rule of the last chain  |
    |  drop      |  Drop matching entries                                                          |
    |  queue     |  Enqueue packet to userspace                                                    |
    |  offload   |  Offload packet via flowtable                                                   |
    |  synproxy  |  Synproxy connections                                                           |
`,
		},

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

		"conntrack_helper": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Match related traffic from conntrack helpers

    |  Format  |  Description                         |
    |----------|--------------------------------------|
    |  ftp     |  Related traffic from FTP helper     |
    |  h323    |  Related traffic from H.323 helper   |
    |  pptp    |  Related traffic from PPTP helper    |
    |  nfs     |  Related traffic from NFS helper     |
    |  rtsp    |  Related traffic from RTSP helper    |
    |  sip     |  Related traffic from SIP helper     |
    |  tftp    |  Related traffic from TFTP helper    |
    |  sqlnet  |  Related traffic from SQLNet helper  |
`,
			Description: `Match related traffic from conntrack helpers

    |  Format  |  Description                         |
    |----------|--------------------------------------|
    |  ftp     |  Related traffic from FTP helper     |
    |  h323    |  Related traffic from H.323 helper   |
    |  pptp    |  Related traffic from PPTP helper    |
    |  nfs     |  Related traffic from NFS helper     |
    |  rtsp    |  Related traffic from RTSP helper    |
    |  sip     |  Related traffic from SIP helper     |
    |  tftp    |  Related traffic from TFTP helper    |
    |  sqlnet  |  Related traffic from SQLNet helper  |
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

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
			Description: `Disable instance

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
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

		"jump_target": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set jump target. Action jump must be defined to use this setting

`,
			Description: `Set jump target. Action jump must be defined to use this setting

`,
		},

		"offload_target": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set flowtable offload target. Action offload must be defined to use this setting

`,
			Description: `Set flowtable offload target. Action offload must be defined to use this setting

`,
		},

		// Nodes

		"fragment": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixForwardFilterRuleFragment{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IP fragment match

`,
			Description: `IP fragment match

`,
		},

		"limit": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixForwardFilterRuleLimit{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Rate limit using a token bucket filter

`,
			Description: `Rate limit using a token bucket filter

`,
		},

		"log_options": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixForwardFilterRuleLogOptions{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Log options

`,
			Description: `Log options

`,
		},

		"connection_status": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixForwardFilterRuleConnectionStatus{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Connection status

`,
			Description: `Connection status

`,
		},

		"recent": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixForwardFilterRuleRecent{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Parameters for matching recently seen sources

`,
			Description: `Parameters for matching recently seen sources

`,
		},

		"synproxy": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixForwardFilterRuleSynproxy{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Synproxy options

`,
			Description: `Synproxy options

`,
		},

		"tcp": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixForwardFilterRuleTCP{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `TCP options to match

`,
			Description: `TCP options to match

`,
		},

		"time": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixForwardFilterRuleTime{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Time to match rule

`,
			Description: `Time to match rule

`,
		},

		"hop_limit": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixForwardFilterRuleHopLimit{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Hop limit

`,
			Description: `Hop limit

`,
		},

		"add_address_to_group": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixForwardFilterRuleAddAddressToGroup{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Add ipv6 address to dynamic ipv6-address-group

`,
			Description: `Add ipv6 address to dynamic ipv6-address-group

`,
		},

		"destination": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixForwardFilterRuleDestination{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Destination parameters

`,
			Description: `Destination parameters

`,
		},

		"icmpv6": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixForwardFilterRuleIcmpvsix{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `ICMPv6 type and code information

`,
			Description: `ICMPv6 type and code information

`,
		},

		"source": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixForwardFilterRuleSource{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Source parameters

`,
			Description: `Source parameters

`,
		},

		"inbound_interface": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixForwardFilterRuleInboundInterface{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Match inbound-interface

`,
			Description: `Match inbound-interface

`,
		},

		"ipsec": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixForwardFilterRuleIPsec{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Inbound IPsec packets

`,
			Description: `Inbound IPsec packets

`,
		},

		"outbound_interface": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixForwardFilterRuleOutboundInterface{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Match outbound-interface

`,
			Description: `Match outbound-interface

`,
		},
	}
}
