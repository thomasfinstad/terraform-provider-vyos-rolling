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

var _ helpers.VyosResourceDataModel = &InterfacesOpenvpnServerClient{}

// InterfacesOpenvpnServerClient describes the resource data model.
type InterfacesOpenvpnServerClient struct {
	// LeafNodes
	LeafInterfacesOpenvpnServerClientDisable   types.Bool `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafInterfacesOpenvpnServerClientIP        types.List `tfsdk:"ip" vyos:"ip,omitempty"`
	LeafInterfacesOpenvpnServerClientPushRoute types.List `tfsdk:"push_route" vyos:"push-route,omitempty"`
	LeafInterfacesOpenvpnServerClientSubnet    types.List `tfsdk:"subnet" vyos:"subnet,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesOpenvpnServerClient) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
			Description: `Disable instance

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"ip": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `IP address of the client

    |  Format  |  Description          |
    |----------|-----------------------|
    |  ipv4    |  Client IPv4 address  |
    |  ipv6    |  Client IPv6 address  |
`,
			Description: `IP address of the client

    |  Format  |  Description          |
    |----------|-----------------------|
    |  ipv4    |  Client IPv4 address  |
    |  ipv6    |  Client IPv6 address  |
`,
		},

		"push_route": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Route to be pushed to the client

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  ipv4net  |  IPv4 network and prefix length  |
    |  ipv6net  |  IPv6 network and prefix length  |
`,
			Description: `Route to be pushed to the client

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  ipv4net  |  IPv4 network and prefix length  |
    |  ipv6net  |  IPv6 network and prefix length  |
`,
		},

		"subnet": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Subnet belonging to the client (iroute)

    |  Format   |  Description                                             |
    |-----------|----------------------------------------------------------|
    |  ipv4net  |  IPv4 network and prefix length belonging to the client  |
    |  ipv6net  |  IPv6 network and prefix length belonging to the client  |
`,
			Description: `Subnet belonging to the client (iroute)

    |  Format   |  Description                                             |
    |-----------|----------------------------------------------------------|
    |  ipv4net  |  IPv4 network and prefix length belonging to the client  |
    |  ipv6net  |  IPv6 network and prefix length belonging to the client  |
`,
		},

		// Nodes

	}
}
