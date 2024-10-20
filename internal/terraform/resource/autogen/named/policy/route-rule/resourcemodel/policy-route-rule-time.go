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

var _ helpers.VyosResourceDataModel = &PolicyRouteRuleTime{}

// PolicyRouteRuleTime describes the resource data model.
type PolicyRouteRuleTime struct {
	// LeafNodes
	LeafPolicyRouteRuleTimeMonthdays types.String `tfsdk:"monthdays" vyos:"monthdays,omitempty"`
	LeafPolicyRouteRuleTimeStartdate types.String `tfsdk:"startdate" vyos:"startdate,omitempty"`
	LeafPolicyRouteRuleTimeStarttime types.String `tfsdk:"starttime" vyos:"starttime,omitempty"`
	LeafPolicyRouteRuleTimeStopdate  types.String `tfsdk:"stopdate" vyos:"stopdate,omitempty"`
	LeafPolicyRouteRuleTimeStoptime  types.String `tfsdk:"stoptime" vyos:"stoptime,omitempty"`
	LeafPolicyRouteRuleTimeUtc       types.Bool   `tfsdk:"utc" vyos:"utc,omitempty"`
	LeafPolicyRouteRuleTimeWeekdays  types.String `tfsdk:"weekdays" vyos:"weekdays,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteRuleTime) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"monthdays":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Monthdays to match rule on

`,
			Description: `Monthdays to match rule on

`,
		},

		"startdate":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Date to start matching rule

`,
			Description: `Date to start matching rule

`,
		},

		"starttime":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Time of day to start matching rule

`,
			Description: `Time of day to start matching rule

`,
		},

		"stopdate":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Date to stop matching rule

`,
			Description: `Date to stop matching rule

`,
		},

		"stoptime":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Time of day to stop matching rule

`,
			Description: `Time of day to stop matching rule

`,
		},

		"utc":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Interpret times for startdate, stopdate, starttime and stoptime to be UTC

`,
			Description: `Interpret times for startdate, stopdate, starttime and stoptime to be UTC

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"weekdays":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Weekdays to match rule on

`,
			Description: `Weekdays to match rule on

`,
		},

		// Nodes

	}
}
