// Package resourcemodel code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvsixForwardFilterRuleTime{}

// FirewallIPvsixForwardFilterRuleTime describes the resource data model.
type FirewallIPvsixForwardFilterRuleTime struct {
	// LeafNodes
	LeafFirewallIPvsixForwardFilterRuleTimeStartdate types.String `tfsdk:"startdate" vyos:"startdate,omitempty"`
	LeafFirewallIPvsixForwardFilterRuleTimeStarttime types.String `tfsdk:"starttime" vyos:"starttime,omitempty"`
	LeafFirewallIPvsixForwardFilterRuleTimeStopdate  types.String `tfsdk:"stopdate" vyos:"stopdate,omitempty"`
	LeafFirewallIPvsixForwardFilterRuleTimeStoptime  types.String `tfsdk:"stoptime" vyos:"stoptime,omitempty"`
	LeafFirewallIPvsixForwardFilterRuleTimeWeekdays  types.String `tfsdk:"weekdays" vyos:"weekdays,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixForwardFilterRuleTime) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"startdate": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Date to start matching rule

    |  Format  |  Description                                       |
    |----------|----------------------------------------------------|
    |  txt     |  Enter date using following notation - YYYY-MM-DD  |
`,
			Description: `Date to start matching rule

    |  Format  |  Description                                       |
    |----------|----------------------------------------------------|
    |  txt     |  Enter date using following notation - YYYY-MM-DD  |
`,
		},

		"starttime": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Time of day to start matching rule

    |  Format  |  Description                                         |
    |----------|------------------------------------------------------|
    |  txt     |  Enter time using using 24 hour notation - hh:mm:ss  |
`,
			Description: `Time of day to start matching rule

    |  Format  |  Description                                         |
    |----------|------------------------------------------------------|
    |  txt     |  Enter time using using 24 hour notation - hh:mm:ss  |
`,
		},

		"stopdate": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Date to stop matching rule

    |  Format  |  Description                                       |
    |----------|----------------------------------------------------|
    |  txt     |  Enter date using following notation - YYYY-MM-DD  |
`,
			Description: `Date to stop matching rule

    |  Format  |  Description                                       |
    |----------|----------------------------------------------------|
    |  txt     |  Enter date using following notation - YYYY-MM-DD  |
`,
		},

		"stoptime": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Time of day to stop matching rule

    |  Format  |  Description                                         |
    |----------|------------------------------------------------------|
    |  txt     |  Enter time using using 24 hour notation - hh:mm:ss  |
`,
			Description: `Time of day to stop matching rule

    |  Format  |  Description                                         |
    |----------|------------------------------------------------------|
    |  txt     |  Enter time using using 24 hour notation - hh:mm:ss  |
`,
		},

		"weekdays": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Comma separated weekdays to match rule on

    |  Format  |  Description                                                                    |
    |----------|---------------------------------------------------------------------------------|
    |  txt     |  Name of day (Monday, Tuesday, Wednesday, Thursdays, Friday, Saturday, Sunday)  |
    |  0-6     |  Day number (0 = Sunday ... 6 = Saturday)                                       |
`,
			Description: `Comma separated weekdays to match rule on

    |  Format  |  Description                                                                    |
    |----------|---------------------------------------------------------------------------------|
    |  txt     |  Name of day (Monday, Tuesday, Wednesday, Thursdays, Friday, Saturday, Sunday)  |
    |  0-6     |  Day number (0 = Sunday ... 6 = Saturday)                                       |
`,
		},

		// Nodes

	}
}
