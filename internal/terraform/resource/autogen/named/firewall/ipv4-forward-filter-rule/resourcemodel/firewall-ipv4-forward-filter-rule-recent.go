// Package resourcemodel code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvfourForwardFilterRuleRecent{}

// FirewallIPvfourForwardFilterRuleRecent describes the resource data model.
type FirewallIPvfourForwardFilterRuleRecent struct {
	// LeafNodes
	LeafFirewallIPvfourForwardFilterRuleRecentCount types.Number `tfsdk:"count" vyos:"count,omitempty"`
	LeafFirewallIPvfourForwardFilterRuleRecentTime  types.String `tfsdk:"time" vyos:"time,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvfourForwardFilterRuleRecent) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
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
