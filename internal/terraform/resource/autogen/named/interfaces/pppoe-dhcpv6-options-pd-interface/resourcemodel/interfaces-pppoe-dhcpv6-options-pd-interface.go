// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesPppoeDhcpvsixOptionsPdInterface describes the resource data model.
type InterfacesPppoeDhcpvsixOptionsPdInterface struct {
	ID types.String `tfsdk:"identifier" vyos:",self-id"`

	ParentIDInterfacesPppoe types.String `tfsdk:"pppoe" vyos:"pppoe_identifier,parent-id"`

	ParentIDInterfacesPppoeDhcpvsixOptionsPd types.String `tfsdk:"pd" vyos:"pd_identifier,parent-id"`

	// LeafNodes
	LeafInterfacesPppoeDhcpvsixOptionsPdInterfaceAddress types.String `tfsdk:"address" vyos:"address,omitempty"`
	LeafInterfacesPppoeDhcpvsixOptionsPdInterfaceSLAID   types.Number `tfsdk:"sla_id" vyos:"sla-id,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesPppoeDhcpvsixOptionsPdInterface) GetVyosPath() []string {
	return []string{
		"interfaces",

		"pppoe",
		o.ParentIDInterfacesPppoe.ValueString(),

		"dhcpv6-options",

		"pd",
		o.ParentIDInterfacesPppoeDhcpvsixOptionsPd.ValueString(),

		"interface",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesPppoeDhcpvsixOptionsPdInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Delegate IPv6 prefix from provider to this interface

`,
		},

		"pppoe_identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Point-to-Point Protocol over Ethernet (PPPoE) Interface

    |  Format  |  Description  |
    |----------|---------------|
    |  pppoeN  |  PPPoE dialer interface name  |

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
func (o *InterfacesPppoeDhcpvsixOptionsPdInterface) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesPppoeDhcpvsixOptionsPdInterface) UnmarshalJSON(_ []byte) error {
	return nil
}
