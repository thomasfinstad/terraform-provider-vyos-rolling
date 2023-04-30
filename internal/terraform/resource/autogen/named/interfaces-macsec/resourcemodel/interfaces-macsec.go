// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// InterfacesMacsec describes the resource data model.
type InterfacesMacsec struct {
	// LeafNodes
	InterfacesMacsecAddress         customtypes.CustomStringValue `tfsdk:"address" json:"address,omitempty"`
	InterfacesMacsecDescrIPtion     customtypes.CustomStringValue `tfsdk:"description" json:"description,omitempty"`
	InterfacesMacsecDisable         customtypes.CustomStringValue `tfsdk:"disable" json:"disable,omitempty"`
	InterfacesMacsecMtu             customtypes.CustomStringValue `tfsdk:"mtu" json:"mtu,omitempty"`
	InterfacesMacsecSourceInterface customtypes.CustomStringValue `tfsdk:"source_interface" json:"source-interface,omitempty"`
	InterfacesMacsecRedirect        customtypes.CustomStringValue `tfsdk:"redirect" json:"redirect,omitempty"`
	InterfacesMacsecVrf             customtypes.CustomStringValue `tfsdk:"vrf" json:"vrf,omitempty"`

	// TagNodes

	// Nodes
	InterfacesMacsecDhcpOptions     types.Object `tfsdk:"dhcp_options" json:"dhcp-options,omitempty"`
	InterfacesMacsecDhcpvsixOptions types.Object `tfsdk:"dhcpv6_options" json:"dhcpv6-options,omitempty"`
	InterfacesMacsecIP              types.Object `tfsdk:"ip" json:"ip,omitempty"`
	InterfacesMacsecIPvsix          types.Object `tfsdk:"ipv6" json:"ipv6,omitempty"`
	InterfacesMacsecMirror          types.Object `tfsdk:"mirror" json:"mirror,omitempty"`
	InterfacesMacsecSecURIty        types.Object `tfsdk:"security" json:"security,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o InterfacesMacsec) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `IP address

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  IPv4 address and prefix length  |
|  ipv6net  |  IPv6 address and prefix length  |
|  dhcp  |  Dynamic Host Configuration Protocol  |
|  dhcpv6  |  Dynamic Host Configuration Protocol for IPv6  |

`,
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

		"mtu": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Maximum Transmission Unit (MTU)

|  Format  |  Description  |
|----------|---------------|
|  u32:68-16000  |  Maximum Transmission Unit in byte  |

`,

			// Default:          stringdefault.StaticString(`1460`),
			Computed: true,
		},

		"source_interface": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Physical interface the traffic will go through

|  Format  |  Description  |
|----------|---------------|
|  interface  |  Physical interface used for traffic forwarding  |

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

		"dhcp_options": schema.SingleNestedAttribute{
			Attributes: InterfacesMacsecDhcpOptions{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCP client settings/options

`,
		},

		"dhcpv6_options": schema.SingleNestedAttribute{
			Attributes: InterfacesMacsecDhcpvsixOptions{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCPv6 client settings/options

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesMacsecIP{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesMacsecIPvsix{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesMacsecMirror{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
		},

		"security": schema.SingleNestedAttribute{
			Attributes: InterfacesMacsecSecURIty{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Security/Encryption Settings

`,
		},
	}
}
