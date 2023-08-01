// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VpnIPsecProfileAuthentication describes the resource data model.
type VpnIPsecProfileAuthentication struct {
	// LeafNodes
	LeafVpnIPsecProfileAuthenticationMode            types.String `tfsdk:"mode" vyos:"mode,omitempty"`
	LeafVpnIPsecProfileAuthenticationPreSharedSecret types.String `tfsdk:"pre_shared_secret" vyos:"pre-shared-secret,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnIPsecProfileAuthentication) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"mode": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Authentication mode

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  pre-shared-secret  &emsp; |  Use a pre-shared secret key  |

`,
		},

		"pre_shared_secret": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Pre-shared secret key

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Pre-shared secret key  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VpnIPsecProfileAuthentication) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VpnIPsecProfileAuthentication) UnmarshalJSON(_ []byte) error {
	return nil
}
