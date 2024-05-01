// Package resourcemodel code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &NatStaticRuleDestination{}

// NatStaticRuleDestination describes the resource data model.
type NatStaticRuleDestination struct {
	// LeafNodes
	LeafNatStaticRuleDestinationAddress types.String `tfsdk:"address" vyos:"address,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o NatStaticRuleDestination) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address, prefix

    |  Format   |  Description            |
    |-----------|-------------------------|
    |  ipv4     |  IPv4 address to match  |
    |  ipv4net  |  IPv4 prefix to match   |
`,
			Description: `IP address, prefix

    |  Format   |  Description            |
    |-----------|-------------------------|
    |  ipv4     |  IPv4 address to match  |
    |  ipv4net  |  IPv4 prefix to match   |
`,
		},

		// Nodes

	}
}
