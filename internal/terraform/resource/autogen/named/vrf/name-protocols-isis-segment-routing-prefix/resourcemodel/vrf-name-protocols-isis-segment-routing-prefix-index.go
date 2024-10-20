/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsIsisSegmentRoutingPrefixIndex{}

// VrfNameProtocolsIsisSegmentRoutingPrefixIndex describes the resource data model.
type VrfNameProtocolsIsisSegmentRoutingPrefixIndex struct {
	// LeafNodes
	LeafVrfNameProtocolsIsisSegmentRoutingPrefixIndexValue        types.Number `tfsdk:"value" vyos:"value,omitempty"`
	LeafVrfNameProtocolsIsisSegmentRoutingPrefixIndexExplicitNull types.Bool   `tfsdk:"explicit_null" vyos:"explicit-null,omitempty"`
	LeafVrfNameProtocolsIsisSegmentRoutingPrefixIndexNoPhpFlag    types.Bool   `tfsdk:"no_php_flag" vyos:"no-php-flag,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisSegmentRoutingPrefixIndex) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"value":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Specify the index value of prefix segment/label ID

    |  Format   |  Description                       |
    |-----------|------------------------------------|
    |  0-65535  |  The index segment/label ID value  |
`,
			Description: `Specify the index value of prefix segment/label ID

    |  Format   |  Description                       |
    |-----------|------------------------------------|
    |  0-65535  |  The index segment/label ID value  |
`,
		},

		"explicit_null":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Request upstream neighbor to replace segment/label with explicit null label

`,
			Description: `Request upstream neighbor to replace segment/label with explicit null label

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"no_php_flag":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not request penultimate hop popping for segment/label

`,
			Description: `Do not request penultimate hop popping for segment/label

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
