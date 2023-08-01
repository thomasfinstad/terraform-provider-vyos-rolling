// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VpnOpenconnectAuthenticationLocalUsersUsernameOtp describes the resource data model.
type VpnOpenconnectAuthenticationLocalUsersUsernameOtp struct {
	// LeafNodes
	LeafVpnOpenconnectAuthenticationLocalUsersUsernameOtpKey       types.String `tfsdk:"key" vyos:"key,omitempty"`
	LeafVpnOpenconnectAuthenticationLocalUsersUsernameOtpOtpLength types.Number `tfsdk:"otp_length" vyos:"otp-length,omitempty"`
	LeafVpnOpenconnectAuthenticationLocalUsersUsernameOtpInterval  types.Number `tfsdk:"interval" vyos:"interval,omitempty"`
	LeafVpnOpenconnectAuthenticationLocalUsersUsernameOtpTokenType types.String `tfsdk:"token_type" vyos:"token-type,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnOpenconnectAuthenticationLocalUsersUsernameOtp) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"key": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Token Key Secret key for the token algorithm (see RFC 4226)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  OTP key in hex-encoded format  |

`,
		},

		"otp_length": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Number of digits in OTP code

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 6-8  &emsp; |  Number of digits in OTP code  |

`,

			// Default:          stringdefault.StaticString(`6`),
			Computed: true,
		},

		"interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Time tokens interval in seconds

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 5-86400  &emsp; |  Time tokens interval in seconds.  |

`,

			// Default:          stringdefault.StaticString(`30`),
			Computed: true,
		},

		"token_type": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Token type

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  hotp-time  &emsp; |  Time-based OTP algorithm  |
    |  hotp-event  &emsp; |  Event-based OTP algorithm  |

`,

			// Default:          stringdefault.StaticString(`hotp-time`),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VpnOpenconnectAuthenticationLocalUsersUsernameOtp) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VpnOpenconnectAuthenticationLocalUsersUsernameOtp) UnmarshalJSON(_ []byte) error {
	return nil
}
