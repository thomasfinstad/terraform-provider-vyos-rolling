// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &ServiceSnmpVthreeTrapTargetPrivacy{}

// ServiceSnmpVthreeTrapTargetPrivacy describes the resource data model.
type ServiceSnmpVthreeTrapTargetPrivacy struct {
	// LeafNodes
	LeafServiceSnmpVthreeTrapTargetPrivacyEncryptedPassword types.String `tfsdk:"encrypted_password" vyos:"encrypted-password,omitempty"`
	LeafServiceSnmpVthreeTrapTargetPrivacyPlaintextPassword types.String `tfsdk:"plaintext_password" vyos:"plaintext-password,omitempty"`
	LeafServiceSnmpVthreeTrapTargetPrivacyType              types.String `tfsdk:"type" vyos:"type,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceSnmpVthreeTrapTargetPrivacy) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"encrypted_password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Defines the encrypted key for privacy protocol

`,
			Description: `Defines the encrypted key for privacy protocol

`,
		},

		"plaintext_password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Defines the clear text key for privacy protocol

`,
			Description: `Defines the clear text key for privacy protocol

`,
		},

		"type": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Defines the protocol for privacy

    |  Format  |  Description                   |
    |----------|--------------------------------|
    |  des     |  Data Encryption Standard      |
    |  aes     |  Advanced Encryption Standard  |
`,
			Description: `Defines the protocol for privacy

    |  Format  |  Description                   |
    |----------|--------------------------------|
    |  des     |  Data Encryption Standard      |
    |  aes     |  Advanced Encryption Standard  |
`,

			// Default:          stringdefault.StaticString(`des`),
			Computed: true,
		},

		// Nodes

	}
}
