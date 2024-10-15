// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
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

var _ helpers.VyosTopResourceDataModel = &PolicyRouteRule{}

// PolicyRouteRule describes the resource data model.
type PolicyRouteRule struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafPolicyRouteRuleAction              types.String `tfsdk:"action" vyos:"action,omitempty"`
	LeafPolicyRouteRuleDescrIPtion         types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafPolicyRouteRuleMark                types.String `tfsdk:"mark" vyos:"mark,omitempty"`
	LeafPolicyRouteRuleDisable             types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafPolicyRouteRuleLog                 types.Bool   `tfsdk:"log" vyos:"log,omitempty"`
	LeafPolicyRouteRuleProtocol            types.String `tfsdk:"protocol" vyos:"protocol,omitempty"`
	LeafPolicyRouteRuleState               types.List   `tfsdk:"state" vyos:"state,omitempty"`
	LeafPolicyRouteRuleDscp                types.List   `tfsdk:"dscp" vyos:"dscp,omitempty"`
	LeafPolicyRouteRuleDscpExclude         types.List   `tfsdk:"dscp_exclude" vyos:"dscp-exclude,omitempty"`
	LeafPolicyRouteRulePacketLength        types.List   `tfsdk:"packet_length" vyos:"packet-length,omitempty"`
	LeafPolicyRouteRulePacketLengthExclude types.List   `tfsdk:"packet_length_exclude" vyos:"packet-length-exclude,omitempty"`
	LeafPolicyRouteRulePacketType          types.String `tfsdk:"packet_type" vyos:"packet-type,omitempty"`
	LeafPolicyRouteRuleConnectionMark      types.List   `tfsdk:"connection_mark" vyos:"connection-mark,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodePolicyRouteRuleDestination *PolicyRouteRuleDestination `tfsdk:"destination" vyos:"destination,omitempty"`
	NodePolicyRouteRuleSource      *PolicyRouteRuleSource      `tfsdk:"source" vyos:"source,omitempty"`
	NodePolicyRouteRuleFragment    *PolicyRouteRuleFragment    `tfsdk:"fragment" vyos:"fragment,omitempty"`
	NodePolicyRouteRuleIPsec       *PolicyRouteRuleIPsec       `tfsdk:"ipsec" vyos:"ipsec,omitempty"`
	NodePolicyRouteRuleLimit       *PolicyRouteRuleLimit       `tfsdk:"limit" vyos:"limit,omitempty"`
	NodePolicyRouteRuleRecent      *PolicyRouteRuleRecent      `tfsdk:"recent" vyos:"recent,omitempty"`
	NodePolicyRouteRuleSet         *PolicyRouteRuleSet         `tfsdk:"set" vyos:"set,omitempty"`
	NodePolicyRouteRuleTCP         *PolicyRouteRuleTCP         `tfsdk:"tcp" vyos:"tcp,omitempty"`
	NodePolicyRouteRuleTime        *PolicyRouteRuleTime        `tfsdk:"time" vyos:"time,omitempty"`
	NodePolicyRouteRuleIcmp        *PolicyRouteRuleIcmp        `tfsdk:"icmp" vyos:"icmp,omitempty"`
	NodePolicyRouteRuleTTL         *PolicyRouteRuleTTL         `tfsdk:"ttl" vyos:"ttl,omitempty"`
}

// SetID configures the resource ID
func (o *PolicyRouteRule) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *PolicyRouteRule) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *PolicyRouteRule) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *PolicyRouteRule) GetVyosPath() []string {
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
func (o *PolicyRouteRule) GetVyosParentPath() []string {
	return []string{
		"policy",

		"route",

		o.SelfIdentifier.Attributes()["route"].(types.String).ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *PolicyRouteRule) GetVyosNamedParentPath() []string {
	return []string{
		"policy",

		"route",

		o.SelfIdentifier.Attributes()["route"].(types.String).ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteRule) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
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

				"route": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Policy route rule set name for IPv4

`,
					Description: `Policy route rule set name for IPv4

`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in route, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_]+$`),
								"illegal character in  route, value must match: ^[.:a-zA-Z0-9-_]+$",
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
			Attributes: PolicyRouteRuleDestination{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Destination parameters

`,
			Description: `Destination parameters

`,
		},

		"source": schema.SingleNestedAttribute{
			Attributes: PolicyRouteRuleSource{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Source parameters

`,
			Description: `Source parameters

`,
		},

		"fragment": schema.SingleNestedAttribute{
			Attributes: PolicyRouteRuleFragment{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IP fragment match

`,
			Description: `IP fragment match

`,
		},

		"ipsec": schema.SingleNestedAttribute{
			Attributes: PolicyRouteRuleIPsec{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPsec encapsulated packets

`,
			Description: `IPsec encapsulated packets

`,
		},

		"limit": schema.SingleNestedAttribute{
			Attributes: PolicyRouteRuleLimit{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Rate limit using a token bucket filter

`,
			Description: `Rate limit using a token bucket filter

`,
		},

		"recent": schema.SingleNestedAttribute{
			Attributes: PolicyRouteRuleRecent{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Parameters for matching recently seen sources

`,
			Description: `Parameters for matching recently seen sources

`,
		},

		"set": schema.SingleNestedAttribute{
			Attributes: PolicyRouteRuleSet{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Packet modifications

`,
			Description: `Packet modifications

`,
		},

		"tcp": schema.SingleNestedAttribute{
			Attributes: PolicyRouteRuleTCP{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `TCP options to match

`,
			Description: `TCP options to match

`,
		},

		"time": schema.SingleNestedAttribute{
			Attributes: PolicyRouteRuleTime{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Time to match rule

`,
			Description: `Time to match rule

`,
		},

		"icmp": schema.SingleNestedAttribute{
			Attributes: PolicyRouteRuleIcmp{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `ICMP type and code information

`,
			Description: `ICMP type and code information

`,
		},

		"ttl": schema.SingleNestedAttribute{
			Attributes: PolicyRouteRuleTTL{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Time to live limit

`,
			Description: `Time to live limit

`,
		},
	}
}
