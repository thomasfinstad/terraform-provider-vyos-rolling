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

var _ helpers.VyosResourceDataModel = &ProtocolsBgpNeighborPathAttribute{}

// ProtocolsBgpNeighborPathAttribute describes the resource data model.
type ProtocolsBgpNeighborPathAttribute struct {
	// LeafNodes
	LeafProtocolsBgpNeighborPathAttributeDiscard         types.List   `tfsdk:"discard" vyos:"discard,omitempty"`
	LeafProtocolsBgpNeighborPathAttributeTreatAsWithdraw types.Number `tfsdk:"treat_as_withdraw" vyos:"treat-as-withdraw,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborPathAttribute) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"discard":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype-multi.gotmpl */
		schema.ListAttribute{
			ElementType: types.NumberType,
			Optional:    true,
			MarkdownDescription: `Drop specified attributes from incoming UPDATE messages

    |  Format  |  Description       |
    |----------|--------------------|
    |  1-255   |  Attribute number  |
`,
			Description: `Drop specified attributes from incoming UPDATE messages

    |  Format  |  Description       |
    |----------|--------------------|
    |  1-255   |  Attribute number  |
`,
		},

		"treat_as_withdraw":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Treat-as-withdraw any incoming BGP UPDATE messages that contain the specified attribute

    |  Format  |  Description       |
    |----------|--------------------|
    |  1-255   |  Attribute number  |
`,
			Description: `Treat-as-withdraw any incoming BGP UPDATE messages that contain the specified attribute

    |  Format  |  Description       |
    |----------|--------------------|
    |  1-255   |  Attribute number  |
`,
		},

		// Nodes

	}
}
