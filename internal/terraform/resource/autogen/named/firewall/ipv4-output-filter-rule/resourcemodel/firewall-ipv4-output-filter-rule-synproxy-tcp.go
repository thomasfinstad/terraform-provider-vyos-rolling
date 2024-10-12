// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvfourOutputFilterRuleSynproxyTCP{}

// FirewallIPvfourOutputFilterRuleSynproxyTCP describes the resource data model.
type FirewallIPvfourOutputFilterRuleSynproxyTCP struct {
	// LeafNodes
	LeafFirewallIPvfourOutputFilterRuleSynproxyTCPMss         types.Number `tfsdk:"mss" vyos:"mss,omitempty"`
	LeafFirewallIPvfourOutputFilterRuleSynproxyTCPWindowScale types.Number `tfsdk:"window_scale" vyos:"window-scale,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvfourOutputFilterRuleSynproxyTCP) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"mss": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `TCP Maximum segment size

    |  Format     |  Description                                    |
    |-------------|-------------------------------------------------|
    |  501-65535  |  Maximum segment size for synproxy connections  |
`,
			Description: `TCP Maximum segment size

    |  Format     |  Description                                    |
    |-------------|-------------------------------------------------|
    |  501-65535  |  Maximum segment size for synproxy connections  |
`,
		},

		"window_scale": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `TCP window scale for synproxy connections

    |  Format  |  Description       |
    |----------|--------------------|
    |  1-14    |  TCP window scale  |
`,
			Description: `TCP window scale for synproxy connections

    |  Format  |  Description       |
    |----------|--------------------|
    |  1-14    |  TCP window scale  |
`,
		},

		// Nodes

	}
}
