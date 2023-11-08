// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ServiceSnmpVthreeUserAuth describes the resource data model.
type ServiceSnmpVthreeUserAuth struct {
	// LeafNodes
	LeafServiceSnmpVthreeUserAuthEncryptedPassword types.String `tfsdk:"encrypted_password" vyos:"encrypted-password,omitempty"`
	LeafServiceSnmpVthreeUserAuthPlaintextPassword types.String `tfsdk:"plaintext_password" vyos:"plaintext-password,omitempty"`
	LeafServiceSnmpVthreeUserAuthType              types.String `tfsdk:"type" vyos:"type,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceSnmpVthreeUserAuth) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"encrypted_password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Defines the encrypted key for authentication

`,
		},

		"plaintext_password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Defines the clear text key for authentication

`,
		},

		"type": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Define used protocol

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  md5  &emsp; |  Message Digest 5  |
    |  sha  &emsp; |  Secure Hash Algorithm  |

`,

			// Default:          stringdefault.StaticString(`md5`),
			Computed: true,
		},

		// Nodes

	}
}
