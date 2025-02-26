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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (authentication) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsOspfAreaVirtualLinkAuthentication{}

// VrfNameProtocolsOspfAreaVirtualLinkAuthentication describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VrfNameProtocolsOspfAreaVirtualLinkAuthentication struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfAreaVirtualLinkAuthenticationPlaintextPassword types.String `tfsdk:"plaintext_password" vyos:"plaintext-password,omitempty"`

	// TagNodes

	// Nodes

	// Ignoring Node `VrfNameProtocolsOspfAreaVirtualLinkAuthenticationMdfive`.
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfAreaVirtualLinkAuthentication) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"plaintext_password":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (plaintext-password) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Plain text password

    |  Format  |  Description                                 |
    |----------|----------------------------------------------|
    |  txt     |  Plain text password (8 characters or less)  |
`,
			Description: `Plain text password

    |  Format  |  Description                                 |
    |----------|----------------------------------------------|
    |  txt     |  Plain text password (8 characters or less)  |
`,
		},

		// TagNodes

		// Nodes

	}
}
