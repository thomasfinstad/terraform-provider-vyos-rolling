// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesBonding describes the resource data model.
type InterfacesBonding struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"bonding_id" vyos:",self-id"`

	// LeafNodes
	LeafInterfacesBondingAddress           types.List   `tfsdk:"address" vyos:"address,omitempty"`
	LeafInterfacesBondingDescrIPtion       types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafInterfacesBondingDisableLinkDetect types.Bool   `tfsdk:"disable_link_detect" vyos:"disable-link-detect,omitempty"`
	LeafInterfacesBondingDisable           types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafInterfacesBondingVrf               types.String `tfsdk:"vrf" vyos:"vrf,omitempty"`
	LeafInterfacesBondingHashPolicy        types.String `tfsdk:"hash_policy" vyos:"hash-policy,omitempty"`
	LeafInterfacesBondingMac               types.String `tfsdk:"mac" vyos:"mac,omitempty"`
	LeafInterfacesBondingMiiMonInterval    types.Number `tfsdk:"mii_mon_interval" vyos:"mii-mon-interval,omitempty"`
	LeafInterfacesBondingMinLinks          types.Number `tfsdk:"min_links" vyos:"min-links,omitempty"`
	LeafInterfacesBondingLacpRate          types.String `tfsdk:"lacp_rate" vyos:"lacp-rate,omitempty"`
	LeafInterfacesBondingMode              types.String `tfsdk:"mode" vyos:"mode,omitempty"`
	LeafInterfacesBondingMtu               types.Number `tfsdk:"mtu" vyos:"mtu,omitempty"`
	LeafInterfacesBondingPrimary           types.String `tfsdk:"primary" vyos:"primary,omitempty"`
	LeafInterfacesBondingRedirect          types.String `tfsdk:"redirect" vyos:"redirect,omitempty"`
	LeafInterfacesBondingXdp               types.Bool   `tfsdk:"xdp" vyos:"xdp,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagInterfacesBondingVifS bool `tfsdk:"vif_s" vyos:"vif-s,child"`
	ExistsTagInterfacesBondingVif  bool `tfsdk:"vif" vyos:"vif,child"`

	// Nodes
	NodeInterfacesBondingArpMonitor      *InterfacesBondingArpMonitor      `tfsdk:"arp_monitor" vyos:"arp-monitor,omitempty"`
	NodeInterfacesBondingDhcpOptions     *InterfacesBondingDhcpOptions     `tfsdk:"dhcp_options" vyos:"dhcp-options,omitempty"`
	NodeInterfacesBondingDhcpvsixOptions *InterfacesBondingDhcpvsixOptions `tfsdk:"dhcpv6_options" vyos:"dhcpv6-options,omitempty"`
	NodeInterfacesBondingMirror          *InterfacesBondingMirror          `tfsdk:"mirror" vyos:"mirror,omitempty"`
	NodeInterfacesBondingIP              *InterfacesBondingIP              `tfsdk:"ip" vyos:"ip,omitempty"`
	NodeInterfacesBondingIPvsix          *InterfacesBondingIPvsix          `tfsdk:"ipv6" vyos:"ipv6,omitempty"`
	NodeInterfacesBondingMember          *InterfacesBondingMember          `tfsdk:"member" vyos:"member,omitempty"`
}

// GetID returns the resource ID
func (o InterfacesBonding) GetID() *types.String {
	return &o.ID
}

// SetID configures the resource ID
func (o InterfacesBonding) SetID(id types.String) {
	o.ID = id
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesBonding) GetVyosPath() []string {
	return []string{
		"interfaces",

		"bonding",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesBonding) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"bonding_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Bonding Interface/Link Aggregation

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  bondN  &emsp; |  Bonding interface name  |

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

		"vrf": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VRF instance name

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  VRF instance name  |

`,
		},

		"hash_policy": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Bonding transmit hash policy

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  layer2  &emsp; |  use MAC addresses to generate the hash  |
    |  layer2+3  &emsp; |  combine MAC address and IP address to make hash  |
    |  layer3+4  &emsp; |  combine IP address and port to make hash  |
    |  encap2+3  &emsp; |  combine encapsulated MAC address and IP address to make hash  |
    |  encap3+4  &emsp; |  combine encapsulated IP address and port to make hash  |

`,

			// Default:          stringdefault.StaticString(`layer2`),
			Computed: true,
		},

		"mac": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Media Access Control (MAC) address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  macaddr  &emsp; |  Hardware (MAC) address  |

`,
		},

		"mii_mon_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Specifies the MII link monitoring frequency in milliseconds

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0  &emsp; |  Disable MII link monitoring  |
    |  number: 50-1000  &emsp; |  MII link monitoring frequency in milliseconds  |

`,

			// Default:          stringdefault.StaticString(`100`),
			Computed: true,
		},

		"min_links": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Minimum number of member interfaces required up before enabling bond

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-16  &emsp; |  Minimum number of member interfaces required up before enabling bond  |

`,

			// Default:          stringdefault.StaticString(`0`),
			Computed: true,
		},

		"lacp_rate": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Rate in which we will ask our link partner to transmit LACPDU packets

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  slow  &emsp; |  Request partner to transmit LACPDUs every 30 seconds  |
    |  fast  &emsp; |  Request partner to transmit LACPDUs every 1 second  |

`,

			// Default:          stringdefault.StaticString(`slow`),
			Computed: true,
		},

		"mode": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Bonding mode

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  802.3ad  &emsp; |  IEEE 802.3ad Dynamic link aggregation  |
    |  active-backup  &emsp; |  Fault tolerant: only one slave in the bond is active  |
    |  broadcast  &emsp; |  Fault tolerant: transmits everything on all slave interfaces  |
    |  round-robin  &emsp; |  Load balance: transmit packets in sequential order  |
    |  transmit-load-balance  &emsp; |  Load balance: adapts based on transmit load and speed  |
    |  adaptive-load-balance  &emsp; |  Load balance: adapts based on transmit and receive plus ARP  |
    |  xor-hash  &emsp; |  Distribute based on MAC address  |

`,

			// Default:          stringdefault.StaticString(`802.3ad`),
			Computed: true,
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

		"primary": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Primary device interface

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Interface name  |

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

		"xdp": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable eXpress Data Path

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

		"arp_monitor": schema.SingleNestedAttribute{
			Attributes: InterfacesBondingArpMonitor{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `ARP link monitoring parameters

`,
		},

		"dhcp_options": schema.SingleNestedAttribute{
			Attributes: InterfacesBondingDhcpOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCP client settings/options

`,
		},

		"dhcpv6_options": schema.SingleNestedAttribute{
			Attributes: InterfacesBondingDhcpvsixOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCPv6 client settings/options

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesBondingMirror{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesBondingIP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesBondingIPvsix{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
		},

		"member": schema.SingleNestedAttribute{
			Attributes: InterfacesBondingMember{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Bridge member interfaces

`,
		},
	}
}
