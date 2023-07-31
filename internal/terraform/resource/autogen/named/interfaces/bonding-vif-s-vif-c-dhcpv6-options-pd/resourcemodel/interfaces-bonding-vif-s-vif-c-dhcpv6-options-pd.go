// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesBondingVifSVifCDhcpvsixOptionsPd describes the resource data model.
type InterfacesBondingVifSVifCDhcpvsixOptionsPd struct {
	ID types.String `tfsdk:"identifier" vyos:",self-id"`

	ParentIDInterfacesBonding types.String `tfsdk:"bonding" vyos:"bonding_identifier,parent-id"`

	ParentIDInterfacesBondingVifS types.String `tfsdk:"vif_s" vyos:"vif-s_identifier,parent-id"`

	ParentIDInterfacesBondingVifSVifC types.String `tfsdk:"vif_c" vyos:"vif-c_identifier,parent-id"`

	// LeafNodes
	LeafInterfacesBondingVifSVifCDhcpvsixOptionsPdLength types.Number `tfsdk:"length" vyos:"length,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagInterfacesBondingVifSVifCDhcpvsixOptionsPdInterface bool `tfsdk:"interface" vyos:"interface,child"`

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesBondingVifSVifCDhcpvsixOptionsPd) GetVyosPath() []string {
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
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesBondingVifSVifCDhcpvsixOptionsPd) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `DHCPv6 prefix delegation interface statement

    |  Format  |  Description  |
    |----------|---------------|
    |  instance number  |  Prefix delegation instance (>= 0)  |

`,
		},

		"bonding_identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Bonding Interface/Link Aggregation

    |  Format  |  Description  |
    |----------|---------------|
    |  bondN  |  Bonding interface name  |

`,
		},

		"vif_s_identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `QinQ TAG-S Virtual Local Area Network (VLAN) ID

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:0-4094  |  QinQ Virtual Local Area Network (VLAN) ID  |

`,
		},

		"vif_c_identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `QinQ TAG-C Virtual Local Area Network (VLAN) ID

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
func (o *InterfacesBondingVifSVifCDhcpvsixOptionsPd) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesBondingVifSVifCDhcpvsixOptionsPd) UnmarshalJSON(_ []byte) error {
	return nil
}
