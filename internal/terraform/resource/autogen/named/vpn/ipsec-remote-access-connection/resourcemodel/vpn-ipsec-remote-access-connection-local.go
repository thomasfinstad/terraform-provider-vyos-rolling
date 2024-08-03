// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VpnIPsecRemoteAccessConnectionLocal{}

// VpnIPsecRemoteAccessConnectionLocal describes the resource data model.
type VpnIPsecRemoteAccessConnectionLocal struct {
	// LeafNodes
	LeafVpnIPsecRemoteAccessConnectionLocalPort   types.Number `tfsdk:"port" vyos:"port,omitempty"`
	LeafVpnIPsecRemoteAccessConnectionLocalPrefix types.List   `tfsdk:"prefix" vyos:"prefix,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnIPsecRemoteAccessConnectionLocal) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

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

		"prefix": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Local IPv4 or IPv6 prefix

    |  Format   |  Description        |
    |-----------|---------------------|
    |  ipv4net  |  Local IPv4 prefix  |
    |  ipv6net  |  Local IPv6 prefix  |
`,
			Description: `Local IPv4 or IPv6 prefix

    |  Format   |  Description        |
    |-----------|---------------------|
    |  ipv4net  |  Local IPv4 prefix  |
    |  ipv6net  |  Local IPv6 prefix  |
`,
		},

		// Nodes

	}
}
