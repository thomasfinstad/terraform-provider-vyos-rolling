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

var _ helpers.VyosResourceDataModel = &SystemLoginUserAuthenticationOtp{}

// SystemLoginUserAuthenticationOtp describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type SystemLoginUserAuthenticationOtp struct {
	// LeafNodes
	LeafSystemLoginUserAuthenticationOtpRateLimit  types.Number `tfsdk:"rate_limit" vyos:"rate-limit,omitempty"`
	LeafSystemLoginUserAuthenticationOtpRateTime   types.Number `tfsdk:"rate_time" vyos:"rate-time,omitempty"`
	LeafSystemLoginUserAuthenticationOtpWindowSize types.Number `tfsdk:"window_size" vyos:"window-size,omitempty"`
	LeafSystemLoginUserAuthenticationOtpKey        types.String `tfsdk:"key" vyos:"key,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemLoginUserAuthenticationOtp) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"rate_limit":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Limit number of logins (rate-limit) per rate-time

    |  Format  |  Description         |
    |----------|----------------------|
    |  1-10    |  Number of attempts  |
`,
			Description: `Limit number of logins (rate-limit) per rate-time

    |  Format  |  Description         |
    |----------|----------------------|
    |  1-10    |  Number of attempts  |
`,

			// Default:          stringdefault.StaticString(`3`),
			Computed: true,
		},

		"rate_time":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Limit number of logins (rate-limit) per rate-time

    |  Format  |  Description    |
    |----------|-----------------|
    |  15-600  |  Time interval  |
`,
			Description: `Limit number of logins (rate-limit) per rate-time

    |  Format  |  Description    |
    |----------|-----------------|
    |  15-600  |  Time interval  |
`,

			// Default:          stringdefault.StaticString(`30`),
			Computed: true,
		},

		"window_size":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set window of concurrently valid codes

    |  Format  |  Description  |
    |----------|---------------|
    |  1-21    |  Window size  |
`,
			Description: `Set window of concurrently valid codes

    |  Format  |  Description  |
    |----------|---------------|
    |  1-21    |  Window size  |
`,

			// Default:          stringdefault.StaticString(`3`),
			Computed: true,
		},

		"key":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Key/secret the token algorithm (see RFC4226)

    |  Format  |  Description               |
    |----------|----------------------------|
    |  txt     |  Base32 encoded key/token  |
`,
			Description: `Key/secret the token algorithm (see RFC4226)

    |  Format  |  Description               |
    |----------|----------------------------|
    |  txt     |  Base32 encoded key/token  |
`,
		},

		// TagNodes

		// Nodes

	}
}
