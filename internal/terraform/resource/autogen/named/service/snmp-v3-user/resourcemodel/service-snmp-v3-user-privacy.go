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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &ServiceSnmpVthreeUserPrivacy{}

// ServiceSnmpVthreeUserPrivacy describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type ServiceSnmpVthreeUserPrivacy struct {
	// LeafNodes
	LeafServiceSnmpVthreeUserPrivacyEncryptedPassword types.String `tfsdk:"encrypted_password" vyos:"encrypted-password,omitempty"`
	LeafServiceSnmpVthreeUserPrivacyPlaintextPassword types.String `tfsdk:"plaintext_password" vyos:"plaintext-password,omitempty"`
	LeafServiceSnmpVthreeUserPrivacyType              types.String `tfsdk:"type" vyos:"type,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceSnmpVthreeUserPrivacy) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"encrypted_password":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Defines the encrypted key for privacy protocol

`,
			Description: `Defines the encrypted key for privacy protocol

`,
		},

		"plaintext_password":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Defines the clear text key for privacy protocol

`,
			Description: `Defines the clear text key for privacy protocol

`,
		},

		"type":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
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

		// TagNodes

		// Nodes

	}
}
