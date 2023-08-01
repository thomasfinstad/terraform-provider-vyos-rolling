// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesMacsecSecURIty describes the resource data model.
type InterfacesMacsecSecURIty struct {
	// LeafNodes
	LeafInterfacesMacsecSecURItyCIPher       types.String `tfsdk:"cipher" vyos:"cipher,omitempty"`
	LeafInterfacesMacsecSecURItyEncrypt      types.Bool   `tfsdk:"encrypt" vyos:"encrypt,omitempty"`
	LeafInterfacesMacsecSecURItyReplayWindow types.Number `tfsdk:"replay_window" vyos:"replay-window,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeInterfacesMacsecSecURItyMka *InterfacesMacsecSecURItyMka `tfsdk:"mka" vyos:"mka,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesMacsecSecURIty) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"cipher": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Cipher suite used

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  gcm-aes-128  &emsp; |  Galois/Counter Mode of AES cipher with 128-bit key  |
    |  gcm-aes-256  &emsp; |  Galois/Counter Mode of AES cipher with 256-bit key  |

`,
		},

		"encrypt": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable optional MACsec encryption

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"replay_window": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `IEEE 802.1X/MACsec replay protection window

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0  &emsp; |  No replay window, strict check  |
    |  number: 1-4294967295  &emsp; |  Number of packets that could be misordered  |

`,
		},

		// Nodes

		"mka": schema.SingleNestedAttribute{
			Attributes: InterfacesMacsecSecURItyMka{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `MACsec Key Agreement protocol (MKA)

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesMacsecSecURIty) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesMacsecSecURIty) UnmarshalJSON(_ []byte) error {
	return nil
}