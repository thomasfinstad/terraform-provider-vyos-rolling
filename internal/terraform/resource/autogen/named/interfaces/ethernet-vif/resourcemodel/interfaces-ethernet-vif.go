// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesEthernetVif describes the resource data model.
type InterfacesEthernetVif struct {
	SelfIdentifier types.Number `tfsdk:"vif_id" vyos:",self-id"`

	ParentIDInterfacesEthernet types.String `tfsdk:"ethernet" vyos:"ethernet,parent-id"`

	// LeafNodes
	LeafInterfacesEthernetVifDescrIPtion       types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafInterfacesEthernetVifAddress           types.List   `tfsdk:"address" vyos:"address,omitempty"`
	LeafInterfacesEthernetVifDisableLinkDetect types.Bool   `tfsdk:"disable_link_detect" vyos:"disable-link-detect,omitempty"`
	LeafInterfacesEthernetVifDisable           types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafInterfacesEthernetVifEgressQos         types.String `tfsdk:"egress_qos" vyos:"egress-qos,omitempty"`
	LeafInterfacesEthernetVifIngressQos        types.String `tfsdk:"ingress_qos" vyos:"ingress-qos,omitempty"`
	LeafInterfacesEthernetVifMac               types.String `tfsdk:"mac" vyos:"mac,omitempty"`
	LeafInterfacesEthernetVifMtu               types.Number `tfsdk:"mtu" vyos:"mtu,omitempty"`
	LeafInterfacesEthernetVifRedirect          types.String `tfsdk:"redirect" vyos:"redirect,omitempty"`
	LeafInterfacesEthernetVifVrf               types.String `tfsdk:"vrf" vyos:"vrf,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeInterfacesEthernetVifDhcpOptions     *InterfacesEthernetVifDhcpOptions     `tfsdk:"dhcp_options" vyos:"dhcp-options,omitempty"`
	NodeInterfacesEthernetVifDhcpvsixOptions *InterfacesEthernetVifDhcpvsixOptions `tfsdk:"dhcpv6_options" vyos:"dhcpv6-options,omitempty"`
	NodeInterfacesEthernetVifIP              *InterfacesEthernetVifIP              `tfsdk:"ip" vyos:"ip,omitempty"`
	NodeInterfacesEthernetVifIPvsix          *InterfacesEthernetVifIPvsix          `tfsdk:"ipv6" vyos:"ipv6,omitempty"`
	NodeInterfacesEthernetVifMirror          *InterfacesEthernetVifMirror          `tfsdk:"mirror" vyos:"mirror,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesEthernetVif) GetVyosPath() []string {
	return []string{
		"interfaces",

		"ethernet",
		o.ParentIDInterfacesEthernet.ValueString(),

		"vif",
		o.SelfIdentifier.ValueBigFloat().String(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesEthernetVif) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, an amalgamation of the `vif_id` and the parents `*_id` fields seperated by dunder `__` starting with top level ancestor.",
		},
		"vif_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Virtual Local Area Network (VLAN) ID

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-4094  &emsp; |  Virtual Local Area Network (VLAN) ID  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
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

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Description  |

`,
		},

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

		"egress_qos": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VLAN egress QoS

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Format for qos mapping, e.g.: '0:1 1:6 7:6'  |

`,
		},

		"ingress_qos": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VLAN ingress QoS

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Format for qos mapping, e.g.: '0:1 1:6 7:6'  |

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
			Attributes: InterfacesEthernetVifDhcpOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCP client settings/options

`,
		},

		"dhcpv6_options": schema.SingleNestedAttribute{
			Attributes: InterfacesEthernetVifDhcpvsixOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCPv6 client settings/options

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesEthernetVifIP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesEthernetVifIPvsix{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesEthernetVifMirror{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesEthernetVif) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesEthernetVif) UnmarshalJSON(_ []byte) error {
	return nil
}
