// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// FirewallIPvsixOutputFilterRuleRecent describes the resource data model.
type FirewallIPvsixOutputFilterRuleRecent struct {
	// LeafNodes
	LeafFirewallIPvsixOutputFilterRuleRecentCount types.Number `tfsdk:"count" vyos:"count,omitempty"`
	LeafFirewallIPvsixOutputFilterRuleRecentTime  types.String `tfsdk:"time" vyos:"time,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixOutputFilterRuleRecent) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"count": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Source addresses seen more than N times

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-255  &emsp; |  Source addresses seen more than N times  |

`,
		},

		"time": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Source addresses seen in the last second/minute/hour

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  second  &emsp; |  Source addresses seen COUNT times in the last second  |
    |  minute  &emsp; |  Source addresses seen COUNT times in the last minute  |
    |  hour  &emsp; |  Source addresses seen COUNT times in the last hour  |

`,
		},

		// Nodes

	}
}
