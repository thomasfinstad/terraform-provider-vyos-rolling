// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// VrfNameProtocolsIsisDomainPassword describes the resource data model.
type VrfNameProtocolsIsisDomainPassword struct {
	// LeafNodes
	VrfNameProtocolsIsisDomainPasswordPlaintextPassword customtypes.CustomStringValue `tfsdk:"plaintext_password" json:"plaintext-password,omitempty"`
	VrfNameProtocolsIsisDomainPasswordMdfive            customtypes.CustomStringValue `tfsdk:"md5" json:"md5,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VrfNameProtocolsIsisDomainPassword) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"plaintext_password": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Plain-text authentication type

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Circuit password  |

`,
		},

		"md5": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `MD5 authentication type

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Level-wide password  |

`,
		},

		// TagNodes

		// Nodes

	}
}
