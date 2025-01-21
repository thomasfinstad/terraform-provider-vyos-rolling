/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (password) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &ProtocolsIsisInterfacePassword{}

// ProtocolsIsisInterfacePassword describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type ProtocolsIsisInterfacePassword struct {
	// LeafNodes
	LeafProtocolsIsisInterfacePasswordPlaintextPassword types.String `tfsdk:"plaintext_password" vyos:"plaintext-password,omitempty"`
	LeafProtocolsIsisInterfacePasswordMdfive            types.String `tfsdk:"md5" vyos:"md5,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsIsisInterfacePassword) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"plaintext_password":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (plaintext-password) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Plain-text authentication type

    |  Format  |  Description       |
    |----------|--------------------|
    |  txt     |  Circuit password  |
`,
			Description: `Plain-text authentication type

    |  Format  |  Description       |
    |----------|--------------------|
    |  txt     |  Circuit password  |
`,
		},

		"md5":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (md5) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `MD5 authentication type

    |  Format  |  Description          |
    |----------|-----------------------|
    |  txt     |  Level-wide password  |
`,
			Description: `MD5 authentication type

    |  Format  |  Description          |
    |----------|-----------------------|
    |  txt     |  Level-wide password  |
`,
		},

		// TagNodes

		// Nodes

	}
}
