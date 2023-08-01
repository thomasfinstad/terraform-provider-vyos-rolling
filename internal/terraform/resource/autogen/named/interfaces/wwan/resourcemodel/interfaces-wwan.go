// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesWwan describes the resource data model.
type InterfacesWwan struct {
	SelfIdentifier types.String `tfsdk:"wwan_id" vyos:",self-id"`

	// LeafNodes
	LeafInterfacesWwanAddress           types.List   `tfsdk:"address" vyos:"address,omitempty"`
	LeafInterfacesWwanApn               types.String `tfsdk:"apn" vyos:"apn,omitempty"`
	LeafInterfacesWwanDescrIPtion       types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafInterfacesWwanDisable           types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafInterfacesWwanDisableLinkDetect types.Bool   `tfsdk:"disable_link_detect" vyos:"disable-link-detect,omitempty"`
	LeafInterfacesWwanMtu               types.Number `tfsdk:"mtu" vyos:"mtu,omitempty"`
	LeafInterfacesWwanConnectOnDemand   types.Bool   `tfsdk:"connect_on_demand" vyos:"connect-on-demand,omitempty"`
	LeafInterfacesWwanRedirect          types.String `tfsdk:"redirect" vyos:"redirect,omitempty"`
	LeafInterfacesWwanVrf               types.String `tfsdk:"vrf" vyos:"vrf,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeInterfacesWwanDhcpOptions     *InterfacesWwanDhcpOptions     `tfsdk:"dhcp_options" vyos:"dhcp-options,omitempty"`
	NodeInterfacesWwanDhcpvsixOptions *InterfacesWwanDhcpvsixOptions `tfsdk:"dhcpv6_options" vyos:"dhcpv6-options,omitempty"`
	NodeInterfacesWwanAuthentication  *InterfacesWwanAuthentication  `tfsdk:"authentication" vyos:"authentication,omitempty"`
	NodeInterfacesWwanMirror          *InterfacesWwanMirror          `tfsdk:"mirror" vyos:"mirror,omitempty"`
	NodeInterfacesWwanIP              *InterfacesWwanIP              `tfsdk:"ip" vyos:"ip,omitempty"`
	NodeInterfacesWwanIPvsix          *InterfacesWwanIPvsix          `tfsdk:"ipv6" vyos:"ipv6,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesWwan) GetVyosPath() []string {
	return []string{
		"interfaces",

		"wwan",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWwan) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"wwan_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Wireless Modem (WWAN) Interface

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  wwanN  &emsp; |  Wireless Wide Area Network interface name  |

`,
		},

		// LeafNodes

		"address": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `IP address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  IPv4 address and prefix length  |
    |  ipv6net  &emsp; |  IPv6 address and prefix length  |
    |  dhcp  &emsp; |  Dynamic Host Configuration Protocol  |
    |  dhcpv6  &emsp; |  Dynamic Host Configuration Protocol for IPv6  |

`,
		},

		"apn": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Access Point Name (APN)

`,
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

		"disable_link_detect": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Ignore link state changes

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"mtu": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum Transmission Unit (MTU)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 68-1500  &emsp; |  Maximum Transmission Unit in byte  |

`,

			// Default:          stringdefault.StaticString(`1430`),
			Computed: true,
		},

		"connect_on_demand": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Establishment connection automatically when traffic is sent

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
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

		"dhcp_options": schema.SingleNestedAttribute{
			Attributes: InterfacesWwanDhcpOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCP client settings/options

`,
		},

		"dhcpv6_options": schema.SingleNestedAttribute{
			Attributes: InterfacesWwanDhcpvsixOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCPv6 client settings/options

`,
		},

		"authentication": schema.SingleNestedAttribute{
			Attributes: InterfacesWwanAuthentication{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Authentication settings

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesWwanMirror{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesWwanIP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesWwanIPvsix{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesWwan) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesWwan) UnmarshalJSON(_ []byte) error {
	return nil
}
