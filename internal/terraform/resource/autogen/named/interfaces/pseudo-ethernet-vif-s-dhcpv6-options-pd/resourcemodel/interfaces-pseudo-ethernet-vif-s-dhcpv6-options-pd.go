// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesPseudoEthernetVifSDhcpvsixOptionsPd describes the resource data model.
type InterfacesPseudoEthernetVifSDhcpvsixOptionsPd struct {
	SelfIdentifier types.String `tfsdk:"pd_id" vyos:",self-id"`

	ParentIDInterfacesPseudoEthernet types.String `tfsdk:"pseudo_ethernet" vyos:"pseudo-ethernet,parent-id"`

	ParentIDInterfacesPseudoEthernetVifS types.String `tfsdk:"vif_s" vyos:"vif-s,parent-id"`

	// LeafNodes
	LeafInterfacesPseudoEthernetVifSDhcpvsixOptionsPdLength types.Number `tfsdk:"length" vyos:"length,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagInterfacesPseudoEthernetVifSDhcpvsixOptionsPdInterface bool `tfsdk:"interface" vyos:"interface,child"`

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesPseudoEthernetVifSDhcpvsixOptionsPd) GetVyosPath() []string {
	return []string{
		"interfaces",

		"pseudo-ethernet",
		o.ParentIDInterfacesPseudoEthernet.ValueString(),

		"vif-s",
		o.ParentIDInterfacesPseudoEthernetVifS.ValueString(),

		"dhcpv6-options",

		"pd",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesPseudoEthernetVifSDhcpvsixOptionsPd) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"pd_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `DHCPv6 prefix delegation interface statement

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  instance number  &emsp; |  Prefix delegation instance (>= 0)  |

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

		"vif_s_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `QinQ TAG-S Virtual Local Area Network (VLAN) ID

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-4094  &emsp; |  QinQ Virtual Local Area Network (VLAN) ID  |

`,
		},

		// LeafNodes

		"length": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Request IPv6 prefix length from peer

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 32-64  &emsp; |  Length of delegated prefix  |

`,

			// Default:          stringdefault.StaticString(`64`),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesPseudoEthernetVifSDhcpvsixOptionsPd) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesPseudoEthernetVifSDhcpvsixOptionsPd) UnmarshalJSON(_ []byte) error {
	return nil
}
