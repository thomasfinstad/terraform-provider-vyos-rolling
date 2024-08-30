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

var _ helpers.VyosTopResourceDataModel = &PolicyRoutesixRule{}

// PolicyRoutesixRule describes the resource data model.
type PolicyRoutesixRule struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafPolicyRoutesixRuleAction              types.String `tfsdk:"action" vyos:"action,omitempty"`
	LeafPolicyRoutesixRuleDescrIPtion         types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafPolicyRoutesixRuleMark                types.String `tfsdk:"mark" vyos:"mark,omitempty"`
	LeafPolicyRoutesixRuleDisable             types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafPolicyRoutesixRuleLog                 types.Bool   `tfsdk:"log" vyos:"log,omitempty"`
	LeafPolicyRoutesixRuleProtocol            types.String `tfsdk:"protocol" vyos:"protocol,omitempty"`
	LeafPolicyRoutesixRuleState               types.List   `tfsdk:"state" vyos:"state,omitempty"`
	LeafPolicyRoutesixRuleDscp                types.List   `tfsdk:"dscp" vyos:"dscp,omitempty"`
	LeafPolicyRoutesixRuleDscpExclude         types.List   `tfsdk:"dscp_exclude" vyos:"dscp-exclude,omitempty"`
	LeafPolicyRoutesixRulePacketLength        types.List   `tfsdk:"packet_length" vyos:"packet-length,omitempty"`
	LeafPolicyRoutesixRulePacketLengthExclude types.List   `tfsdk:"packet_length_exclude" vyos:"packet-length-exclude,omitempty"`
	LeafPolicyRoutesixRulePacketType          types.String `tfsdk:"packet_type" vyos:"packet-type,omitempty"`
	LeafPolicyRoutesixRuleConnectionMark      types.List   `tfsdk:"connection_mark" vyos:"connection-mark,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodePolicyRoutesixRuleDestination *PolicyRoutesixRuleDestination `tfsdk:"destination" vyos:"destination,omitempty"`
	NodePolicyRoutesixRuleSource      *PolicyRoutesixRuleSource      `tfsdk:"source" vyos:"source,omitempty"`
	NodePolicyRoutesixRuleFragment    *PolicyRoutesixRuleFragment    `tfsdk:"fragment" vyos:"fragment,omitempty"`
	NodePolicyRoutesixRuleIPsec       *PolicyRoutesixRuleIPsec       `tfsdk:"ipsec" vyos:"ipsec,omitempty"`
	NodePolicyRoutesixRuleLimit       *PolicyRoutesixRuleLimit       `tfsdk:"limit" vyos:"limit,omitempty"`
	NodePolicyRoutesixRuleRecent      *PolicyRoutesixRuleRecent      `tfsdk:"recent" vyos:"recent,omitempty"`
	NodePolicyRoutesixRuleSet         *PolicyRoutesixRuleSet         `tfsdk:"set" vyos:"set,omitempty"`
	NodePolicyRoutesixRuleTCP         *PolicyRoutesixRuleTCP         `tfsdk:"tcp" vyos:"tcp,omitempty"`
	NodePolicyRoutesixRuleTime        *PolicyRoutesixRuleTime        `tfsdk:"time" vyos:"time,omitempty"`
	NodePolicyRoutesixRuleIcmpvsix    *PolicyRoutesixRuleIcmpvsix    `tfsdk:"icmpv6" vyos:"icmpv6,omitempty"`
	NodePolicyRoutesixRuleHopLimit    *PolicyRoutesixRuleHopLimit    `tfsdk:"hop_limit" vyos:"hop-limit,omitempty"`
}

// SetID configures the resource ID
func (o *PolicyRoutesixRule) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *PolicyRoutesixRule) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *PolicyRoutesixRule) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *PolicyRoutesixRule) GetVyosPath() []string {
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
func (o *PolicyRoutesixRule) GetVyosParentPath() []string {
	return []string{
		"policy",

		"route6",

		o.SelfIdentifier.Attributes()["route6"].(types.String).ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *PolicyRoutesixRule) GetVyosNamedParentPath() []string {
	return []string{
		"policy",

		"route6",

		o.SelfIdentifier.Attributes()["route6"].(types.String).ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRoutesixRule) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
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
					MarkdownDescription: `Policy rule number

    |  Format    |  Description            |
    |------------|-------------------------|
    |  1-999999  |  Number of policy rule  |
`,
					Description: `Policy rule number

    |  Format    |  Description            |
    |------------|-------------------------|
    |  1-999999  |  Number of policy rule  |
`,
					PlanModifiers: []planmodifier.Number{
						numberplanmodifier.RequiresReplace(),
					},
				},

				"route6": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Policy route rule set name for IPv6

`,
					Description: `Policy route rule set name for IPv6

`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in route6, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
								"illegal character in  route6, value must match: ^[a-zA-Z0-9-_]*$",
							),
						),
					},
				},
			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"action": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Rule action

    |  Format  |  Description                                                                    |
    |----------|---------------------------------------------------------------------------------|
    |  accept  |  Accept matching entries                                                        |
    |  reject  |  Reject matching entries                                                        |
    |  return  |  Return from the current chain and continue at the next rule of the last chain  |
    |  drop    |  Drop matching entries                                                          |
`,
			Description: `Rule action

    |  Format  |  Description                                                                    |
    |----------|---------------------------------------------------------------------------------|
    |  accept  |  Accept matching entries                                                        |
    |  reject  |  Reject matching entries                                                        |
    |  return  |  Return from the current chain and continue at the next rule of the last chain  |
    |  drop    |  Drop matching entries                                                          |
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
    |  !<protocol>  |  IP protocol number  |
`,
			Description: `Protocol to match (protocol name, number, or "all")

    |  Format       |  Description         |
    |---------------|----------------------|
    |  all          |  All IP protocols    |
    |  tcp_udp      |  Both TCP and UDP    |
    |  0-255        |  IP protocol number  |
    |  !<protocol>  |  IP protocol number  |
`,

			// Default:          stringdefault.StaticString(`all`),
			Computed: true,
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

		// Nodes

		"destination": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleDestination{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Destination parameters

`,
			Description: `Destination parameters

`,
		},

		"source": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleSource{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Source parameters

`,
			Description: `Source parameters

`,
		},

		"fragment": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleFragment{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IP fragment match

`,
			Description: `IP fragment match

`,
		},

		"ipsec": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleIPsec{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPsec encapsulated packets

`,
			Description: `IPsec encapsulated packets

`,
		},

		"limit": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleLimit{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Rate limit using a token bucket filter

`,
			Description: `Rate limit using a token bucket filter

`,
		},

		"recent": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleRecent{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Parameters for matching recently seen sources

`,
			Description: `Parameters for matching recently seen sources

`,
		},

		"set": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleSet{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Packet modifications

`,
			Description: `Packet modifications

`,
		},

		"tcp": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleTCP{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `TCP options to match

`,
			Description: `TCP options to match

`,
		},

		"time": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleTime{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Time to match rule

`,
			Description: `Time to match rule

`,
		},

		"icmpv6": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleIcmpvsix{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `ICMPv6 type and code information

`,
			Description: `ICMPv6 type and code information

`,
		},

		"hop_limit": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleHopLimit{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Hop limit

`,
			Description: `Hop limit

`,
		},
	}
}
