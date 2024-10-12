// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &QosPolicyRoundRobinClassMatchIPDestination{}

// QosPolicyRoundRobinClassMatchIPDestination describes the resource data model.
type QosPolicyRoundRobinClassMatchIPDestination struct {
	// LeafNodes
	LeafQosPolicyRoundRobinClassMatchIPDestinationAddress types.String `tfsdk:"address" vyos:"address,omitempty"`
	LeafQosPolicyRoundRobinClassMatchIPDestinationPort    types.Number `tfsdk:"port" vyos:"port,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyRoundRobinClassMatchIPDestination) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv4 destination address for this match

    |  Format   |  Description   |
    |-----------|----------------|
    |  ipv4     |  IPv4 address  |
    |  ipv4net  |  IPv4 prefix   |
`,
			Description: `IPv4 destination address for this match

    |  Format   |  Description   |
    |-----------|----------------|
    |  ipv4     |  IPv4 address  |
    |  ipv4net  |  IPv4 prefix   |
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
