/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallBrIDgePreroutingFilterRuleTTL{}

// FirewallBrIDgePreroutingFilterRuleTTL describes the resource data model.
type FirewallBrIDgePreroutingFilterRuleTTL struct {
	// LeafNodes
	LeafFirewallBrIDgePreroutingFilterRuleTTLEq types.Number `tfsdk:"eq" vyos:"eq,omitempty"`
	LeafFirewallBrIDgePreroutingFilterRuleTTLGt types.Number `tfsdk:"gt" vyos:"gt,omitempty"`
	LeafFirewallBrIDgePreroutingFilterRuleTTLLt types.Number `tfsdk:"lt" vyos:"lt,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallBrIDgePreroutingFilterRuleTTL) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"eq":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
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

		// Nodes

	}
}
