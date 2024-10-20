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

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsIsisAreaPassword{}

// VrfNameProtocolsIsisAreaPassword describes the resource data model.
type VrfNameProtocolsIsisAreaPassword struct {
	// LeafNodes
	LeafVrfNameProtocolsIsisAreaPasswordPlaintextPassword types.String `tfsdk:"plaintext_password" vyos:"plaintext-password,omitempty"`
	LeafVrfNameProtocolsIsisAreaPasswordMdfive            types.String `tfsdk:"md5" vyos:"md5,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisAreaPassword) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"plaintext_password":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
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

		// Nodes

	}
}
