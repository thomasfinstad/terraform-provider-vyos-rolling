// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallBrIDgeOutputFilterRuleVlan{}

// FirewallBrIDgeOutputFilterRuleVlan describes the resource data model.
type FirewallBrIDgeOutputFilterRuleVlan struct {
	// LeafNodes
	LeafFirewallBrIDgeOutputFilterRuleVlanID       types.String `tfsdk:"id" vyos:"id,omitempty"`
	LeafFirewallBrIDgeOutputFilterRuleVlanPriority types.String `tfsdk:"priority" vyos:"priority,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallBrIDgeOutputFilterRuleVlan) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"id": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Vlan id

    |  Format       |  Description             |
    |---------------|--------------------------|
    |  0-4096       |  Vlan id                 |
    |  <start-end>  |  Vlan id range to match  |
`,
			Description: `Vlan id

    |  Format       |  Description             |
    |---------------|--------------------------|
    |  0-4096       |  Vlan id                 |
    |  <start-end>  |  Vlan id range to match  |
`,
		},

		"priority": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Vlan priority(pcp)

    |  Format       |  Description                   |
    |---------------|--------------------------------|
    |  0-7          |  Vlan priority                 |
    |  <start-end>  |  Vlan priority range to match  |
`,
			Description: `Vlan priority(pcp)

    |  Format       |  Description                   |
    |---------------|--------------------------------|
    |  0-7          |  Vlan priority                 |
    |  <start-end>  |  Vlan priority range to match  |
`,
		},

		// Nodes

	}
}
