/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &InterfacesOpenvpnEncryption{}

// InterfacesOpenvpnEncryption describes the resource data model.
type InterfacesOpenvpnEncryption struct {
	// LeafNodes
	LeafInterfacesOpenvpnEncryptionCIPher      types.String `tfsdk:"cipher" vyos:"cipher,omitempty"`
	LeafInterfacesOpenvpnEncryptionDataCIPhers types.List   `tfsdk:"data_ciphers" vyos:"data-ciphers,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesOpenvpnEncryption) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"cipher":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Standard Data Encryption Algorithm

    |  Format     |  Description                           |
    |-------------|----------------------------------------|
    |  none       |  Disable encryption                    |
    |  3des       |  DES algorithm with triple encryption  |
    |  aes128     |  AES algorithm with 128-bit key CBC    |
    |  aes128gcm  |  AES algorithm with 128-bit key GCM    |
    |  aes192     |  AES algorithm with 192-bit key CBC    |
    |  aes192gcm  |  AES algorithm with 192-bit key GCM    |
    |  aes256     |  AES algorithm with 256-bit key CBC    |
    |  aes256gcm  |  AES algorithm with 256-bit key GCM    |
`,
			Description: `Standard Data Encryption Algorithm

    |  Format     |  Description                           |
    |-------------|----------------------------------------|
    |  none       |  Disable encryption                    |
    |  3des       |  DES algorithm with triple encryption  |
    |  aes128     |  AES algorithm with 128-bit key CBC    |
    |  aes128gcm  |  AES algorithm with 128-bit key GCM    |
    |  aes192     |  AES algorithm with 192-bit key CBC    |
    |  aes192gcm  |  AES algorithm with 192-bit key GCM    |
    |  aes256     |  AES algorithm with 256-bit key CBC    |
    |  aes256gcm  |  AES algorithm with 256-bit key GCM    |
`,
		},

		"data_ciphers":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype-multi.gotmpl */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Cipher negotiation list for use in server or client mode

    |  Format     |  Description                           |
    |-------------|----------------------------------------|
    |  none       |  Disable encryption                    |
    |  3des       |  DES algorithm with triple encryption  |
    |  aes128     |  AES algorithm with 128-bit key CBC    |
    |  aes128gcm  |  AES algorithm with 128-bit key GCM    |
    |  aes192     |  AES algorithm with 192-bit key CBC    |
    |  aes192gcm  |  AES algorithm with 192-bit key GCM    |
    |  aes256     |  AES algorithm with 256-bit key CBC    |
    |  aes256gcm  |  AES algorithm with 256-bit key GCM    |
`,
			Description: `Cipher negotiation list for use in server or client mode

    |  Format     |  Description                           |
    |-------------|----------------------------------------|
    |  none       |  Disable encryption                    |
    |  3des       |  DES algorithm with triple encryption  |
    |  aes128     |  AES algorithm with 128-bit key CBC    |
    |  aes128gcm  |  AES algorithm with 128-bit key GCM    |
    |  aes192     |  AES algorithm with 192-bit key CBC    |
    |  aes192gcm  |  AES algorithm with 192-bit key GCM    |
    |  aes256     |  AES algorithm with 256-bit key CBC    |
    |  aes256gcm  |  AES algorithm with 256-bit key GCM    |
`,
		},

		// Nodes

	}
}
