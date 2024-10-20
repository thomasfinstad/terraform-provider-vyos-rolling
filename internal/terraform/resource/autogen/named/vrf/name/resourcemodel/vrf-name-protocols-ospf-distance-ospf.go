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

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsOspfDistanceOspf{}

// VrfNameProtocolsOspfDistanceOspf describes the resource data model.
type VrfNameProtocolsOspfDistanceOspf struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfDistanceOspfExternal  types.Number `tfsdk:"external" vyos:"external,omitempty"`
	LeafVrfNameProtocolsOspfDistanceOspfInterArea types.Number `tfsdk:"inter_area" vyos:"inter-area,omitempty"`
	LeafVrfNameProtocolsOspfDistanceOspfIntraArea types.Number `tfsdk:"intra_area" vyos:"intra-area,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfDistanceOspf) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"external":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Distance for external routes

    |  Format  |  Description                   |
    |----------|--------------------------------|
    |  1-255   |  Distance for external routes  |
`,
			Description: `Distance for external routes

    |  Format  |  Description                   |
    |----------|--------------------------------|
    |  1-255   |  Distance for external routes  |
`,
		},

		"inter_area":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Distance for inter-area routes

    |  Format  |  Description                     |
    |----------|----------------------------------|
    |  1-255   |  Distance for inter-area routes  |
`,
			Description: `Distance for inter-area routes

    |  Format  |  Description                     |
    |----------|----------------------------------|
    |  1-255   |  Distance for inter-area routes  |
`,
		},

		"intra_area":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Distance for intra-area routes

    |  Format  |  Description                     |
    |----------|----------------------------------|
    |  1-255   |  Distance for intra-area routes  |
`,
			Description: `Distance for intra-area routes

    |  Format  |  Description                     |
    |----------|----------------------------------|
    |  1-255   |  Distance for intra-area routes  |
`,
		},

		// Nodes

	}
}
