// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &SystemLoginUserAuthentication{}

// SystemLoginUserAuthentication describes the resource data model.
type SystemLoginUserAuthentication struct {
	// LeafNodes
	LeafSystemLoginUserAuthenticationEncryptedPassword types.String `tfsdk:"encrypted_password" vyos:"encrypted-password,omitempty"`
	LeafSystemLoginUserAuthenticationPlaintextPassword types.String `tfsdk:"plaintext_password" vyos:"plaintext-password,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	ExistsTagSystemLoginUserAuthenticationPublicKeys bool `tfsdk:"public_keys" vyos:"public-keys,child"`

	// Nodes
	NodeSystemLoginUserAuthenticationOtp *SystemLoginUserAuthenticationOtp `tfsdk:"otp" vyos:"otp,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemLoginUserAuthentication) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"encrypted_password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Encrypted password

`,
			Description: `Encrypted password

`,

			// Default:          stringdefault.StaticString(`!`),
			Computed: true,
		},

		"plaintext_password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Plaintext password used for encryption

`,
			Description: `Plaintext password used for encryption

`,
		},

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