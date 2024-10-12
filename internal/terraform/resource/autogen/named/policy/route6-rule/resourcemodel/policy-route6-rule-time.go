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

var _ helpers.VyosResourceDataModel = &PolicyRoutesixRuleTime{}

// PolicyRoutesixRuleTime describes the resource data model.
type PolicyRoutesixRuleTime struct {
	// LeafNodes
	LeafPolicyRoutesixRuleTimeMonthdays types.String `tfsdk:"monthdays" vyos:"monthdays,omitempty"`
	LeafPolicyRoutesixRuleTimeStartdate types.String `tfsdk:"startdate" vyos:"startdate,omitempty"`
	LeafPolicyRoutesixRuleTimeStarttime types.String `tfsdk:"starttime" vyos:"starttime,omitempty"`
	LeafPolicyRoutesixRuleTimeStopdate  types.String `tfsdk:"stopdate" vyos:"stopdate,omitempty"`
	LeafPolicyRoutesixRuleTimeStoptime  types.String `tfsdk:"stoptime" vyos:"stoptime,omitempty"`
	LeafPolicyRoutesixRuleTimeUtc       types.Bool   `tfsdk:"utc" vyos:"utc,omitempty"`
	LeafPolicyRoutesixRuleTimeWeekdays  types.String `tfsdk:"weekdays" vyos:"weekdays,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRoutesixRuleTime) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"monthdays": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Monthdays to match rule on

`,
			Description: `Monthdays to match rule on

`,
		},

		"startdate": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Date to start matching rule

`,
			Description: `Date to start matching rule

`,
		},

		"starttime": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Time of day to start matching rule

`,
			Description: `Time of day to start matching rule

`,
		},

		"stopdate": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Date to stop matching rule

`,
			Description: `Date to stop matching rule

`,
		},

		"stoptime": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Time of day to stop matching rule

`,
			Description: `Time of day to stop matching rule

`,
		},

		"utc": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Interpret times for startdate, stopdate, starttime and stoptime to be UTC

`,
			Description: `Interpret times for startdate, stopdate, starttime and stoptime to be UTC

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"weekdays": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Weekdays to match rule on

`,
			Description: `Weekdays to match rule on

`,
		},

		// Nodes

	}
}
