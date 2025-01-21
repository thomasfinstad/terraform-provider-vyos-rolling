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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (hop-limit) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallBrIDgeNameRuleHopLimit{}

// FirewallBrIDgeNameRuleHopLimit describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type FirewallBrIDgeNameRuleHopLimit struct {
	// LeafNodes
	LeafFirewallBrIDgeNameRuleHopLimitEq types.Number `tfsdk:"eq" vyos:"eq,omitempty"`
	LeafFirewallBrIDgeNameRuleHopLimitGt types.Number `tfsdk:"gt" vyos:"gt,omitempty"`
	LeafFirewallBrIDgeNameRuleHopLimitLt types.Number `tfsdk:"lt" vyos:"lt,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallBrIDgeNameRuleHopLimit) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"eq":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (eq) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Match on equal value

    |  Format  |  Description     |
    |----------|------------------|
    |  0-255   |  Equal to value  |
`,
			Description: `Match on equal value

    |  Format  |  Description     |
    |----------|------------------|
    |  0-255   |  Equal to value  |
`,
		},

		"gt":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (gt) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Match on greater then value

    |  Format  |  Description         |
    |----------|----------------------|
    |  0-255   |  Greater then value  |
`,
			Description: `Match on greater then value

    |  Format  |  Description         |
    |----------|----------------------|
    |  0-255   |  Greater then value  |
`,
		},

		"lt":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (lt) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Match on less then value

    |  Format  |  Description      |
    |----------|-------------------|
    |  0-255   |  Less then value  |
`,
			Description: `Match on less then value

    |  Format  |  Description      |
    |----------|-------------------|
    |  0-255   |  Less then value  |
`,
		},

		// TagNodes

		// Nodes

	}
}
