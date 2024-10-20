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

var _ helpers.VyosResourceDataModel = &PkiOpenTCPPublic{}

// PkiOpenTCPPublic describes the resource data model.
type PkiOpenTCPPublic struct {
	// LeafNodes
	LeafPkiOpenTCPPublicKey  types.String `tfsdk:"key" vyos:"key,omitempty"`
	LeafPkiOpenTCPPublicType types.String `tfsdk:"type" vyos:"type,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PkiOpenTCPPublic) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"key":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Public key in PEM format

`,
			Description: `Public key in PEM format

`,
		},

		"type":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `SSH public key type

    |  Format   |  Description                      |
    |-----------|-----------------------------------|
    |  ssh-rsa  |  Key pair based on RSA algorithm  |
`,
			Description: `SSH public key type

    |  Format   |  Description                      |
    |-----------|-----------------------------------|
    |  ssh-rsa  |  Key pair based on RSA algorithm  |
`,
		},

		// Nodes

	}
}
