// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &InterfacesOpenvpnServerClientIPvsixPool{}

// InterfacesOpenvpnServerClientIPvsixPool describes the resource data model.
type InterfacesOpenvpnServerClientIPvsixPool struct {
	// LeafNodes
	LeafInterfacesOpenvpnServerClientIPvsixPoolBase    types.String `tfsdk:"base" vyos:"base,omitempty"`
	LeafInterfacesOpenvpnServerClientIPvsixPoolDisable types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesOpenvpnServerClientIPvsixPool) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"base": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Client IPv6 pool base address with optional prefix length

    |  Format   |  Description                                                                                                                                |
    |-----------|---------------------------------------------------------------------------------------------------------------------------------------------|
    |  ipv6net  |  Client IPv6 pool base address with optional prefix length (defaults: base = server subnet + 0x1000, prefix length = server prefix length)  |
`,
			Description: `Client IPv6 pool base address with optional prefix length

    |  Format   |  Description                                                                                                                                |
    |-----------|---------------------------------------------------------------------------------------------------------------------------------------------|
    |  ipv6net  |  Client IPv6 pool base address with optional prefix length (defaults: base = server subnet + 0x1000, prefix length = server prefix length)  |
`,
		},

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
			Description: `Disable instance

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}