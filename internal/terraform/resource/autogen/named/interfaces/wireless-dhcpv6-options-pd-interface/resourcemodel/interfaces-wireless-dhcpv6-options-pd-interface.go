// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesWirelessDhcpvsixOptionsPdInterface describes the resource data model.
type InterfacesWirelessDhcpvsixOptionsPdInterface struct {
	SelfIdentifier types.String `tfsdk:"interface_id" vyos:",self-id"`

	ParentIDInterfacesWireless types.String `tfsdk:"wireless" vyos:"wireless,parent-id"`

	ParentIDInterfacesWirelessDhcpvsixOptionsPd types.String `tfsdk:"pd" vyos:"pd,parent-id"`

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
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWirelessDhcpvsixOptionsPdInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, an amalgamation of the `interface_id` and the parents `*_id` fields seperated by dunder `__` starting with top level ancestor.",
		},
		"interface_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Delegate IPv6 prefix from provider to this interface

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		"wireless_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Wireless (WiFi/WLAN) Network Interface

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  wlanN  &emsp; |  Wireless (WiFi/WLAN) interface name  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		"pd_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `DHCPv6 prefix delegation interface statement

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  instance number  &emsp; |  Prefix delegation instance (>= 0)  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
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
func (o *InterfacesWirelessDhcpvsixOptionsPdInterface) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesWirelessDhcpvsixOptionsPdInterface) UnmarshalJSON(_ []byte) error {
	return nil
}
