// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SystemLoginUserAuthenticationOtp describes the resource data model.
type SystemLoginUserAuthenticationOtp struct {
	// LeafNodes
	LeafSystemLoginUserAuthenticationOtpRateLimit  types.Number `tfsdk:"rate_limit" vyos:"rate-limit,omitempty"`
	LeafSystemLoginUserAuthenticationOtpRateTime   types.Number `tfsdk:"rate_time" vyos:"rate-time,omitempty"`
	LeafSystemLoginUserAuthenticationOtpWindowSize types.Number `tfsdk:"window_size" vyos:"window-size,omitempty"`
	LeafSystemLoginUserAuthenticationOtpKey        types.String `tfsdk:"key" vyos:"key,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemLoginUserAuthenticationOtp) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"rate_limit": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Limit number of logins (rate-limit) per rate-time

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-10  &emsp; |  Number of attempts  |

`,

			// Default:          stringdefault.StaticString(`3`),
			Computed: true,
		},

		"rate_time": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Limit number of logins (rate-limit) per rate-time

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 15-600  &emsp; |  Time interval  |

`,

			// Default:          stringdefault.StaticString(`30`),
			Computed: true,
		},

		"window_size": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set window of concurrently valid codes

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-21  &emsp; |  Window size  |

`,

			// Default:          stringdefault.StaticString(`3`),
			Computed: true,
		},

		"key": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Key/secret the token algorithm (see RFC4226)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Base32 encoded key/token  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *SystemLoginUserAuthenticationOtp) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *SystemLoginUserAuthenticationOtp) UnmarshalJSON(_ []byte) error {
	return nil
}