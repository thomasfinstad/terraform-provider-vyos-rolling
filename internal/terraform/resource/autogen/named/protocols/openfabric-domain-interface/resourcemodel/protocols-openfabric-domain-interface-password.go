/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &ProtocolsOpenfabricDomainInterfacePassword{}

// ProtocolsOpenfabricDomainInterfacePassword describes the resource data model.
type ProtocolsOpenfabricDomainInterfacePassword struct {
	// LeafNodes
	LeafProtocolsOpenfabricDomainInterfacePasswordPlaintextPassword types.String `tfsdk:"plaintext_password" vyos:"plaintext-password,omitempty"`
	LeafProtocolsOpenfabricDomainInterfacePasswordMdfive            types.String `tfsdk:"md5" vyos:"md5,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsOpenfabricDomainInterfacePassword) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"plaintext_password":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
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

		"md5":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
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
