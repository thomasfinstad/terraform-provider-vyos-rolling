// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesPppoeDhcpvsixOptionsPd describes the resource data model.
type InterfacesPppoeDhcpvsixOptionsPd struct {
	SelfIdentifier types.String `tfsdk:"pd_id" vyos:",self-id"`

	ParentIDInterfacesPppoe types.String `tfsdk:"pppoe" vyos:"pppoe,parent-id"`

	// LeafNodes
	LeafInterfacesPppoeDhcpvsixOptionsPdLength types.Number `tfsdk:"length" vyos:"length,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagInterfacesPppoeDhcpvsixOptionsPdInterface bool `tfsdk:"interface" vyos:"interface,child"`

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesPppoeDhcpvsixOptionsPd) GetVyosPath() []string {
	return []string{
		"interfaces",

		"pppoe",
		o.ParentIDInterfacesPppoe.ValueString(),

		"dhcpv6-options",

		"pd",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesPppoeDhcpvsixOptionsPd) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"pd_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `DHCPv6 prefix delegation interface statement

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  instance number  &emsp; |  Prefix delegation instance (>= 0)  |

`,
		},

		"pppoe_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Point-to-Point Protocol over Ethernet (PPPoE) Interface

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  pppoeN  &emsp; |  PPPoE dialer interface name  |

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
func (o *InterfacesPppoeDhcpvsixOptionsPd) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesPppoeDhcpvsixOptionsPd) UnmarshalJSON(_ []byte) error {
	return nil
}
