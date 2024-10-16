// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvsixOutputRawRuleTCPFlags{}

// FirewallIPvsixOutputRawRuleTCPFlags describes the resource data model.
type FirewallIPvsixOutputRawRuleTCPFlags struct {
	// LeafNodes
	LeafFirewallIPvsixOutputRawRuleTCPFlagsSyn types.Bool `tfsdk:"syn" vyos:"syn,omitempty"`
	LeafFirewallIPvsixOutputRawRuleTCPFlagsAck types.Bool `tfsdk:"ack" vyos:"ack,omitempty"`
	LeafFirewallIPvsixOutputRawRuleTCPFlagsFin types.Bool `tfsdk:"fin" vyos:"fin,omitempty"`
	LeafFirewallIPvsixOutputRawRuleTCPFlagsRst types.Bool `tfsdk:"rst" vyos:"rst,omitempty"`
	LeafFirewallIPvsixOutputRawRuleTCPFlagsUrg types.Bool `tfsdk:"urg" vyos:"urg,omitempty"`
	LeafFirewallIPvsixOutputRawRuleTCPFlagsPsh types.Bool `tfsdk:"psh" vyos:"psh,omitempty"`
	LeafFirewallIPvsixOutputRawRuleTCPFlagsEcn types.Bool `tfsdk:"ecn" vyos:"ecn,omitempty"`
	LeafFirewallIPvsixOutputRawRuleTCPFlagsCwr types.Bool `tfsdk:"cwr" vyos:"cwr,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeFirewallIPvsixOutputRawRuleTCPFlagsNot *FirewallIPvsixOutputRawRuleTCPFlagsNot `tfsdk:"not" vyos:"not,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixOutputRawRuleTCPFlags) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"syn": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Synchronise flag

`,
			Description: `Synchronise flag

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"ack": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Acknowledge flag

`,
			Description: `Acknowledge flag

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"fin": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Finish flag

`,
			Description: `Finish flag

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"rst": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Reset flag

`,
			Description: `Reset flag

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"urg": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Urgent flag

`,
			Description: `Urgent flag

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"psh": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Push flag

`,
			Description: `Push flag

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"ecn": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Explicit Congestion Notification flag

`,
			Description: `Explicit Congestion Notification flag

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"cwr": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Congestion Window Reduced flag

`,
			Description: `Congestion Window Reduced flag

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

		"not": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixOutputRawRuleTCPFlagsNot{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Match flags not set

`,
			Description: `Match flags not set

`,
		},
	}
}
