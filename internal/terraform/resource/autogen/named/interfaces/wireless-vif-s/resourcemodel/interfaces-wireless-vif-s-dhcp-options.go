// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &InterfacesWirelessVifSDhcpOptions{}

// InterfacesWirelessVifSDhcpOptions describes the resource data model.
type InterfacesWirelessVifSDhcpOptions struct {
	// LeafNodes
	LeafInterfacesWirelessVifSDhcpOptionsClientID             types.String `tfsdk:"client_id" vyos:"client-id,omitempty"`
	LeafInterfacesWirelessVifSDhcpOptionsHostName             types.String `tfsdk:"host_name" vyos:"host-name,omitempty"`
	LeafInterfacesWirelessVifSDhcpOptionsMtu                  types.Bool   `tfsdk:"mtu" vyos:"mtu,omitempty"`
	LeafInterfacesWirelessVifSDhcpOptionsVendorClassID        types.String `tfsdk:"vendor_class_id" vyos:"vendor-class-id,omitempty"`
	LeafInterfacesWirelessVifSDhcpOptionsUserClass            types.String `tfsdk:"user_class" vyos:"user-class,omitempty"`
	LeafInterfacesWirelessVifSDhcpOptionsNoDefaultRoute       types.Bool   `tfsdk:"no_default_route" vyos:"no-default-route,omitempty"`
	LeafInterfacesWirelessVifSDhcpOptionsDefaultRouteDistance types.Number `tfsdk:"default_route_distance" vyos:"default-route-distance,omitempty"`
	LeafInterfacesWirelessVifSDhcpOptionsReject               types.List   `tfsdk:"reject" vyos:"reject,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWirelessVifSDhcpOptions) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"client_id": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Identifier used by client to identify itself to the DHCP server

    |  Format  |  Description         |
    |----------|----------------------|
    |  txt     |  DHCP option string  |
`,
			Description: `Identifier used by client to identify itself to the DHCP server

    |  Format  |  Description         |
    |----------|----------------------|
    |  txt     |  DHCP option string  |
`,
		},

		"host_name": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Override system host-name sent to DHCP server

`,
			Description: `Override system host-name sent to DHCP server

`,
		},

		"mtu": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Use MTU value from DHCP server - ignore interface setting

`,
			Description: `Use MTU value from DHCP server - ignore interface setting

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"vendor_class_id": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Identify the vendor client type to the DHCP server

    |  Format  |  Description         |
    |----------|----------------------|
    |  txt     |  DHCP option string  |
`,
			Description: `Identify the vendor client type to the DHCP server

    |  Format  |  Description         |
    |----------|----------------------|
    |  txt     |  DHCP option string  |
`,
		},

		"user_class": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Identify to the DHCP server, user configurable option

    |  Format  |  Description         |
    |----------|----------------------|
    |  txt     |  DHCP option string  |
`,
			Description: `Identify to the DHCP server, user configurable option

    |  Format  |  Description         |
    |----------|----------------------|
    |  txt     |  DHCP option string  |
`,
		},

		"no_default_route": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not install default route to system

`,
			Description: `Do not install default route to system

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"default_route_distance": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Distance for installed default route

    |  Format  |  Description                                      |
    |----------|---------------------------------------------------|
    |  1-255   |  Distance for the default route from DHCP server  |
`,
			Description: `Distance for installed default route

    |  Format  |  Description                                      |
    |----------|---------------------------------------------------|
    |  1-255   |  Distance for the default route from DHCP server  |
`,

			// Default:          stringdefault.StaticString(`210`),
			Computed: true,
		},

		"reject": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `IP addresses or subnets from which to reject DHCP leases

    |  Format   |  Description            |
    |-----------|-------------------------|
    |  ipv4     |  IPv4 address to match  |
    |  ipv4net  |  IPv4 prefix to match   |
`,
			Description: `IP addresses or subnets from which to reject DHCP leases

    |  Format   |  Description            |
    |-----------|-------------------------|
    |  ipv4     |  IPv4 address to match  |
    |  ipv4net  |  IPv4 prefix to match   |
`,
		},

		// Nodes

	}
}
