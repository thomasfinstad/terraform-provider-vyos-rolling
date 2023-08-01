// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsOspfAreaVirtualLinkAuthentication describes the resource data model.
type VrfNameProtocolsOspfAreaVirtualLinkAuthentication struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfAreaVirtualLinkAuthenticationPlaintextPassword types.String `tfsdk:"plaintext_password" vyos:"plaintext-password,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsOspfAreaVirtualLinkAuthenticationMdfive *VrfNameProtocolsOspfAreaVirtualLinkAuthenticationMdfive `tfsdk:"md5" vyos:"md5,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfAreaVirtualLinkAuthentication) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"plaintext_password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Plain text password

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Plain text password (8 characters or less)  |

`,
		},

		// Nodes

		"md5": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfAreaVirtualLinkAuthenticationMdfive{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `MD5 key id

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsOspfAreaVirtualLinkAuthentication) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsOspfAreaVirtualLinkAuthentication) UnmarshalJSON(_ []byte) error {
	return nil
}
