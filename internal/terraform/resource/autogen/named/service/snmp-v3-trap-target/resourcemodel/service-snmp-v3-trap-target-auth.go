// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ServiceSnmpVthreeTrapTargetAuth describes the resource data model.
type ServiceSnmpVthreeTrapTargetAuth struct {
	// LeafNodes
	LeafServiceSnmpVthreeTrapTargetAuthEncryptedPassword types.String `tfsdk:"encrypted_password" vyos:"encrypted-password,omitempty"`
	LeafServiceSnmpVthreeTrapTargetAuthPlaintextPassword types.String `tfsdk:"plaintext_password" vyos:"plaintext-password,omitempty"`
	LeafServiceSnmpVthreeTrapTargetAuthType              types.String `tfsdk:"type" vyos:"type,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceSnmpVthreeTrapTargetAuth) ResourceSchemaAttributes() map[string]schema.Attribute {
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

    |  Format  |  Description  |
    |----------|---------------|
    |  md5  |  Message Digest 5  |
    |  sha  |  Secure Hash Algorithm  |

`,

			// Default:          stringdefault.StaticString(`md5`),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ServiceSnmpVthreeTrapTargetAuth) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ServiceSnmpVthreeTrapTargetAuth) UnmarshalJSON(_ []byte) error {
	return nil
}
