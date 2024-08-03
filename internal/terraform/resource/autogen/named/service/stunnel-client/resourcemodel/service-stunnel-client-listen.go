// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &ServiceStunnelClientListen{}

// ServiceStunnelClientListen describes the resource data model.
type ServiceStunnelClientListen struct {
	// LeafNodes
	LeafServiceStunnelClientListenAddress types.String `tfsdk:"address" vyos:"address,omitempty"`
	LeafServiceStunnelClientListenPort    types.Number `tfsdk:"port" vyos:"port,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceStunnelClientListen) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Hostname or IP address

    |  Format    |  Description   |
    |------------|----------------|
    |  ipv4      |  IPv4 address  |
    |  hostname  |  hostname      |
`,
			Description: `Hostname or IP address

    |  Format    |  Description   |
    |------------|----------------|
    |  ipv4      |  IPv4 address  |
    |  hostname  |  hostname      |
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
