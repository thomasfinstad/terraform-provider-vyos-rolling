// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesVxlan describes the resource data model.
type InterfacesVxlan struct {
	SelfIdentifier types.String `tfsdk:"vxlan_id" vyos:",self-id"`

	// LeafNodes
	LeafInterfacesVxlanAddress         types.List   `tfsdk:"address" vyos:"address,omitempty"`
	LeafInterfacesVxlanDescrIPtion     types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafInterfacesVxlanDisable         types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafInterfacesVxlanExternal        types.Bool   `tfsdk:"external" vyos:"external,omitempty"`
	LeafInterfacesVxlanGpe             types.Bool   `tfsdk:"gpe" vyos:"gpe,omitempty"`
	LeafInterfacesVxlanGroup           types.String `tfsdk:"group" vyos:"group,omitempty"`
	LeafInterfacesVxlanMac             types.String `tfsdk:"mac" vyos:"mac,omitempty"`
	LeafInterfacesVxlanMtu             types.Number `tfsdk:"mtu" vyos:"mtu,omitempty"`
	LeafInterfacesVxlanPort            types.Number `tfsdk:"port" vyos:"port,omitempty"`
	LeafInterfacesVxlanSourceAddress   types.String `tfsdk:"source_address" vyos:"source-address,omitempty"`
	LeafInterfacesVxlanSourceInterface types.String `tfsdk:"source_interface" vyos:"source-interface,omitempty"`
	LeafInterfacesVxlanRemote          types.List   `tfsdk:"remote" vyos:"remote,omitempty"`
	LeafInterfacesVxlanRedirect        types.String `tfsdk:"redirect" vyos:"redirect,omitempty"`
	LeafInterfacesVxlanVrf             types.String `tfsdk:"vrf" vyos:"vrf,omitempty"`
	LeafInterfacesVxlanVni             types.Number `tfsdk:"vni" vyos:"vni,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeInterfacesVxlanIP         *InterfacesVxlanIP         `tfsdk:"ip" vyos:"ip,omitempty"`
	NodeInterfacesVxlanIPvsix     *InterfacesVxlanIPvsix     `tfsdk:"ipv6" vyos:"ipv6,omitempty"`
	NodeInterfacesVxlanMirror     *InterfacesVxlanMirror     `tfsdk:"mirror" vyos:"mirror,omitempty"`
	NodeInterfacesVxlanParameters *InterfacesVxlanParameters `tfsdk:"parameters" vyos:"parameters,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesVxlan) GetVyosPath() []string {
	return []string{
		"interfaces",

		"vxlan",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesVxlan) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, an amalgamation of the `vxlan_id` and the parents `*_id` fields seperated by dunder `__` starting with top level ancestor.",
		},
		"vxlan_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Virtual Extensible LAN (VXLAN) Interface

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  vxlanN  &emsp; |  VXLAN interface name  |

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

		"external": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Use external control plane

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"gpe": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable Generic Protocol extension (VXLAN-GPE)

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Multicast group address for VXLAN interface

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  Multicast IPv4 group address  |
    |  ipv6  &emsp; |  Multicast IPv6 group address  |

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
    |  number: 1200-16000  &emsp; |  Maximum Transmission Unit in byte  |

`,

			// Default:          stringdefault.StaticString(`1450`),
			Computed: true,
		},

		"port": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Port number used by connection

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Numeric IP port  |

`,

			// Default:          stringdefault.StaticString(`8472`),
			Computed: true,
		},

		"source_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Source IP address used to initiate connection

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IPv4 source address  |
    |  ipv6  &emsp; |  IPv6 source address  |

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

		"remote": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Tunnel remote address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  Tunnel remote IPv4 address  |
    |  ipv6  &emsp; |  Tunnel remote IPv6 address  |

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

		"vni": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Virtual Network Identifier

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-16777214  &emsp; |  VXLAN virtual network identifier  |

`,
		},

		// Nodes

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesVxlanIP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesVxlanIPvsix{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesVxlanMirror{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
		},

		"parameters": schema.SingleNestedAttribute{
			Attributes: InterfacesVxlanParameters{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `VXLAN tunnel parameters

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesVxlan) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesVxlan) UnmarshalJSON(_ []byte) error {
	return nil
}
