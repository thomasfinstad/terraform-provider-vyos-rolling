// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// PolicyRouteMapRuleOnMatch describes the resource data model.
type PolicyRouteMapRuleOnMatch struct {
	// LeafNodes
	PolicyRouteMapRuleOnMatchGoto customtypes.CustomStringValue `tfsdk:"goto" json:"goto,omitempty"`
	PolicyRouteMapRuleOnMatchNext customtypes.CustomStringValue `tfsdk:"next" json:"next,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o PolicyRouteMapRuleOnMatch) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"goto": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Rule number to goto on match

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Rule number  |

`,
		},

		"next": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Next sequence number to goto on match

`,
		},

		// TagNodes

		// Nodes

	}
}
