// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesWwanDhcpOptions describes the resource data model.
type InterfacesWwanDhcpOptions struct {
	// LeafNodes
	LeafInterfacesWwanDhcpOptionsClientID             types.String `tfsdk:"client_id" vyos:"client-id,omitempty"`
	LeafInterfacesWwanDhcpOptionsHostName             types.String `tfsdk:"host_name" vyos:"host-name,omitempty"`
	LeafInterfacesWwanDhcpOptionsMtu                  types.Bool   `tfsdk:"mtu" vyos:"mtu,omitempty"`
	LeafInterfacesWwanDhcpOptionsVendorClassID        types.String `tfsdk:"vendor_class_id" vyos:"vendor-class-id,omitempty"`
	LeafInterfacesWwanDhcpOptionsNoDefaultRoute       types.Bool   `tfsdk:"no_default_route" vyos:"no-default-route,omitempty"`
	LeafInterfacesWwanDhcpOptionsDefaultRouteDistance types.Number `tfsdk:"default_route_distance" vyos:"default-route-distance,omitempty"`
	LeafInterfacesWwanDhcpOptionsReject               types.List   `tfsdk:"reject" vyos:"reject,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWwanDhcpOptions) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"client_id": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Identifier used by client to identify itself to the DHCP server

`,
		},

		"host_name": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Override system host-name sent to DHCP server

`,
		},

		"mtu": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Use MTU value from DHCP server - ignore interface setting

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"vendor_class_id": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Identify the vendor client type to the DHCP server

`,
		},

		"no_default_route": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not install default route to system

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"default_route_distance": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Distance for installed default route

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-255  |  Distance for the default route from DHCP server  |

`,

			// Default:          stringdefault.StaticString(`210`),
			Computed: true,
		},

		"reject": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `IP addresses or subnets from which to reject DHCP leases

    |  Format  |  Description  |
    |----------|---------------|
    |  ipv4  |  IPv4 address to match  |
    |  ipv4net  |  IPv4 prefix to match  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesWwanDhcpOptions) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesWwanDhcpOptions) UnmarshalJSON(_ []byte) error {
	return nil
}
