/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvfourForwardFilterRuleTCPFlagsNot{}

// FirewallIPvfourForwardFilterRuleTCPFlagsNot describes the resource data model.
type FirewallIPvfourForwardFilterRuleTCPFlagsNot struct {
	// LeafNodes
	LeafFirewallIPvfourForwardFilterRuleTCPFlagsNotSyn types.Bool `tfsdk:"syn" vyos:"syn,omitempty"`
	LeafFirewallIPvfourForwardFilterRuleTCPFlagsNotAck types.Bool `tfsdk:"ack" vyos:"ack,omitempty"`
	LeafFirewallIPvfourForwardFilterRuleTCPFlagsNotFin types.Bool `tfsdk:"fin" vyos:"fin,omitempty"`
	LeafFirewallIPvfourForwardFilterRuleTCPFlagsNotRst types.Bool `tfsdk:"rst" vyos:"rst,omitempty"`
	LeafFirewallIPvfourForwardFilterRuleTCPFlagsNotUrg types.Bool `tfsdk:"urg" vyos:"urg,omitempty"`
	LeafFirewallIPvfourForwardFilterRuleTCPFlagsNotPsh types.Bool `tfsdk:"psh" vyos:"psh,omitempty"`
	LeafFirewallIPvfourForwardFilterRuleTCPFlagsNotEcn types.Bool `tfsdk:"ecn" vyos:"ecn,omitempty"`
	LeafFirewallIPvfourForwardFilterRuleTCPFlagsNotCwr types.Bool `tfsdk:"cwr" vyos:"cwr,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvfourForwardFilterRuleTCPFlagsNot) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"syn":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Synchronise flag

`,
			Description: `Synchronise flag

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"ack":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Acknowledge flag

`,
			Description: `Acknowledge flag

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"fin":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Finish flag

`,
			Description: `Finish flag

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"rst":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Reset flag

`,
			Description: `Reset flag

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"urg":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Urgent flag

`,
			Description: `Urgent flag

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"psh":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Push flag

`,
			Description: `Push flag

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"ecn":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Explicit Congestion Notification flag

`,
			Description: `Explicit Congestion Notification flag

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"cwr":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Congestion Window Reduced flag

`,
			Description: `Congestion Window Reduced flag

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
