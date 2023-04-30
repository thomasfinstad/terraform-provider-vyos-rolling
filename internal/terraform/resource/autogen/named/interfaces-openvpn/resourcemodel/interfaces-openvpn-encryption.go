// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// InterfacesOpenvpnEncryption describes the resource data model.
type InterfacesOpenvpnEncryption struct {
	// LeafNodes
	InterfacesOpenvpnEncryptionCIPher     customtypes.CustomStringValue `tfsdk:"cipher" json:"cipher,omitempty"`
	InterfacesOpenvpnEncryptionNcpCIPhers customtypes.CustomStringValue `tfsdk:"ncp_ciphers" json:"ncp-ciphers,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o InterfacesOpenvpnEncryption) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"cipher": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Standard Data Encryption Algorithm

|  Format  |  Description  |
|----------|---------------|
|  none  |  Disable encryption  |
|  des  |  DES algorithm  |
|  3des  |  DES algorithm with triple encryption  |
|  bf128  |  Blowfish algorithm with 128-bit key  |
|  bf256  |  Blowfish algorithm with 256-bit key  |
|  aes128  |  AES algorithm with 128-bit key CBC  |
|  aes128gcm  |  AES algorithm with 128-bit key GCM  |
|  aes192  |  AES algorithm with 192-bit key CBC  |
|  aes192gcm  |  AES algorithm with 192-bit key GCM  |
|  aes256  |  AES algorithm with 256-bit key CBC  |
|  aes256gcm  |  AES algorithm with 256-bit key GCM  |

`,
		},

		"ncp_ciphers": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Cipher negotiation list for use in server or client mode

|  Format  |  Description  |
|----------|---------------|
|  none  |  Disable encryption  |
|  des  |  DES algorithm  |
|  3des  |  DES algorithm with triple encryption  |
|  aes128  |  AES algorithm with 128-bit key CBC  |
|  aes128gcm  |  AES algorithm with 128-bit key GCM  |
|  aes192  |  AES algorithm with 192-bit key CBC  |
|  aes192gcm  |  AES algorithm with 192-bit key GCM  |
|  aes256  |  AES algorithm with 256-bit key CBC  |
|  aes256gcm  |  AES algorithm with 256-bit key GCM  |

`,
		},

		// TagNodes

		// Nodes

	}
}
