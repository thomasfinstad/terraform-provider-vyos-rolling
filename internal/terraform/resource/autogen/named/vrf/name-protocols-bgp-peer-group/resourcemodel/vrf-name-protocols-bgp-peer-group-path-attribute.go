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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpPeerGroupPathAttribute{}

// VrfNameProtocolsBgpPeerGroupPathAttribute describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VrfNameProtocolsBgpPeerGroupPathAttribute struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpPeerGroupPathAttributeDiscard         types.List   `tfsdk:"discard" vyos:"discard,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupPathAttributeTreatAsWithdraw types.Number `tfsdk:"treat_as_withdraw" vyos:"treat-as-withdraw,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpPeerGroupPathAttribute) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"discard":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
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

		// TagNodes

		// Nodes

	}
}
