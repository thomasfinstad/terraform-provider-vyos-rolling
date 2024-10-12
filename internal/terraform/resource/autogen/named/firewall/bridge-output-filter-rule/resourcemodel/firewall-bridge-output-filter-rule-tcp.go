// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallBrIDgeOutputFilterRuleTCP{}

// FirewallBrIDgeOutputFilterRuleTCP describes the resource data model.
type FirewallBrIDgeOutputFilterRuleTCP struct {
	// LeafNodes
	LeafFirewallBrIDgeOutputFilterRuleTCPMss types.String `tfsdk:"mss" vyos:"mss,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeFirewallBrIDgeOutputFilterRuleTCPFlags *FirewallBrIDgeOutputFilterRuleTCPFlags `tfsdk:"flags" vyos:"flags,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallBrIDgeOutputFilterRuleTCP) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"mss": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum segment size (MSS)

    |  Format       |  Description                           |
    |---------------|----------------------------------------|
    |  1-16384      |  Maximum segment size                  |
    |  <min>-<max>  |  TCP MSS range (use '-' as delimiter)  |
`,
			Description: `Maximum segment size (MSS)

    |  Format       |  Description                           |
    |---------------|----------------------------------------|
    |  1-16384      |  Maximum segment size                  |
    |  <min>-<max>  |  TCP MSS range (use '-' as delimiter)  |
`,
		},

		// Nodes

		"flags": schema.SingleNestedAttribute{
			Attributes: FirewallBrIDgeOutputFilterRuleTCPFlags{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `TCP flags to match

`,
			Description: `TCP flags to match

`,
		},
	}
}
