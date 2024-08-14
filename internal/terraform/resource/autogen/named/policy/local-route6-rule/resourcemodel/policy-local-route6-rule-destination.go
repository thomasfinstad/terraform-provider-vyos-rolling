// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &PolicyLocalRoutesixRuleDestination{}

// PolicyLocalRoutesixRuleDestination describes the resource data model.
type PolicyLocalRoutesixRuleDestination struct {
	// LeafNodes
	LeafPolicyLocalRoutesixRuleDestinationAddress types.List   `tfsdk:"address" vyos:"address,omitempty"`
	LeafPolicyLocalRoutesixRuleDestinationPort    types.Number `tfsdk:"port" vyos:"port,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyLocalRoutesixRuleDestination) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `IPv6 address or prefix

    |  Format   |  Description               |
    |-----------|----------------------------|
    |  ipv6     |  Address to match against  |
    |  ipv6net  |  Prefix to match against   |
`,
			Description: `IPv6 address or prefix

    |  Format   |  Description               |
    |-----------|----------------------------|
    |  ipv6     |  Address to match against  |
    |  ipv6net  |  Prefix to match against   |
`,
		},

		"port": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Port number used by connection

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Numeric IP port  |
`,
			Description: `Port number used by connection

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Numeric IP port  |
`,
		},

		// Nodes

	}
}