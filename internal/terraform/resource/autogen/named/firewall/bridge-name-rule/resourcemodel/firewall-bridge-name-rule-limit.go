// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallBrIDgeNameRuleLimit{}

// FirewallBrIDgeNameRuleLimit describes the resource data model.
type FirewallBrIDgeNameRuleLimit struct {
	// LeafNodes
	LeafFirewallBrIDgeNameRuleLimitBurst types.Number `tfsdk:"burst" vyos:"burst,omitempty"`
	LeafFirewallBrIDgeNameRuleLimitRate  types.String `tfsdk:"rate" vyos:"rate,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallBrIDgeNameRuleLimit) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"burst": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum number of packets to allow in excess of rate

    |  Format        |  Description                                           |
    |----------------|--------------------------------------------------------|
    |  0-4294967295  |  Maximum number of packets to allow in excess of rate  |
`,
			Description: `Maximum number of packets to allow in excess of rate

    |  Format        |  Description                                           |
    |----------------|--------------------------------------------------------|
    |  0-4294967295  |  Maximum number of packets to allow in excess of rate  |
`,
		},

		"rate": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum average matching rate

    |  Format  |  Description                       |
    |----------|------------------------------------|
    |  txt     |  integer/unit (Example: 5/minute)  |
`,
			Description: `Maximum average matching rate

    |  Format  |  Description                       |
    |----------|------------------------------------|
    |  txt     |  integer/unit (Example: 5/minute)  |
`,
		},

		// Nodes

	}
}
