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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (public) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &PkiKeyPairPublic{}

// PkiKeyPairPublic describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type PkiKeyPairPublic struct {
	// LeafNodes
	LeafPkiKeyPairPublicKey types.String `tfsdk:"key" vyos:"key,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PkiKeyPairPublic) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"key":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (key) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Public key in PEM format

`,
			Description: `Public key in PEM format

`,
		},

		// TagNodes

		// Nodes

	}
}
