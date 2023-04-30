// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// InterfacesBrIDgeVifIPvsix describes the resource data model.
type InterfacesBrIDgeVifIPvsix struct {
	// LeafNodes
	InterfacesBrIDgeVifIPvsixAdjustMss              customtypes.CustomStringValue `tfsdk:"adjust_mss" json:"adjust-mss,omitempty"`
	InterfacesBrIDgeVifIPvsixDisableForwarding      customtypes.CustomStringValue `tfsdk:"disable_forwarding" json:"disable-forwarding,omitempty"`
	InterfacesBrIDgeVifIPvsixDupAddrDetectTransmits customtypes.CustomStringValue `tfsdk:"dup_addr_detect_transmits" json:"dup-addr-detect-transmits,omitempty"`

	// TagNodes

	// Nodes
	InterfacesBrIDgeVifIPvsixAddress types.Object `tfsdk:"address" json:"address,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o InterfacesBrIDgeVifIPvsix) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"adjust_mss": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Adjust TCP MSS value

|  Format  |  Description  |
|----------|---------------|
|  clamp-mss-to-pmtu  |  Automatically sets the MSS to the proper value  |
|  u32:536-65535  |  TCP Maximum segment size in bytes  |

`,
		},

		"disable_forwarding": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Disable IP forwarding on this interface

`,
		},

		"dup_addr_detect_transmits": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Number of NS messages to send while performing DAD (default: 1)

|  Format  |  Description  |
|----------|---------------|
|  u32:0  |  Disable Duplicate Address Dectection (DAD)  |
|  u32:1-n  |  Number of NS messages to send while performing DAD  |

`,
		},

		// TagNodes

		// Nodes

		"address": schema.SingleNestedAttribute{
			Attributes: InterfacesBrIDgeVifIPvsixAddress{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 address configuration modes

`,
		},
	}
}
