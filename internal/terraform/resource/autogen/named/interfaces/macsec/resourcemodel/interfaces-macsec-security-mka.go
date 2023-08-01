// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesMacsecSecURItyMka describes the resource data model.
type InterfacesMacsecSecURItyMka struct {
	// LeafNodes
	LeafInterfacesMacsecSecURItyMkaCak      types.String `tfsdk:"cak" vyos:"cak,omitempty"`
	LeafInterfacesMacsecSecURItyMkaCkn      types.String `tfsdk:"ckn" vyos:"ckn,omitempty"`
	LeafInterfacesMacsecSecURItyMkaPriority types.Number `tfsdk:"priority" vyos:"priority,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesMacsecSecURItyMka) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"cak": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Secure Connectivity Association Key

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  16-byte (128-bit) hex-string (32 hex-digits) for gcm-aes-128 or 32-byte (256-bit) hex-string (64 hex-digits) for gcm-aes-256  |

`,
		},

		"ckn": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Secure Connectivity Association Key Name

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  1..32-bytes (8..256 bit) hex-string (2..64 hex-digits)  |

`,
		},

		"priority": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Priority of MACsec Key Agreement protocol (MKA) actor

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-255  &emsp; |  MACsec Key Agreement protocol (MKA) priority  |

`,

			// Default:          stringdefault.StaticString(`255`),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesMacsecSecURItyMka) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesMacsecSecURItyMka) UnmarshalJSON(_ []byte) error {
	return nil
}
