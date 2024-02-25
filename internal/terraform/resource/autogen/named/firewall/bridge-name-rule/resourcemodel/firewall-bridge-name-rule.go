// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// FirewallBrIDgeNameRule describes the resource data model.
type FirewallBrIDgeNameRule struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Number `tfsdk:"rule_id" vyos:"-,self-id"`

	ParentIDFirewallBrIDgeName types.String `tfsdk:"name" vyos:"name,parent-id"`

	// LeafNodes
	LeafFirewallBrIDgeNameRuleAction       types.String `tfsdk:"action" vyos:"action,omitempty"`
	LeafFirewallBrIDgeNameRuleQueue        types.Number `tfsdk:"queue" vyos:"queue,omitempty"`
	LeafFirewallBrIDgeNameRuleQueueOptions types.List   `tfsdk:"queue_options" vyos:"queue-options,omitempty"`
	LeafFirewallBrIDgeNameRuleDisable      types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafFirewallBrIDgeNameRuleJumpTarget   types.String `tfsdk:"jump_target" vyos:"jump-target,omitempty"`
	LeafFirewallBrIDgeNameRuleLog          types.Bool   `tfsdk:"log" vyos:"log,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeFirewallBrIDgeNameRuleDestination       *FirewallBrIDgeNameRuleDestination       `tfsdk:"destination" vyos:"destination,omitempty"`
	NodeFirewallBrIDgeNameRuleLogOptions        *FirewallBrIDgeNameRuleLogOptions        `tfsdk:"log_options" vyos:"log-options,omitempty"`
	NodeFirewallBrIDgeNameRuleSource            *FirewallBrIDgeNameRuleSource            `tfsdk:"source" vyos:"source,omitempty"`
	NodeFirewallBrIDgeNameRuleInboundInterface  *FirewallBrIDgeNameRuleInboundInterface  `tfsdk:"inbound_interface" vyos:"inbound-interface,omitempty"`
	NodeFirewallBrIDgeNameRuleOutboundInterface *FirewallBrIDgeNameRuleOutboundInterface `tfsdk:"outbound_interface" vyos:"outbound-interface,omitempty"`
	NodeFirewallBrIDgeNameRuleVlan              *FirewallBrIDgeNameRuleVlan              `tfsdk:"vlan" vyos:"vlan,omitempty"`
}

// SetID configures the resource ID
func (o *FirewallBrIDgeNameRule) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *FirewallBrIDgeNameRule) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"firewall",

		"bridge",

		"name",
		o.ParentIDFirewallBrIDgeName.ValueString(),

		"rule",
		o.SelfIdentifier.ValueBigFloat().String(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallBrIDgeNameRule) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"rule_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Bridge Firewall forward filter rule number

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-999999  &emsp; |  Number for this firewall rule  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		"name_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Bridge custom firewall

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
    |  continue  &emsp; |  Continue parsing next rule  |
    |  jump  &emsp; |  Jump to another chain  |
    |  return  &emsp; |  Return from the current chain and continue at the next rule of the last chain  |
    |  drop  &emsp; |  Drop matching entries  |
    |  queue  &emsp; |  Enqueue packet to userspace  |

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
			MarkdownDescription: `Option to disable firewall rule

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"jump_target": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set jump target. Action jump must be defined to use this setting

`,
		},

		"log": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Log packets hitting this rule

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

		"destination": schema.SingleNestedAttribute{
			Attributes: FirewallBrIDgeNameRuleDestination{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Destination parameters

`,
		},

		"log_options": schema.SingleNestedAttribute{
			Attributes: FirewallBrIDgeNameRuleLogOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Log options

`,
		},

		"source": schema.SingleNestedAttribute{
			Attributes: FirewallBrIDgeNameRuleSource{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Source parameters

`,
		},

		"inbound_interface": schema.SingleNestedAttribute{
			Attributes: FirewallBrIDgeNameRuleInboundInterface{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match inbound-interface

`,
		},

		"outbound_interface": schema.SingleNestedAttribute{
			Attributes: FirewallBrIDgeNameRuleOutboundInterface{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match outbound-interface

`,
		},

		"vlan": schema.SingleNestedAttribute{
			Attributes: FirewallBrIDgeNameRuleVlan{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `VLAN parameters

`,
		},
	}
}
