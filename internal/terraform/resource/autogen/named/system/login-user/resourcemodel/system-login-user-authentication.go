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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (authentication) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &SystemLoginUserAuthentication{}

// SystemLoginUserAuthentication describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type SystemLoginUserAuthentication struct {
	// LeafNodes
	LeafSystemLoginUserAuthenticationEncryptedPassword types.String `tfsdk:"encrypted_password" vyos:"encrypted-password,omitempty"`
	LeafSystemLoginUserAuthenticationPlaintextPassword types.String `tfsdk:"plaintext_password" vyos:"plaintext-password,omitempty"`

	// TagNodes

	ExistsTagSystemLoginUserAuthenticationPublicKeys bool `tfsdk:"-" vyos:"public-keys,child"`

	// Nodes

	NodeSystemLoginUserAuthenticationOtp *SystemLoginUserAuthenticationOtp `tfsdk:"otp" vyos:"otp,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemLoginUserAuthentication) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"encrypted_password":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (encrypted-password) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Encrypted password

`,
			Description: `Encrypted password

`,

			// Default:          stringdefault.StaticString(`!`),
			Computed: true,
		},

		"plaintext_password":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (plaintext-password) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Plaintext password used for encryption

`,
			Description: `Plaintext password used for encryption

`,
		},

		// TagNodes

		// Nodes

		"otp": schema.SingleNestedAttribute{
			Attributes: SystemLoginUserAuthenticationOtp{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `One-Time-Pad (two-factor) authentication parameters

`,
			Description: `One-Time-Pad (two-factor) authentication parameters

`,
		},
	}
}
