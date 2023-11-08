// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// FirewallIPvsixName describes the resource data model.
type FirewallIPvsixName struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"ipv6_name_id" vyos:",self-id"`

	// LeafNodes
	LeafFirewallIPvsixNameDefaultAction     types.String `tfsdk:"default_action" vyos:"default-action,omitempty"`
	LeafFirewallIPvsixNameEnableDefaultLog  types.Bool   `tfsdk:"enable_default_log" vyos:"enable-default-log,omitempty"`
	LeafFirewallIPvsixNameDescrIPtion       types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafFirewallIPvsixNameDefaultJumpTarget types.String `tfsdk:"default_jump_target" vyos:"default-jump-target,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagFirewallIPvsixNameRule bool `tfsdk:"rule" vyos:"rule,child"`

	// Nodes
}

// GetID returns the resource ID
func (o FirewallIPvsixName) GetID() *types.String {
	return &o.ID
}

// SetID configures the resource ID
func (o FirewallIPvsixName) SetID(id types.String) {
	o.ID = id
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *FirewallIPvsixName) GetVyosPath() []string {
	return []string{
		"firewall",

		"ipv6-name",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixName) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"ipv6_name_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `IPv6 firewall rule-set name

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"default_action": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Default-action for rule-set

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  drop  &emsp; |  Drop if no prior rules are hit  |
    |  jump  &emsp; |  Jump to another chain if no prior rules are hit  |
    |  reject  &emsp; |  Drop and notify source if no prior rules are hit  |
    |  return  &emsp; |  Return from the current chain and continue at the next rule of the last chain  |
    |  accept  &emsp; |  Accept if no prior rules are hit  |

`,

			// Default:          stringdefault.StaticString(`drop`),
			Computed: true,
		},

		"enable_default_log": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Log packets hitting default-action

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Description  |

`,
		},

		"default_jump_target": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set jump target. Action jump must be defined in default-action to use this setting

`,
		},

		// Nodes

	}
}
