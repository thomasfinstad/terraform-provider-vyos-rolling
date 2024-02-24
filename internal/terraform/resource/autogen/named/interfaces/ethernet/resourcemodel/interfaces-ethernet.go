// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesEthernet describes the resource data model.
type InterfacesEthernet struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"ethernet_id" vyos:",self-id"`

	// LeafNodes
	LeafInterfacesEthernetAddress            types.List   `tfsdk:"address" vyos:"address,omitempty"`
	LeafInterfacesEthernetDescrIPtion        types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafInterfacesEthernetDisableFlowControl types.Bool   `tfsdk:"disable_flow_control" vyos:"disable-flow-control,omitempty"`
	LeafInterfacesEthernetDisableLinkDetect  types.Bool   `tfsdk:"disable_link_detect" vyos:"disable-link-detect,omitempty"`
	LeafInterfacesEthernetDisable            types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafInterfacesEthernetDuplex             types.String `tfsdk:"duplex" vyos:"duplex,omitempty"`
	LeafInterfacesEthernetHwID               types.String `tfsdk:"hw_id" vyos:"hw-id,omitempty"`
	LeafInterfacesEthernetMac                types.String `tfsdk:"mac" vyos:"mac,omitempty"`
	LeafInterfacesEthernetMtu                types.Number `tfsdk:"mtu" vyos:"mtu,omitempty"`
	LeafInterfacesEthernetSpeed              types.String `tfsdk:"speed" vyos:"speed,omitempty"`
	LeafInterfacesEthernetRedirect           types.String `tfsdk:"redirect" vyos:"redirect,omitempty"`
	LeafInterfacesEthernetVrf                types.String `tfsdk:"vrf" vyos:"vrf,omitempty"`
	LeafInterfacesEthernetXdp                types.Bool   `tfsdk:"xdp" vyos:"xdp,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagInterfacesEthernetVifS bool `tfsdk:"-" vyos:"vif-s,ignore,child"`
	ExistsTagInterfacesEthernetVif  bool `tfsdk:"-" vyos:"vif,ignore,child"`

	// Nodes
	NodeInterfacesEthernetDhcpOptions     *InterfacesEthernetDhcpOptions     `tfsdk:"dhcp_options" vyos:"dhcp-options,omitempty"`
	NodeInterfacesEthernetDhcpvsixOptions *InterfacesEthernetDhcpvsixOptions `tfsdk:"dhcpv6_options" vyos:"dhcpv6-options,omitempty"`
	NodeInterfacesEthernetEapol           *InterfacesEthernetEapol           `tfsdk:"eapol" vyos:"eapol,omitempty"`
	NodeInterfacesEthernetIP              *InterfacesEthernetIP              `tfsdk:"ip" vyos:"ip,omitempty"`
	NodeInterfacesEthernetIPvsix          *InterfacesEthernetIPvsix          `tfsdk:"ipv6" vyos:"ipv6,omitempty"`
	NodeInterfacesEthernetMirror          *InterfacesEthernetMirror          `tfsdk:"mirror" vyos:"mirror,omitempty"`
	NodeInterfacesEthernetOffload         *InterfacesEthernetOffload         `tfsdk:"offload" vyos:"offload,omitempty"`
	NodeInterfacesEthernetRingBuffer      *InterfacesEthernetRingBuffer      `tfsdk:"ring_buffer" vyos:"ring-buffer,omitempty"`
}

// SetID configures the resource ID
func (o *InterfacesEthernet) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesEthernet) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"interfaces",

		"ethernet",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesEthernet) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"ethernet_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Ethernet Interface

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ethN  &emsp; |  Ethernet interface name  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
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

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Description  |

`,
		},

		"disable_flow_control": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable Ethernet flow control (pause frames)

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

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Administratively disable interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"duplex": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Duplex mode

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  auto  &emsp; |  Auto negotiation  |
    |  half  &emsp; |  Half duplex  |
    |  full  &emsp; |  Full duplex  |

`,

			// Default:          stringdefault.StaticString(`auto`),
			Computed: true,
		},

		"hw_id": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Associate Ethernet Interface with given Media Access Control (MAC) address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  macaddr  &emsp; |  Hardware (MAC) address  |

`,
		},

		"mac": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Media Access Control (MAC) address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  macaddr  &emsp; |  Hardware (MAC) address  |

`,
		},

		"mtu": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum Transmission Unit (MTU)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 68-16000  &emsp; |  Maximum Transmission Unit in byte  |

`,

			// Default:          stringdefault.StaticString(`1500`),
			Computed: true,
		},

		"speed": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Link speed

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  auto  &emsp; |  Auto negotiation  |
    |  10  &emsp; |  10 Mbit/sec  |
    |  100  &emsp; |  100 Mbit/sec  |
    |  1000  &emsp; |  1 Gbit/sec  |
    |  2500  &emsp; |  2.5 Gbit/sec  |
    |  5000  &emsp; |  5 Gbit/sec  |
    |  10000  &emsp; |  10 Gbit/sec  |
    |  25000  &emsp; |  25 Gbit/sec  |
    |  40000  &emsp; |  40 Gbit/sec  |
    |  50000  &emsp; |  50 Gbit/sec  |
    |  100000  &emsp; |  100 Gbit/sec  |

`,

			// Default:          stringdefault.StaticString(`auto`),
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

		"xdp": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable eXpress Data Path

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

		"dhcp_options": schema.SingleNestedAttribute{
			Attributes: InterfacesEthernetDhcpOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCP client settings/options

`,
		},

		"dhcpv6_options": schema.SingleNestedAttribute{
			Attributes: InterfacesEthernetDhcpvsixOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCPv6 client settings/options

`,
		},

		"eapol": schema.SingleNestedAttribute{
			Attributes: InterfacesEthernetEapol{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Extensible Authentication Protocol over Local Area Network

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesEthernetIP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesEthernetIPvsix{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesEthernetMirror{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
		},

		"offload": schema.SingleNestedAttribute{
			Attributes: InterfacesEthernetOffload{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Configurable offload options

`,
		},

		"ring_buffer": schema.SingleNestedAttribute{
			Attributes: InterfacesEthernetRingBuffer{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Shared buffer between the device driver and NIC

`,
		},
	}
}
