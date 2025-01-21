/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (auth) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &ServiceSnmpVthreeUserAuth{}

// ServiceSnmpVthreeUserAuth describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type ServiceSnmpVthreeUserAuth struct {
	// LeafNodes
	LeafServiceSnmpVthreeUserAuthEncryptedPassword types.String `tfsdk:"encrypted_password" vyos:"encrypted-password,omitempty"`
	LeafServiceSnmpVthreeUserAuthPlaintextPassword types.String `tfsdk:"plaintext_password" vyos:"plaintext-password,omitempty"`
	LeafServiceSnmpVthreeUserAuthType              types.String `tfsdk:"type" vyos:"type,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceSnmpVthreeUserAuth) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"encrypted_password":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (encrypted-password) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Defines the encrypted key for authentication

`,
			Description: `Defines the encrypted key for authentication

`,
		},

		"plaintext_password":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (plaintext-password) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Defines the clear text key for authentication

`,
			Description: `Defines the clear text key for authentication

`,
		},

		"type":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (type) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Define used protocol

    |  Format  |  Description            |
    |----------|-------------------------|
    |  md5     |  Message Digest 5       |
    |  sha     |  Secure Hash Algorithm  |
`,
			Description: `Define used protocol

    |  Format  |  Description            |
    |----------|-------------------------|
    |  md5     |  Message Digest 5       |
    |  sha     |  Secure Hash Algorithm  |
`,

			// Default:          stringdefault.StaticString(`md5`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
