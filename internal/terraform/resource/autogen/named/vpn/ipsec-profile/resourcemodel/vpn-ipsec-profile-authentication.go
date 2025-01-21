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

var _ helpers.VyosResourceDataModel = &VpnIPsecProfileAuthentication{}

// VpnIPsecProfileAuthentication describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VpnIPsecProfileAuthentication struct {
	// LeafNodes
	LeafVpnIPsecProfileAuthenticationMode            types.String `tfsdk:"mode" vyos:"mode,omitempty"`
	LeafVpnIPsecProfileAuthenticationPreSharedSecret types.String `tfsdk:"pre_shared_secret" vyos:"pre-shared-secret,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnIPsecProfileAuthentication) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"mode":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (mode) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Authentication mode

    |  Format             |  Description                  |
    |---------------------|-------------------------------|
    |  pre-shared-secret  |  Use a pre-shared secret key  |
`,
			Description: `Authentication mode

    |  Format             |  Description                  |
    |---------------------|-------------------------------|
    |  pre-shared-secret  |  Use a pre-shared secret key  |
`,
		},

		"pre_shared_secret":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (pre-shared-secret) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Pre-shared secret key

    |  Format  |  Description            |
    |----------|-------------------------|
    |  txt     |  Pre-shared secret key  |
`,
			Description: `Pre-shared secret key

    |  Format  |  Description            |
    |----------|-------------------------|
    |  txt     |  Pre-shared secret key  |
`,
		},

		// TagNodes

		// Nodes

	}
}
