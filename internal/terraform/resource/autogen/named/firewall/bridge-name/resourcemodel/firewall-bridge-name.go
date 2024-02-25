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

// FirewallBrIDgeName describes the resource data model.
type FirewallBrIDgeName struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"name_id" vyos:"-,self-id"`

	// LeafNodes
	LeafFirewallBrIDgeNameDefaultAction     types.String `tfsdk:"default_action" vyos:"default-action,omitempty"`
	LeafFirewallBrIDgeNameDefaultLog        types.Bool   `tfsdk:"default_log" vyos:"default-log,omitempty"`
	LeafFirewallBrIDgeNameDescrIPtion       types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafFirewallBrIDgeNameDefaultJumpTarget types.String `tfsdk:"default_jump_target" vyos:"default-jump-target,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagFirewallBrIDgeNameRule bool `tfsdk:"-" vyos:"rule,child"`

	// Nodes
}

// SetID configures the resource ID
func (o *FirewallBrIDgeName) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *FirewallBrIDgeName) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"firewall",

		"bridge",

		"name",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallBrIDgeName) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
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
    |  continue  &emsp; |  Continue parsing next rule  |

`,

			// Default:          stringdefault.StaticString(`drop`),
			Computed: true,
		},

		"default_log": schema.BoolAttribute{
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
