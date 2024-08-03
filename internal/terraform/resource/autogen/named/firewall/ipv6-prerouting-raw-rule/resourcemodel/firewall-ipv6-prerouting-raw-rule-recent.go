// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvsixPreroutingRawRuleRecent{}

// FirewallIPvsixPreroutingRawRuleRecent describes the resource data model.
type FirewallIPvsixPreroutingRawRuleRecent struct {
	// LeafNodes
	LeafFirewallIPvsixPreroutingRawRuleRecentCount types.Number `tfsdk:"count" vyos:"count,omitempty"`
	LeafFirewallIPvsixPreroutingRawRuleRecentTime  types.String `tfsdk:"time" vyos:"time,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixPreroutingRawRuleRecent) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"count": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Source addresses seen more than N times

    |  Format  |  Description                              |
    |----------|-------------------------------------------|
    |  1-255   |  Source addresses seen more than N times  |
`,
			Description: `Source addresses seen more than N times

    |  Format  |  Description                              |
    |----------|-------------------------------------------|
    |  1-255   |  Source addresses seen more than N times  |
`,
		},

		"time": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Source addresses seen in the last second/minute/hour

    |  Format  |  Description                                           |
    |----------|--------------------------------------------------------|
    |  second  |  Source addresses seen COUNT times in the last second  |
    |  minute  |  Source addresses seen COUNT times in the last minute  |
    |  hour    |  Source addresses seen COUNT times in the last hour    |
`,
			Description: `Source addresses seen in the last second/minute/hour

    |  Format  |  Description                                           |
    |----------|--------------------------------------------------------|
    |  second  |  Source addresses seen COUNT times in the last second  |
    |  minute  |  Source addresses seen COUNT times in the last minute  |
    |  hour    |  Source addresses seen COUNT times in the last hour    |
`,
		},

		// Nodes

	}
}
