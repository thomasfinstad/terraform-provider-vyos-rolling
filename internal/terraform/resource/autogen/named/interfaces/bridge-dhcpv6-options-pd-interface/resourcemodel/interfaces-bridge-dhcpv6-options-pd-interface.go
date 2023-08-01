// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesBrIDgeDhcpvsixOptionsPdInterface describes the resource data model.
type InterfacesBrIDgeDhcpvsixOptionsPdInterface struct {
	SelfIdentifier types.String `tfsdk:"interface_id" vyos:",self-id"`

	ParentIDInterfacesBrIDge types.String `tfsdk:"bridge" vyos:"bridge,parent-id"`

	ParentIDInterfacesBrIDgeDhcpvsixOptionsPd types.String `tfsdk:"pd" vyos:"pd,parent-id"`

	// LeafNodes
	LeafInterfacesBrIDgeDhcpvsixOptionsPdInterfaceAddress types.String `tfsdk:"address" vyos:"address,omitempty"`
	LeafInterfacesBrIDgeDhcpvsixOptionsPdInterfaceSLAID   types.Number `tfsdk:"sla_id" vyos:"sla-id,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesBrIDgeDhcpvsixOptionsPdInterface) GetVyosPath() []string {
	return []string{
		"interfaces",

		"bridge",
		o.ParentIDInterfacesBrIDge.ValueString(),

		"dhcpv6-options",

		"pd",
		o.ParentIDInterfacesBrIDgeDhcpvsixOptionsPd.ValueString(),

		"interface",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesBrIDgeDhcpvsixOptionsPdInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"interface_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Delegate IPv6 prefix from provider to this interface

`,
		},

		"bridge_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Bridge Interface

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  brN  &emsp; |  Bridge interface name  |

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
func (o *InterfacesBrIDgeDhcpvsixOptionsPdInterface) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesBrIDgeDhcpvsixOptionsPdInterface) UnmarshalJSON(_ []byte) error {
	return nil
}
