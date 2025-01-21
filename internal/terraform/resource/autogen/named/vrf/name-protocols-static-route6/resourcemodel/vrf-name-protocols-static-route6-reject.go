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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (reject) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsStaticRoutesixReject{}

// VrfNameProtocolsStaticRoutesixReject describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VrfNameProtocolsStaticRoutesixReject struct {
	// LeafNodes
	LeafVrfNameProtocolsStaticRoutesixRejectDistance types.Number `tfsdk:"distance" vyos:"distance,omitempty"`
	LeafVrfNameProtocolsStaticRoutesixRejectTag      types.Number `tfsdk:"tag" vyos:"tag,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsStaticRoutesixReject) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"distance":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (distance) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Distance for this route

    |  Format  |  Description              |
    |----------|---------------------------|
    |  1-255   |  Distance for this route  |
`,
			Description: `Distance for this route

    |  Format  |  Description              |
    |----------|---------------------------|
    |  1-255   |  Distance for this route  |
`,
		},

		"tag":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (tag) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Tag value for this route

    |  Format        |  Description               |
    |----------------|----------------------------|
    |  1-4294967295  |  Tag value for this route  |
`,
			Description: `Tag value for this route

    |  Format        |  Description               |
    |----------------|----------------------------|
    |  1-4294967295  |  Tag value for this route  |
`,
		},

		// TagNodes

		// Nodes

	}
}
