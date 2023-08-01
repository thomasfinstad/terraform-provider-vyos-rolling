// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesPppoe describes the resource data model.
type InterfacesPppoe struct {
	SelfIdentifier types.String `tfsdk:"pppoe_id" vyos:",self-id"`

	// LeafNodes
	LeafInterfacesPppoeAccessConcentrator   types.String `tfsdk:"access_concentrator" vyos:"access-concentrator,omitempty"`
	LeafInterfacesPppoeConnectOnDemand      types.Bool   `tfsdk:"connect_on_demand" vyos:"connect-on-demand,omitempty"`
	LeafInterfacesPppoeNoDefaultRoute       types.Bool   `tfsdk:"no_default_route" vyos:"no-default-route,omitempty"`
	LeafInterfacesPppoeDefaultRouteDistance types.Number `tfsdk:"default_route_distance" vyos:"default-route-distance,omitempty"`
	LeafInterfacesPppoeDescrIPtion          types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafInterfacesPppoeDisable              types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafInterfacesPppoeIDleTimeout          types.Number `tfsdk:"idle_timeout" vyos:"idle-timeout,omitempty"`
	LeafInterfacesPppoeHostUniq             types.String `tfsdk:"host_uniq" vyos:"host-uniq,omitempty"`
	LeafInterfacesPppoeSourceInterface      types.String `tfsdk:"source_interface" vyos:"source-interface,omitempty"`
	LeafInterfacesPppoeLocalAddress         types.String `tfsdk:"local_address" vyos:"local-address,omitempty"`
	LeafInterfacesPppoeMtu                  types.Number `tfsdk:"mtu" vyos:"mtu,omitempty"`
	LeafInterfacesPppoeNoPeerDNS            types.Bool   `tfsdk:"no_peer_dns" vyos:"no-peer-dns,omitempty"`
	LeafInterfacesPppoeRemoteAddress        types.String `tfsdk:"remote_address" vyos:"remote-address,omitempty"`
	LeafInterfacesPppoeServiceName          types.String `tfsdk:"service_name" vyos:"service-name,omitempty"`
	LeafInterfacesPppoeRedirect             types.String `tfsdk:"redirect" vyos:"redirect,omitempty"`
	LeafInterfacesPppoeVrf                  types.String `tfsdk:"vrf" vyos:"vrf,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeInterfacesPppoeAuthentication  *InterfacesPppoeAuthentication  `tfsdk:"authentication" vyos:"authentication,omitempty"`
	NodeInterfacesPppoeDhcpvsixOptions *InterfacesPppoeDhcpvsixOptions `tfsdk:"dhcpv6_options" vyos:"dhcpv6-options,omitempty"`
	NodeInterfacesPppoeIP              *InterfacesPppoeIP              `tfsdk:"ip" vyos:"ip,omitempty"`
	NodeInterfacesPppoeIPvsix          *InterfacesPppoeIPvsix          `tfsdk:"ipv6" vyos:"ipv6,omitempty"`
	NodeInterfacesPppoeMirror          *InterfacesPppoeMirror          `tfsdk:"mirror" vyos:"mirror,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesPppoe) GetVyosPath() []string {
	return []string{
		"interfaces",

		"pppoe",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesPppoe) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"pppoe_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Point-to-Point Protocol over Ethernet (PPPoE) Interface

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  pppoeN  &emsp; |  PPPoE dialer interface name  |

`,
		},

		// LeafNodes

		"access_concentrator": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Access concentrator name

`,
		},

		"connect_on_demand": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Establishment connection automatically when traffic is sent

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
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

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-255  &emsp; |  Distance for the default route from DHCP server  |

`,

			// Default:          stringdefault.StaticString(`210`),
			Computed: true,
		},

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Description  |

`,
		},

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Administratively disable interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"idle_timeout": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Delay before disconnecting idle session (in seconds)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-86400  &emsp; |  Idle timeout in seconds  |

`,
		},

		"host_uniq": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `PPPoE RFC2516 host-uniq tag

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Host-uniq tag as byte string in HEX  |

`,
		},

		"source_interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interface used to establish connection

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  interface  &emsp; |  Interface name  |

`,
		},

		"local_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv4 address of local end of the PPPoE link

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  Address of local end of the PPPoE link  |

`,
		},

		"mtu": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum Transmission Unit (MTU)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 68-1500  &emsp; |  Maximum Transmission Unit in byte  |

`,

			// Default:          stringdefault.StaticString(`1492`),
			Computed: true,
		},

		"no_peer_dns": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not use DNS servers provided by the peer

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"remote_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv4 address of remote end of the PPPoE link

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  Address of remote end of the PPPoE link  |

`,
		},

		"service_name": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Service name, only connect to access concentrators advertising this

`,
		},

		"redirect": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Redirect incoming packet to destination

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Destination interface name  |

`,
		},

		"vrf": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VRF instance name

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  VRF instance name  |

`,
		},

		// Nodes

		"authentication": schema.SingleNestedAttribute{
			Attributes: InterfacesPppoeAuthentication{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Authentication settings

`,
		},

		"dhcpv6_options": schema.SingleNestedAttribute{
			Attributes: InterfacesPppoeDhcpvsixOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCPv6 client settings/options

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesPppoeIP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesPppoeIPvsix{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesPppoeMirror{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesPppoe) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesPppoe) UnmarshalJSON(_ []byte) error {
	return nil
}
