// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesPseudoEthernetDhcpvsixOptionsPdInterface describes the resource data model.
type InterfacesPseudoEthernetDhcpvsixOptionsPdInterface struct {
	SelfIdentifier types.String `tfsdk:"interface_id" vyos:",self-id"`

	ParentIDInterfacesPseudoEthernet types.String `tfsdk:"pseudo_ethernet" vyos:"pseudo-ethernet,parent-id"`

	ParentIDInterfacesPseudoEthernetDhcpvsixOptionsPd types.String `tfsdk:"pd" vyos:"pd,parent-id"`

	// LeafNodes
	LeafInterfacesPseudoEthernetDhcpvsixOptionsPdInterfaceAddress types.String `tfsdk:"address" vyos:"address,omitempty"`
	LeafInterfacesPseudoEthernetDhcpvsixOptionsPdInterfaceSLAID   types.Number `tfsdk:"sla_id" vyos:"sla-id,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesPseudoEthernetDhcpvsixOptionsPdInterface) GetVyosPath() []string {
	return []string{
		"interfaces",

		"pseudo-ethernet",
		o.ParentIDInterfacesPseudoEthernet.ValueString(),

		"dhcpv6-options",

		"pd",
		o.ParentIDInterfacesPseudoEthernetDhcpvsixOptionsPd.ValueString(),

		"interface",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesPseudoEthernetDhcpvsixOptionsPdInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"interface_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Delegate IPv6 prefix from provider to this interface

`,
		},

		"pseudo_ethernet_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Pseudo Ethernet Interface (Macvlan)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  pethN  &emsp; |  Pseudo Ethernet interface name  |

`,
		},

		"pd_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `DHCPv6 prefix delegation interface statement

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  instance number  &emsp; |  Prefix delegation instance (>= 0)  |

`,
		},

		// LeafNodes

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Local interface address assigned to interface (default: EUI-64)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  >0  &emsp; |  Used to form IPv6 interface address  |

`,
		},

		"sla_id": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interface site-Level aggregator (SLA)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-65535  &emsp; |  Decimal integer which fits in the length of SLA IDs  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesPseudoEthernetDhcpvsixOptionsPdInterface) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesPseudoEthernetDhcpvsixOptionsPdInterface) UnmarshalJSON(_ []byte) error {
	return nil
}
