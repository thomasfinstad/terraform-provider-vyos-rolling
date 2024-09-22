// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &ProtocolsOpenfabricDomainDomainPassword{}

// ProtocolsOpenfabricDomainDomainPassword describes the resource data model.
type ProtocolsOpenfabricDomainDomainPassword struct {
	// LeafNodes
	LeafProtocolsOpenfabricDomainDomainPasswordPlaintextPassword types.String `tfsdk:"plaintext_password" vyos:"plaintext-password,omitempty"`
	LeafProtocolsOpenfabricDomainDomainPasswordMdfive            types.String `tfsdk:"md5" vyos:"md5,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsOpenfabricDomainDomainPassword) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"plaintext_password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Use plain text password

    |  Format  |  Description              |
    |----------|---------------------------|
    |  txt     |  Authentication password  |
`,
			Description: `Use plain text password

    |  Format  |  Description              |
    |----------|---------------------------|
    |  txt     |  Authentication password  |
`,
		},

		"md5": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Use MD5 hash authentication

    |  Format  |  Description              |
    |----------|---------------------------|
    |  txt     |  Authentication password  |
`,
			Description: `Use MD5 hash authentication

    |  Format  |  Description              |
    |----------|---------------------------|
    |  txt     |  Authentication password  |
`,
		},

		// Nodes

	}
}