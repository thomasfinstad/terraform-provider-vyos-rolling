// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// InterfacesEthernetVif describes the resource data model.
type InterfacesEthernetVif struct {
	// LeafNodes
	InterfacesEthernetVifDescrIPtion       customtypes.CustomStringValue `tfsdk:"description" json:"description,omitempty"`
	InterfacesEthernetVifAddress           customtypes.CustomStringValue `tfsdk:"address" json:"address,omitempty"`
	InterfacesEthernetVifDisableLinkDetect customtypes.CustomStringValue `tfsdk:"disable_link_detect" json:"disable-link-detect,omitempty"`
	InterfacesEthernetVifDisable           customtypes.CustomStringValue `tfsdk:"disable" json:"disable,omitempty"`
	InterfacesEthernetVifEgressQos         customtypes.CustomStringValue `tfsdk:"egress_qos" json:"egress-qos,omitempty"`
	InterfacesEthernetVifIngressQos        customtypes.CustomStringValue `tfsdk:"ingress_qos" json:"ingress-qos,omitempty"`
	InterfacesEthernetVifMac               customtypes.CustomStringValue `tfsdk:"mac" json:"mac,omitempty"`
	InterfacesEthernetVifMtu               customtypes.CustomStringValue `tfsdk:"mtu" json:"mtu,omitempty"`
	InterfacesEthernetVifRedirect          customtypes.CustomStringValue `tfsdk:"redirect" json:"redirect,omitempty"`
	InterfacesEthernetVifVrf               customtypes.CustomStringValue `tfsdk:"vrf" json:"vrf,omitempty"`

	// TagNodes

	// Nodes
	InterfacesEthernetVifDhcpOptions     types.Object `tfsdk:"dhcp_options" json:"dhcp-options,omitempty"`
	InterfacesEthernetVifDhcpvsixOptions types.Object `tfsdk:"dhcpv6_options" json:"dhcpv6-options,omitempty"`
	InterfacesEthernetVifIP              types.Object `tfsdk:"ip" json:"ip,omitempty"`
	InterfacesEthernetVifIPvsix          types.Object `tfsdk:"ipv6" json:"ipv6,omitempty"`
	InterfacesEthernetVifMirror          types.Object `tfsdk:"mirror" json:"mirror,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o InterfacesEthernetVif) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"description": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |

`,
		},

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

		"disable_link_detect": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Ignore link state changes

`,
		},

		"disable": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Administratively disable interface

`,
		},

		"egress_qos": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `VLAN egress QoS

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Format for qos mapping, e.g.: '0:1 1:6 7:6'  |

`,
		},

		"ingress_qos": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `VLAN ingress QoS

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Format for qos mapping, e.g.: '0:1 1:6 7:6'  |

`,
		},

		"mac": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Media Access Control (MAC) address

|  Format  |  Description  |
|----------|---------------|
|  macaddr  |  Hardware (MAC) address  |

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

			// Default:          stringdefault.StaticString(`1500`),
			Computed: true,
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
			Attributes: InterfacesEthernetVifDhcpOptions{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCP client settings/options

`,
		},

		"dhcpv6_options": schema.SingleNestedAttribute{
			Attributes: InterfacesEthernetVifDhcpvsixOptions{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCPv6 client settings/options

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesEthernetVifIP{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesEthernetVifIPvsix{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesEthernetVifMirror{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
		},
	}
}
