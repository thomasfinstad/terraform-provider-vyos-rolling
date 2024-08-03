// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &SystemLoginUserAuthenticationOtp{}

// SystemLoginUserAuthenticationOtp describes the resource data model.
type SystemLoginUserAuthenticationOtp struct {
	// LeafNodes
	LeafSystemLoginUserAuthenticationOtpRateLimit  types.Number `tfsdk:"rate_limit" vyos:"rate-limit,omitempty"`
	LeafSystemLoginUserAuthenticationOtpRateTime   types.Number `tfsdk:"rate_time" vyos:"rate-time,omitempty"`
	LeafSystemLoginUserAuthenticationOtpWindowSize types.Number `tfsdk:"window_size" vyos:"window-size,omitempty"`
	LeafSystemLoginUserAuthenticationOtpKey        types.String `tfsdk:"key" vyos:"key,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemLoginUserAuthenticationOtp) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"rate_limit": schema.NumberAttribute{
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

		"rate_time": schema.NumberAttribute{
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

		"window_size": schema.NumberAttribute{
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

		"key": schema.StringAttribute{
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

		// Nodes

	}
}
