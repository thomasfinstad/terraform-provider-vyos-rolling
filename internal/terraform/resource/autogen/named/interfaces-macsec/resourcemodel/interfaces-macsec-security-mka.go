// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// InterfacesMacsecSecURItyMka describes the resource data model.
type InterfacesMacsecSecURItyMka struct {
	// LeafNodes
	InterfacesMacsecSecURItyMkaCak      customtypes.CustomStringValue `tfsdk:"cak" json:"cak,omitempty"`
	InterfacesMacsecSecURItyMkaCkn      customtypes.CustomStringValue `tfsdk:"ckn" json:"ckn,omitempty"`
	InterfacesMacsecSecURItyMkaPriority customtypes.CustomStringValue `tfsdk:"priority" json:"priority,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o InterfacesMacsecSecURItyMka) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"cak": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Secure Connectivity Association Key

|  Format  |  Description  |
|----------|---------------|
|  txt  |  16-byte (128-bit) hex-string (32 hex-digits) for gcm-aes-128 or 32-byte (256-bit) hex-string (64 hex-digits) for gcm-aes-256  |

`,
		},

		"ckn": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Secure Connectivity Association Key Name

|  Format  |  Description  |
|----------|---------------|
|  txt  |  1..32-bytes (8..256 bit) hex-string (2..64 hex-digits)  |

`,
		},

		"priority": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Priority of MACsec Key Agreement protocol (MKA) actor

|  Format  |  Description  |
|----------|---------------|
|  u32:0-255  |  MACsec Key Agreement protocol (MKA) priority  |

`,

			// Default:          stringdefault.StaticString(`255`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
