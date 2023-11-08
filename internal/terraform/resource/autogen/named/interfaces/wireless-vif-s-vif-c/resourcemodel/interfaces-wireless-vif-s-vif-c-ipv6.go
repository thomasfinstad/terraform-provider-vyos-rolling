// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesWirelessVifSVifCIPvsix describes the resource data model.
type InterfacesWirelessVifSVifCIPvsix struct {
	// LeafNodes
	LeafInterfacesWirelessVifSVifCIPvsixAdjustMss              types.String `tfsdk:"adjust_mss" vyos:"adjust-mss,omitempty"`
	LeafInterfacesWirelessVifSVifCIPvsixDisableForwarding      types.Bool   `tfsdk:"disable_forwarding" vyos:"disable-forwarding,omitempty"`
	LeafInterfacesWirelessVifSVifCIPvsixDupAddrDetectTransmits types.Number `tfsdk:"dup_addr_detect_transmits" vyos:"dup-addr-detect-transmits,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeInterfacesWirelessVifSVifCIPvsixAddress *InterfacesWirelessVifSVifCIPvsixAddress `tfsdk:"address" vyos:"address,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWirelessVifSVifCIPvsix) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"adjust_mss": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Adjust TCP MSS value

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  clamp-mss-to-pmtu  &emsp; |  Automatically sets the MSS to the proper value  |
    |  number: 536-65535  &emsp; |  TCP Maximum segment size in bytes  |

`,
		},

		"disable_forwarding": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable IP forwarding on this interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"dup_addr_detect_transmits": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Number of NS messages to send while performing DAD (default: 1)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0  &emsp; |  Disable Duplicate Address Dectection (DAD)  |
    |  number: 1-n  &emsp; |  Number of NS messages to send while performing DAD  |

`,
		},

		// Nodes

		"address": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessVifSVifCIPvsixAddress{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 address configuration modes

`,
		},
	}
}
