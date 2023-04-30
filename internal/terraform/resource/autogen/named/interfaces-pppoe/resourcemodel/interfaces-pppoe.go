// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// InterfacesPppoe describes the resource data model.
type InterfacesPppoe struct {
	// LeafNodes
	InterfacesPppoeAccessConcentrator   customtypes.CustomStringValue `tfsdk:"access_concentrator" json:"access-concentrator,omitempty"`
	InterfacesPppoeConnectOnDemand      customtypes.CustomStringValue `tfsdk:"connect_on_demand" json:"connect-on-demand,omitempty"`
	InterfacesPppoeNoDefaultRoute       customtypes.CustomStringValue `tfsdk:"no_default_route" json:"no-default-route,omitempty"`
	InterfacesPppoeDefaultRouteDistance customtypes.CustomStringValue `tfsdk:"default_route_distance" json:"default-route-distance,omitempty"`
	InterfacesPppoeDescrIPtion          customtypes.CustomStringValue `tfsdk:"description" json:"description,omitempty"`
	InterfacesPppoeDisable              customtypes.CustomStringValue `tfsdk:"disable" json:"disable,omitempty"`
	InterfacesPppoeIDleTimeout          customtypes.CustomStringValue `tfsdk:"idle_timeout" json:"idle-timeout,omitempty"`
	InterfacesPppoeHostUniq             customtypes.CustomStringValue `tfsdk:"host_uniq" json:"host-uniq,omitempty"`
	InterfacesPppoeSourceInterface      customtypes.CustomStringValue `tfsdk:"source_interface" json:"source-interface,omitempty"`
	InterfacesPppoeLocalAddress         customtypes.CustomStringValue `tfsdk:"local_address" json:"local-address,omitempty"`
	InterfacesPppoeMtu                  customtypes.CustomStringValue `tfsdk:"mtu" json:"mtu,omitempty"`
	InterfacesPppoeNoPeerDNS            customtypes.CustomStringValue `tfsdk:"no_peer_dns" json:"no-peer-dns,omitempty"`
	InterfacesPppoeRemoteAddress        customtypes.CustomStringValue `tfsdk:"remote_address" json:"remote-address,omitempty"`
	InterfacesPppoeServiceName          customtypes.CustomStringValue `tfsdk:"service_name" json:"service-name,omitempty"`
	InterfacesPppoeRedirect             customtypes.CustomStringValue `tfsdk:"redirect" json:"redirect,omitempty"`
	InterfacesPppoeVrf                  customtypes.CustomStringValue `tfsdk:"vrf" json:"vrf,omitempty"`

	// TagNodes

	// Nodes
	InterfacesPppoeAuthentication  types.Object `tfsdk:"authentication" json:"authentication,omitempty"`
	InterfacesPppoeDhcpvsixOptions types.Object `tfsdk:"dhcpv6_options" json:"dhcpv6-options,omitempty"`
	InterfacesPppoeIP              types.Object `tfsdk:"ip" json:"ip,omitempty"`
	InterfacesPppoeIPvsix          types.Object `tfsdk:"ipv6" json:"ipv6,omitempty"`
	InterfacesPppoeMirror          types.Object `tfsdk:"mirror" json:"mirror,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o InterfacesPppoe) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"access_concentrator": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Access concentrator name

`,
		},

		"connect_on_demand": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Establishment connection automatically when traffic is sent

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

		"description": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |

`,
		},

		"disable": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Administratively disable interface

`,
		},

		"idle_timeout": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Delay before disconnecting idle session (in seconds)

|  Format  |  Description  |
|----------|---------------|
|  u32:0-86400  |  Idle timeout in seconds  |

`,
		},

		"host_uniq": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `PPPoE RFC2516 host-uniq tag

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Host-uniq tag as byte string in HEX  |

`,
		},

		"source_interface": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Interface used to establish connection

|  Format  |  Description  |
|----------|---------------|
|  interface  |  Interface name  |

`,
		},

		"local_address": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `IPv4 address of local end of the PPPoE link

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Address of local end of the PPPoE link  |

`,
		},

		"mtu": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Maximum Transmission Unit (MTU)

|  Format  |  Description  |
|----------|---------------|
|  u32:68-1500  |  Maximum Transmission Unit in byte  |

`,

			// Default:          stringdefault.StaticString(`1492`),
			Computed: true,
		},

		"no_peer_dns": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Do not use DNS servers provided by the peer

`,
		},

		"remote_address": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `IPv4 address of remote end of the PPPoE link

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Address of remote end of the PPPoE link  |

`,
		},

		"service_name": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Service name, only connect to access concentrators advertising this

`,
		},

		"redirect": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Redirect incoming packet to destination

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Destination interface name  |

`,
		},

		"vrf": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `VRF instance name

|  Format  |  Description  |
|----------|---------------|
|  txt  |  VRF instance name  |

`,
		},

		// TagNodes

		// Nodes

		"authentication": schema.SingleNestedAttribute{
			Attributes: InterfacesPppoeAuthentication{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Authentication settings

`,
		},

		"dhcpv6_options": schema.SingleNestedAttribute{
			Attributes: InterfacesPppoeDhcpvsixOptions{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCPv6 client settings/options

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesPppoeIP{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesPppoeIPvsix{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesPppoeMirror{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
		},
	}
}
