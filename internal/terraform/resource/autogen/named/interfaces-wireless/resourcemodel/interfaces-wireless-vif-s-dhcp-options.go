// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// InterfacesWirelessVifSDhcpOptions describes the resource data model.
type InterfacesWirelessVifSDhcpOptions struct {
	// LeafNodes
	InterfacesWirelessVifSDhcpOptionsClientID             customtypes.CustomStringValue `tfsdk:"client_id" json:"client-id,omitempty"`
	InterfacesWirelessVifSDhcpOptionsHostName             customtypes.CustomStringValue `tfsdk:"host_name" json:"host-name,omitempty"`
	InterfacesWirelessVifSDhcpOptionsMtu                  customtypes.CustomStringValue `tfsdk:"mtu" json:"mtu,omitempty"`
	InterfacesWirelessVifSDhcpOptionsVendorClassID        customtypes.CustomStringValue `tfsdk:"vendor_class_id" json:"vendor-class-id,omitempty"`
	InterfacesWirelessVifSDhcpOptionsNoDefaultRoute       customtypes.CustomStringValue `tfsdk:"no_default_route" json:"no-default-route,omitempty"`
	InterfacesWirelessVifSDhcpOptionsDefaultRouteDistance customtypes.CustomStringValue `tfsdk:"default_route_distance" json:"default-route-distance,omitempty"`
	InterfacesWirelessVifSDhcpOptionsReject               customtypes.CustomStringValue `tfsdk:"reject" json:"reject,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o InterfacesWirelessVifSDhcpOptions) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"client_id": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Identifier used by client to identify itself to the DHCP server

`,
		},

		"host_name": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Override system host-name sent to DHCP server

`,
		},

		"mtu": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Use MTU value from DHCP server - ignore interface setting

`,
		},

		"vendor_class_id": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Identify the vendor client type to the DHCP server

`,
		},

		"no_default_route": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Do not install default route to system

`,
		},

		"default_route_distance": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Distance for installed default route

|  Format  |  Description  |
|----------|---------------|
|  u32:1-255  |  Distance for the default route from DHCP server  |

`,

			// Default:          stringdefault.StaticString(`210`),
			Computed: true,
		},

		"reject": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `IP addresses or subnets from which to reject DHCP leases

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address to match  |
|  ipv4net  |  IPv4 prefix to match  |

`,
		},

		// TagNodes

		// Nodes

	}
}
