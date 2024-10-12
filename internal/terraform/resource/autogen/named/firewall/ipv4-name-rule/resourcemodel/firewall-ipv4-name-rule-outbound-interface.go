// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvfourNameRuleOutboundInterface{}

// FirewallIPvfourNameRuleOutboundInterface describes the resource data model.
type FirewallIPvfourNameRuleOutboundInterface struct {
	// LeafNodes
	LeafFirewallIPvfourNameRuleOutboundInterfaceName  types.String `tfsdk:"name" vyos:"name,omitempty"`
	LeafFirewallIPvfourNameRuleOutboundInterfaceGroup types.String `tfsdk:"group" vyos:"group,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvfourNameRuleOutboundInterface) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"name": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match interface

    |  Format  |  Description                       |
    |----------|------------------------------------|
    |  txt     |  Interface name                    |
    |  txt&    |  Interface name with wildcard      |
    |  !txt    |  Inverted interface name to match  |
`,
			Description: `Match interface

    |  Format  |  Description                       |
    |----------|------------------------------------|
    |  txt     |  Interface name                    |
    |  txt&    |  Interface name with wildcard      |
    |  !txt    |  Inverted interface name to match  |
`,
		},

		"group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match interface-group

    |  Format  |  Description                             |
    |----------|------------------------------------------|
    |  txt     |  Interface-group name to match           |
    |  !txt    |  Inverted interface-group name to match  |
`,
			Description: `Match interface-group

    |  Format  |  Description                             |
    |----------|------------------------------------------|
    |  txt     |  Interface-group name to match           |
    |  !txt    |  Inverted interface-group name to match  |
`,
		},

		// Nodes

	}
}
