// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesBrIDgeIPvsix describes the resource data model.
type InterfacesBrIDgeIPvsix struct {
	// LeafNodes
	LeafInterfacesBrIDgeIPvsixAdjustMss              types.String `tfsdk:"adjust_mss" vyos:"adjust-mss,omitempty"`
	LeafInterfacesBrIDgeIPvsixDisableForwarding      types.Bool   `tfsdk:"disable_forwarding" vyos:"disable-forwarding,omitempty"`
	LeafInterfacesBrIDgeIPvsixDupAddrDetectTransmits types.Number `tfsdk:"dup_addr_detect_transmits" vyos:"dup-addr-detect-transmits,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeInterfacesBrIDgeIPvsixAddress *InterfacesBrIDgeIPvsixAddress `tfsdk:"address" vyos:"address,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesBrIDgeIPvsix) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"adjust_mss": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Adjust TCP MSS value

    |  Format  |  Description  |
    |----------|---------------|
    |  clamp-mss-to-pmtu  |  Automatically sets the MSS to the proper value  |
    |  u32:536-65535  |  TCP Maximum segment size in bytes  |

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

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:0  |  Disable Duplicate Address Dectection (DAD)  |
    |  u32:1-n  |  Number of NS messages to send while performing DAD  |

`,
		},

		// Nodes

		"address": schema.SingleNestedAttribute{
			Attributes: InterfacesBrIDgeIPvsixAddress{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 address configuration modes

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesBrIDgeIPvsix) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesBrIDgeIPvsix) UnmarshalJSON(_ []byte) error {
	return nil
}
