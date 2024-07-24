// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/numberplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &FirewallIPvfourNameRule{}

// FirewallIPvfourNameRule describes the resource data model.
type FirewallIPvfourNameRule struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Number `tfsdk:"rule_id" vyos:"-,self-id"`

	ParentIDFirewallIPvfourName types.String `tfsdk:"name_id" vyos:"name,parent-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafFirewallIPvfourNameRuleAction              types.String `tfsdk:"action" vyos:"action,omitempty"`
	LeafFirewallIPvfourNameRuleConntrackHelper     types.List   `tfsdk:"conntrack_helper" vyos:"conntrack-helper,omitempty"`
	LeafFirewallIPvfourNameRuleConnectionMark      types.List   `tfsdk:"connection_mark" vyos:"connection-mark,omitempty"`
	LeafFirewallIPvfourNameRuleDescrIPtion         types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafFirewallIPvfourNameRuleDisable             types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafFirewallIPvfourNameRuleDscp                types.List   `tfsdk:"dscp" vyos:"dscp,omitempty"`
	LeafFirewallIPvfourNameRuleDscpExclude         types.List   `tfsdk:"dscp_exclude" vyos:"dscp-exclude,omitempty"`
	LeafFirewallIPvfourNameRuleLog                 types.Bool   `tfsdk:"log" vyos:"log,omitempty"`
	LeafFirewallIPvfourNameRuleMark                types.String `tfsdk:"mark" vyos:"mark,omitempty"`
	LeafFirewallIPvfourNameRulePacketLength        types.List   `tfsdk:"packet_length" vyos:"packet-length,omitempty"`
	LeafFirewallIPvfourNameRulePacketLengthExclude types.List   `tfsdk:"packet_length_exclude" vyos:"packet-length-exclude,omitempty"`
	LeafFirewallIPvfourNameRulePacketType          types.String `tfsdk:"packet_type" vyos:"packet-type,omitempty"`
	LeafFirewallIPvfourNameRuleProtocol            types.String `tfsdk:"protocol" vyos:"protocol,omitempty"`
	LeafFirewallIPvfourNameRuleQueue               types.Number `tfsdk:"queue" vyos:"queue,omitempty"`
	LeafFirewallIPvfourNameRuleQueueOptions        types.List   `tfsdk:"queue_options" vyos:"queue-options,omitempty"`
	LeafFirewallIPvfourNameRuleState               types.List   `tfsdk:"state" vyos:"state,omitempty"`
	LeafFirewallIPvfourNameRuleJumpTarget          types.String `tfsdk:"jump_target" vyos:"jump-target,omitempty"`
	LeafFirewallIPvfourNameRuleOffloadTarget       types.String `tfsdk:"offload_target" vyos:"offload-target,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeFirewallIPvfourNameRuleAddAddressToGroup *FirewallIPvfourNameRuleAddAddressToGroup `tfsdk:"add_address_to_group" vyos:"add-address-to-group,omitempty"`
	NodeFirewallIPvfourNameRuleConnectionStatus  *FirewallIPvfourNameRuleConnectionStatus  `tfsdk:"connection_status" vyos:"connection-status,omitempty"`
	NodeFirewallIPvfourNameRuleFragment          *FirewallIPvfourNameRuleFragment          `tfsdk:"fragment" vyos:"fragment,omitempty"`
	NodeFirewallIPvfourNameRuleIPsec             *FirewallIPvfourNameRuleIPsec             `tfsdk:"ipsec" vyos:"ipsec,omitempty"`
	NodeFirewallIPvfourNameRuleLimit             *FirewallIPvfourNameRuleLimit             `tfsdk:"limit" vyos:"limit,omitempty"`
	NodeFirewallIPvfourNameRuleLogOptions        *FirewallIPvfourNameRuleLogOptions        `tfsdk:"log_options" vyos:"log-options,omitempty"`
	NodeFirewallIPvfourNameRuleRecent            *FirewallIPvfourNameRuleRecent            `tfsdk:"recent" vyos:"recent,omitempty"`
	NodeFirewallIPvfourNameRuleSynproxy          *FirewallIPvfourNameRuleSynproxy          `tfsdk:"synproxy" vyos:"synproxy,omitempty"`
	NodeFirewallIPvfourNameRuleTCP               *FirewallIPvfourNameRuleTCP               `tfsdk:"tcp" vyos:"tcp,omitempty"`
	NodeFirewallIPvfourNameRuleTime              *FirewallIPvfourNameRuleTime              `tfsdk:"time" vyos:"time,omitempty"`
	NodeFirewallIPvfourNameRuleIcmp              *FirewallIPvfourNameRuleIcmp              `tfsdk:"icmp" vyos:"icmp,omitempty"`
	NodeFirewallIPvfourNameRuleTTL               *FirewallIPvfourNameRuleTTL               `tfsdk:"ttl" vyos:"ttl,omitempty"`
	NodeFirewallIPvfourNameRuleDestination       *FirewallIPvfourNameRuleDestination       `tfsdk:"destination" vyos:"destination,omitempty"`
	NodeFirewallIPvfourNameRuleSource            *FirewallIPvfourNameRuleSource            `tfsdk:"source" vyos:"source,omitempty"`
	NodeFirewallIPvfourNameRuleInboundInterface  *FirewallIPvfourNameRuleInboundInterface  `tfsdk:"inbound_interface" vyos:"inbound-interface,omitempty"`
	NodeFirewallIPvfourNameRuleOutboundInterface *FirewallIPvfourNameRuleOutboundInterface `tfsdk:"outbound_interface" vyos:"outbound-interface,omitempty"`
}

