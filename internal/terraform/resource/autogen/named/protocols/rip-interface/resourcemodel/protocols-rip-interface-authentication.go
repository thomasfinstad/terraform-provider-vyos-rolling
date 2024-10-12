// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &ProtocolsRIPInterfaceAuthentication{}

// ProtocolsRIPInterfaceAuthentication describes the resource data model.
type ProtocolsRIPInterfaceAuthentication struct {
	// LeafNodes
	LeafProtocolsRIPInterfaceAuthenticationPlaintextPassword types.String `tfsdk:"plaintext_password" vyos:"plaintext-password,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	ExistsTagProtocolsRIPInterfaceAuthenticationMdfive bool `tfsdk:"-" vyos:"md5,child"`

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsRIPInterfaceAuthentication) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"plaintext_password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Plain text password

    |  Format  |  Description                                  |
    |----------|-----------------------------------------------|
    |  txt     |  Plain text password (16 characters or less)  |
`,
			Description: `Plain text password

    |  Format  |  Description                                  |
    |----------|-----------------------------------------------|
    |  txt     |  Plain text password (16 characters or less)  |
`,
		},

		// Nodes

	}
}
