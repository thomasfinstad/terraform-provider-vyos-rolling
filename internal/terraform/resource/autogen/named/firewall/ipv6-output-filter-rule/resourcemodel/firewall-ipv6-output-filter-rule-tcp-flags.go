/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (flags) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvsixOutputFilterRuleTCPFlags{}

// FirewallIPvsixOutputFilterRuleTCPFlags describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type FirewallIPvsixOutputFilterRuleTCPFlags struct {
	// LeafNodes
	LeafFirewallIPvsixOutputFilterRuleTCPFlagsSyn types.Bool `tfsdk:"syn" vyos:"syn,omitempty"`
	LeafFirewallIPvsixOutputFilterRuleTCPFlagsAck types.Bool `tfsdk:"ack" vyos:"ack,omitempty"`
	LeafFirewallIPvsixOutputFilterRuleTCPFlagsFin types.Bool `tfsdk:"fin" vyos:"fin,omitempty"`
	LeafFirewallIPvsixOutputFilterRuleTCPFlagsRst types.Bool `tfsdk:"rst" vyos:"rst,omitempty"`
	LeafFirewallIPvsixOutputFilterRuleTCPFlagsUrg types.Bool `tfsdk:"urg" vyos:"urg,omitempty"`
	LeafFirewallIPvsixOutputFilterRuleTCPFlagsPsh types.Bool `tfsdk:"psh" vyos:"psh,omitempty"`
	LeafFirewallIPvsixOutputFilterRuleTCPFlagsEcn types.Bool `tfsdk:"ecn" vyos:"ecn,omitempty"`
	LeafFirewallIPvsixOutputFilterRuleTCPFlagsCwr types.Bool `tfsdk:"cwr" vyos:"cwr,omitempty"`

	// TagNodes

	// Nodes

	NodeFirewallIPvsixOutputFilterRuleTCPFlagsNot *FirewallIPvsixOutputFilterRuleTCPFlagsNot `tfsdk:"not" vyos:"not,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixOutputFilterRuleTCPFlags) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"syn":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (syn) */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (ack) */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (fin) */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (rst) */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (urg) */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (psh) */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (ecn) */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (cwr) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Congestion Window Reduced flag

`,
			Description: `Congestion Window Reduced flag

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// TagNodes

		// Nodes

		"not": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixOutputFilterRuleTCPFlagsNot{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Match flags not set

`,
			Description: `Match flags not set

`,
		},
	}
}
