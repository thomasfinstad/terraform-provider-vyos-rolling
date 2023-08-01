// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesBondingVifSVifC describes the resource data model.
type InterfacesBondingVifSVifC struct {
	SelfIdentifier types.String `tfsdk:"vif_c_id" vyos:",self-id"`

	ParentIDInterfacesBonding types.String `tfsdk:"bonding" vyos:"bonding,parent-id"`

	ParentIDInterfacesBondingVifS types.String `tfsdk:"vif_s" vyos:"vif-s,parent-id"`

	// LeafNodes
	LeafInterfacesBondingVifSVifCDescrIPtion       types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafInterfacesBondingVifSVifCAddress           types.List   `tfsdk:"address" vyos:"address,omitempty"`
	LeafInterfacesBondingVifSVifCDisableLinkDetect types.Bool   `tfsdk:"disable_link_detect" vyos:"disable-link-detect,omitempty"`
	LeafInterfacesBondingVifSVifCDisable           types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafInterfacesBondingVifSVifCMac               types.String `tfsdk:"mac" vyos:"mac,omitempty"`
	LeafInterfacesBondingVifSVifCMtu               types.Number `tfsdk:"mtu" vyos:"mtu,omitempty"`
	LeafInterfacesBondingVifSVifCRedirect          types.String `tfsdk:"redirect" vyos:"redirect,omitempty"`
	LeafInterfacesBondingVifSVifCVrf               types.String `tfsdk:"vrf" vyos:"vrf,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeInterfacesBondingVifSVifCDhcpOptions     *InterfacesBondingVifSVifCDhcpOptions     `tfsdk:"dhcp_options" vyos:"dhcp-options,omitempty"`
	NodeInterfacesBondingVifSVifCDhcpvsixOptions *InterfacesBondingVifSVifCDhcpvsixOptions `tfsdk:"dhcpv6_options" vyos:"dhcpv6-options,omitempty"`
	NodeInterfacesBondingVifSVifCIP              *InterfacesBondingVifSVifCIP              `tfsdk:"ip" vyos:"ip,omitempty"`
	NodeInterfacesBondingVifSVifCIPvsix          *InterfacesBondingVifSVifCIPvsix          `tfsdk:"ipv6" vyos:"ipv6,omitempty"`
	NodeInterfacesBondingVifSVifCMirror          *InterfacesBondingVifSVifCMirror          `tfsdk:"mirror" vyos:"mirror,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesBondingVifSVifC) GetVyosPath() []string {
	return []string{
		"interfaces",

		"bonding",
		o.ParentIDInterfacesBonding.ValueString(),

		"vif-s",
		o.ParentIDInterfacesBondingVifS.ValueString(),

		"vif-c",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesBondingVifSVifC) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"vif_c_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `QinQ TAG-C Virtual Local Area Network (VLAN) ID

`,
		},

		"bonding_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Bonding Interface/Link Aggregation

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  bondN  &emsp; |  Bonding interface name  |

`,
		},

		"vif_s_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `QinQ TAG-S Virtual Local Area Network (VLAN) ID

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-4094  &emsp; |  QinQ Virtual Local Area Network (VLAN) ID  |

`,
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
			Attributes: InterfacesBondingVifSVifCDhcpOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCP client settings/options

`,
		},

		"dhcpv6_options": schema.SingleNestedAttribute{
			Attributes: InterfacesBondingVifSVifCDhcpvsixOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCPv6 client settings/options

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesBondingVifSVifCIP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesBondingVifSVifCIPvsix{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesBondingVifSVifCMirror{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesBondingVifSVifC) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesBondingVifSVifC) UnmarshalJSON(_ []byte) error {
	return nil
}
