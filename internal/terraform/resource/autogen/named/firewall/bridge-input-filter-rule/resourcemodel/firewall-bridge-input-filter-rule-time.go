/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (time) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallBrIDgeInputFilterRuleTime{}

// FirewallBrIDgeInputFilterRuleTime describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type FirewallBrIDgeInputFilterRuleTime struct {
	// LeafNodes
	LeafFirewallBrIDgeInputFilterRuleTimeStartdate types.String `tfsdk:"startdate" vyos:"startdate,omitempty"`
	LeafFirewallBrIDgeInputFilterRuleTimeStarttime types.String `tfsdk:"starttime" vyos:"starttime,omitempty"`
	LeafFirewallBrIDgeInputFilterRuleTimeStopdate  types.String `tfsdk:"stopdate" vyos:"stopdate,omitempty"`
	LeafFirewallBrIDgeInputFilterRuleTimeStoptime  types.String `tfsdk:"stoptime" vyos:"stoptime,omitempty"`
	LeafFirewallBrIDgeInputFilterRuleTimeWeekdays  types.String `tfsdk:"weekdays" vyos:"weekdays,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallBrIDgeInputFilterRuleTime) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"startdate":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (startdate) */
		schema.StringAttribute{
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

		"starttime":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (starttime) */
		schema.StringAttribute{
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

		"stopdate":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (stopdate) */
		schema.StringAttribute{
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

		"stoptime":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (stoptime) */
		schema.StringAttribute{
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

		"weekdays":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (weekdays) */
		schema.StringAttribute{
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

		// TagNodes

		// Nodes

	}
}
