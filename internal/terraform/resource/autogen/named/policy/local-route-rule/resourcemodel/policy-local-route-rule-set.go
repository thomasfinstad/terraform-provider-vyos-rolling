// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &PolicyLocalRouteRuleSet{}

// PolicyLocalRouteRuleSet describes the resource data model.
type PolicyLocalRouteRuleSet struct {
	// LeafNodes
	LeafPolicyLocalRouteRuleSetTable types.Number `tfsdk:"table" vyos:"table,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyLocalRouteRuleSet) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"table": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Routing table to forward packet with

    |  Format  |  Description   |
    |----------|----------------|
    |  1-200   |  Table number  |
`,
			Description: `Routing table to forward packet with

    |  Format  |  Description   |
    |----------|----------------|
    |  1-200   |  Table number  |
`,
		},

		// Nodes

	}
}
