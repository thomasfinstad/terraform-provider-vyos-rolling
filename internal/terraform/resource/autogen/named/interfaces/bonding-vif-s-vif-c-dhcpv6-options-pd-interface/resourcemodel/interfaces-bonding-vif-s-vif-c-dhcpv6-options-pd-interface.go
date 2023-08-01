// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesBondingVifSVifCDhcpvsixOptionsPdInterface describes the resource data model.
type InterfacesBondingVifSVifCDhcpvsixOptionsPdInterface struct {
	SelfIdentifier types.String `tfsdk:"interface_id" vyos:",self-id"`

	ParentIDInterfacesBonding types.String `tfsdk:"bonding" vyos:"bonding,parent-id"`

	ParentIDInterfacesBondingVifS types.String `tfsdk:"vif_s" vyos:"vif-s,parent-id"`

	ParentIDInterfacesBondingVifSVifC types.String `tfsdk:"vif_c" vyos:"vif-c,parent-id"`

	ParentIDInterfacesBondingVifSVifCDhcpvsixOptionsPd types.String `tfsdk:"pd" vyos:"pd,parent-id"`

	// LeafNodes
	LeafInterfacesBondingVifSVifCDhcpvsixOptionsPdInterfaceAddress types.String `tfsdk:"address" vyos:"address,omitempty"`
	LeafInterfacesBondingVifSVifCDhcpvsixOptionsPdInterfaceSLAID   types.Number `tfsdk:"sla_id" vyos:"sla-id,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesBondingVifSVifCDhcpvsixOptionsPdInterface) GetVyosPath() []string {
	return []string{
		"interfaces",

		"bonding",
		o.ParentIDInterfacesBonding.ValueString(),

		"vif-s",
		o.ParentIDInterfacesBondingVifS.ValueString(),

		"vif-c",
		o.ParentIDInterfacesBondingVifSVifC.ValueString(),

		"dhcpv6-options",

		"pd",
		o.ParentIDInterfacesBondingVifSVifCDhcpvsixOptionsPd.ValueString(),

		"interface",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesBondingVifSVifCDhcpvsixOptionsPdInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"interface_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Delegate IPv6 prefix from provider to this interface

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

		"vif_c_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `QinQ TAG-C Virtual Local Area Network (VLAN) ID

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
func (o *InterfacesBondingVifSVifCDhcpvsixOptionsPdInterface) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesBondingVifSVifCDhcpvsixOptionsPdInterface) UnmarshalJSON(_ []byte) error {
	return nil
}
