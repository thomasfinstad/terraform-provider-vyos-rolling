// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// HighAvailabilityVrrpGroupAuthentication describes the resource data model.
type HighAvailabilityVrrpGroupAuthentication struct {
	// LeafNodes
	LeafHighAvailabilityVrrpGroupAuthenticationPassword types.String `tfsdk:"password" vyos:"password,omitempty"`
	LeafHighAvailabilityVrrpGroupAuthenticationType     types.String `tfsdk:"type" vyos:"type,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o HighAvailabilityVrrpGroupAuthentication) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VRRP password

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  Password string (up to 8 characters)  |

`,
		},

		"type": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Authentication type

    |  Format  |  Description  |
    |----------|---------------|
    |  plaintext-password  |  Simple password string  |
    |  ah  |  AH - IPSEC (not recommended)  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *HighAvailabilityVrrpGroupAuthentication) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *HighAvailabilityVrrpGroupAuthentication) UnmarshalJSON(_ []byte) error {
	return nil
}
