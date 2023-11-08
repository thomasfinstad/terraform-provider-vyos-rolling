// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsIsisDomainPassword describes the resource data model.
type VrfNameProtocolsIsisDomainPassword struct {
	// LeafNodes
	LeafVrfNameProtocolsIsisDomainPasswordPlaintextPassword types.String `tfsdk:"plaintext_password" vyos:"plaintext-password,omitempty"`
	LeafVrfNameProtocolsIsisDomainPasswordMdfive            types.String `tfsdk:"md5" vyos:"md5,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisDomainPassword) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"plaintext_password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Plain-text authentication type

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Circuit password  |

`,
		},

		"md5": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `MD5 authentication type

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Level-wide password  |

`,
		},

		// Nodes

	}
}