// SetID configures the resource ID
func (o *FirewallIPvfourNameRule) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *FirewallIPvfourNameRule) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *FirewallIPvfourNameRule) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *FirewallIPvfourNameRule) GetVyosPath() []string {
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
func (o *FirewallIPvfourNameRule) GetVyosParentPath() []string {
	return []string{
		"firewall",

		"ipv4",

		"name",
		o.ParentIDFirewallIPvfourName.ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *FirewallIPvfourNameRule) GetVyosNamedParentPath() []string {
	return []string{
		"firewall",

		"ipv4",

		"name",
		o.ParentIDFirewallIPvfourName.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvfourNameRule) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"rule_id": schema.NumberAttribute{
			Required: true,
			MarkdownDescription: `IPv4 Firewall custom rule number

    |  Format    |  Description                    |
    |------------|---------------------------------|
    |  1-999999  |  Number for this firewall rule  |
`,
			Description: `IPv4 Firewall custom rule number

    |  Format    |  Description                    |
    |------------|---------------------------------|
    |  1-999999  |  Number for this firewall rule  |
`,
			PlanModifiers: []planmodifier.Number{
				numberplanmodifier.RequiresReplace(),
			},
		},

		"name_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `IPv4 custom firewall

`,
			Description: `IPv4 custom firewall

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
						"illegal character in  name_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
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

		"log": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Log packets hitting this rule

`,
			Description: `Log packets hitting this rule

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
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

		"add_address_to_group": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourNameRuleAddAddressToGroup{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Add ip address to dynamic address-group

`,
			Description: `Add ip address to dynamic address-group

`,
		},

		"connection_status": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourNameRuleConnectionStatus{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Connection status

`,
			Description: `Connection status

`,
		},

		"fragment": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourNameRuleFragment{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IP fragment match

`,
			Description: `IP fragment match

`,
		},

		"ipsec": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourNameRuleIPsec{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Inbound IPsec packets

`,
			Description: `Inbound IPsec packets

`,
		},

		"limit": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourNameRuleLimit{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Rate limit using a token bucket filter

`,
			Description: `Rate limit using a token bucket filter

`,
		},

		"log_options": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourNameRuleLogOptions{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Log options

`,
			Description: `Log options

`,
		},

		"recent": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourNameRuleRecent{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Parameters for matching recently seen sources

`,
			Description: `Parameters for matching recently seen sources

`,
		},

		"synproxy": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourNameRuleSynproxy{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Synproxy options

`,
			Description: `Synproxy options

`,
		},

		"tcp": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourNameRuleTCP{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `TCP options to match

`,
			Description: `TCP options to match

`,
		},

		"time": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourNameRuleTime{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Time to match rule

`,
			Description: `Time to match rule

`,
		},

		"icmp": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourNameRuleIcmp{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `ICMP type and code information

`,
			Description: `ICMP type and code information

`,
		},

		"ttl": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourNameRuleTTL{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Time to live limit

`,
			Description: `Time to live limit

`,
		},

		"destination": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourNameRuleDestination{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Destination parameters

`,
			Description: `Destination parameters

`,
		},

		"source": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourNameRuleSource{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Source parameters

`,
			Description: `Source parameters

`,
		},

		"inbound_interface": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourNameRuleInboundInterface{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Match inbound-interface

`,
			Description: `Match inbound-interface

`,
		},

		"outbound_interface": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourNameRuleOutboundInterface{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Match outbound-interface

`,
			Description: `Match outbound-interface

`,
		},
	}
}
