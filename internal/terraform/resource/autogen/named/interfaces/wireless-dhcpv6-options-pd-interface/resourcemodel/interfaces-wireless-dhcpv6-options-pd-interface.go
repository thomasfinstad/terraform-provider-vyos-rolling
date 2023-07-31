// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesWirelessDhcpvsixOptionsPdInterface describes the resource data model.
type InterfacesWirelessDhcpvsixOptionsPdInterface struct {
	ID types.String `tfsdk:"identifier" vyos:",self-id"`

	ParentIDInterfacesWireless types.String `tfsdk:"wireless" vyos:"wireless_identifier,parent-id"`

	ParentIDInterfacesWirelessDhcpvsixOptionsPd types.String `tfsdk:"pd" vyos:"pd_identifier,parent-id"`

	// LeafNodes
	LeafInterfacesWirelessDhcpvsixOptionsPdInterfaceAddress types.String `tfsdk:"address" vyos:"address,omitempty"`
	LeafInterfacesWirelessDhcpvsixOptionsPdInterfaceSLAID   types.Number `tfsdk:"sla_id" vyos:"sla-id,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesWirelessDhcpvsixOptionsPdInterface) GetVyosPath() []string {
	return []string{
		"interfaces",

		"wireless",
		o.ParentIDInterfacesWireless.ValueString(),

		"dhcpv6-options",

		"pd",
		o.ParentIDInterfacesWirelessDhcpvsixOptionsPd.ValueString(),

		"interface",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWirelessDhcpvsixOptionsPdInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Delegate IPv6 prefix from provider to this interface

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
func (o *InterfacesWirelessDhcpvsixOptionsPdInterface) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesWirelessDhcpvsixOptionsPdInterface) UnmarshalJSON(_ []byte) error {
	return nil
}
