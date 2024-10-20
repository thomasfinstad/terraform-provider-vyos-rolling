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

var _ helpers.VyosResourceDataModel = &InterfacesOpenvpnAuthentication{}

// InterfacesOpenvpnAuthentication describes the resource data model.
type InterfacesOpenvpnAuthentication struct {
	// LeafNodes
	LeafInterfacesOpenvpnAuthenticationUsername types.String `tfsdk:"username" vyos:"username,omitempty"`
	LeafInterfacesOpenvpnAuthenticationPassword types.String `tfsdk:"password" vyos:"password,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesOpenvpnAuthentication) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"username":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Username used for authentication

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Username     |
`,
			Description: `Username used for authentication

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Username     |
`,
		},

		"password":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Password used for authentication

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Password     |
`,
			Description: `Password used for authentication

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Password     |
`,
		},

		// Nodes

	}
}
