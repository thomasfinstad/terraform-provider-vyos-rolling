// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesWirelessVifDhcpvsixOptionsPd describes the resource data model.
type InterfacesWirelessVifDhcpvsixOptionsPd struct {
	ID types.String `tfsdk:"identifier" vyos:",self-id"`

	ParentIDInterfacesWireless types.String `tfsdk:"wireless" vyos:"wireless_identifier,parent-id"`

	ParentIDInterfacesWirelessVif types.String `tfsdk:"vif" vyos:"vif_identifier,parent-id"`

	// LeafNodes
	LeafInterfacesWirelessVifDhcpvsixOptionsPdLength types.Number `tfsdk:"length" vyos:"length,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagInterfacesWirelessVifDhcpvsixOptionsPdInterface bool `tfsdk:"interface" vyos:"interface,child"`

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesWirelessVifDhcpvsixOptionsPd) GetVyosPath() []string {
	return []string{
		"interfaces",

		"wireless",
		o.ParentIDInterfacesWireless.ValueString(),

		"vif",
		o.ParentIDInterfacesWirelessVif.ValueString(),

		"dhcpv6-options",

		"pd",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWirelessVifDhcpvsixOptionsPd) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `DHCPv6 prefix delegation interface statement

    |  Format  |  Description  |
    |----------|---------------|
    |  instance number  |  Prefix delegation instance (>= 0)  |

`,
		},

		"wireless_identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Wireless (WiFi/WLAN) Network Interface

    |  Format  |  Description  |
    |----------|---------------|
    |  wlanN  |  Wireless (WiFi/WLAN) interface name  |

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

		// LeafNodes

		"length": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Request IPv6 prefix length from peer

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:32-64  |  Length of delegated prefix  |

`,

			// Default:          stringdefault.StaticString(`64`),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesWirelessVifDhcpvsixOptionsPd) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesWirelessVifDhcpvsixOptionsPd) UnmarshalJSON(_ []byte) error {
	return nil
}
