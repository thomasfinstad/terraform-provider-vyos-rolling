// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvfourOutputRawRuleTTL{}

// FirewallIPvfourOutputRawRuleTTL describes the resource data model.
type FirewallIPvfourOutputRawRuleTTL struct {
	// LeafNodes
	LeafFirewallIPvfourOutputRawRuleTTLEq types.Number `tfsdk:"eq" vyos:"eq,omitempty"`
	LeafFirewallIPvfourOutputRawRuleTTLGt types.Number `tfsdk:"gt" vyos:"gt,omitempty"`
	LeafFirewallIPvfourOutputRawRuleTTLLt types.Number `tfsdk:"lt" vyos:"lt,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvfourOutputRawRuleTTL) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"eq": schema.NumberAttribute{
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

		"gt": schema.NumberAttribute{
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

		"lt": schema.NumberAttribute{
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
