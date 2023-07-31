// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesEthernetVifDhcpvsixOptionsPdInterface describes the resource data model.
type InterfacesEthernetVifDhcpvsixOptionsPdInterface struct {
	ID types.String `tfsdk:"identifier" vyos:",self-id"`

	ParentIDInterfacesEthernet types.String `tfsdk:"ethernet" vyos:"ethernet_identifier,parent-id"`

	ParentIDInterfacesEthernetVif types.String `tfsdk:"vif" vyos:"vif_identifier,parent-id"`

	ParentIDInterfacesEthernetVifDhcpvsixOptionsPd types.String `tfsdk:"pd" vyos:"pd_identifier,parent-id"`

	// LeafNodes
	LeafInterfacesEthernetVifDhcpvsixOptionsPdInterfaceAddress types.String `tfsdk:"address" vyos:"address,omitempty"`
	LeafInterfacesEthernetVifDhcpvsixOptionsPdInterfaceSLAID   types.Number `tfsdk:"sla_id" vyos:"sla-id,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesEthernetVifDhcpvsixOptionsPdInterface) GetVyosPath() []string {
	return []string{
		"interfaces",

		"ethernet",
		o.ParentIDInterfacesEthernet.ValueString(),

		"vif",
		o.ParentIDInterfacesEthernetVif.ValueString(),

		"dhcpv6-options",

		"pd",
		o.ParentIDInterfacesEthernetVifDhcpvsixOptionsPd.ValueString(),

		"interface",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesEthernetVifDhcpvsixOptionsPdInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Delegate IPv6 prefix from provider to this interface

`,
		},

		"ethernet_identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Ethernet Interface

    |  Format  |  Description  |
    |----------|---------------|
    |  ethN  |  Ethernet interface name  |

`,
		},

		"vif_identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Virtual Local Area Network (VLAN) ID

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:0-4094  |  Virtual Local Area Network (VLAN) ID  |

`,
		},

		"pd_identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `DHCPv6 prefix delegation interface statement

    |  Format  |  Description  |
    |----------|---------------|
    |  instance number  |  Prefix delegation instance (>= 0)  |

`,
		},

		// LeafNodes

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Local interface address assigned to interface (default: EUI-64)

    |  Format  |  Description  |
    |----------|---------------|
    |  >0  |  Used to form IPv6 interface address  |

`,
		},

		"sla_id": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interface site-Level aggregator (SLA)

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:0-65535  |  Decimal integer which fits in the length of SLA IDs  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesEthernetVifDhcpvsixOptionsPdInterface) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesEthernetVifDhcpvsixOptionsPdInterface) UnmarshalJSON(_ []byte) error {
	return nil
}
