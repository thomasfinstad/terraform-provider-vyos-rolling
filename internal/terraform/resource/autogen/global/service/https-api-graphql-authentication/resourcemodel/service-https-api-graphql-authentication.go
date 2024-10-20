/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/global/resource-model.gotmpl */
// Validate compliance
var _ helpers.VyosTopResourceDataModel = &ServiceHTTPSAPIGraphqlAuthentication{}

// ServiceHTTPSAPIGraphqlAuthentication describes the resource data model.
type ServiceHTTPSAPIGraphqlAuthentication struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceHTTPSAPIGraphqlAuthenticationType         types.String `tfsdk:"type" vyos:"type,omitempty"`
	LeafServiceHTTPSAPIGraphqlAuthenticationExpiration   types.Number `tfsdk:"expiration" vyos:"expiration,omitempty"`
	LeafServiceHTTPSAPIGraphqlAuthenticationSecretLength types.Number `tfsdk:"secret_length" vyos:"secret-length,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *ServiceHTTPSAPIGraphqlAuthentication) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceHTTPSAPIGraphqlAuthentication) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceHTTPSAPIGraphqlAuthentication) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceHTTPSAPIGraphqlAuthentication) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"authentication",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceHTTPSAPIGraphqlAuthentication) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/global/resource-model-parent-vyos-path-hack.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/global/resource-model-parent-vyos-path-hack.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/global/resource-model-parent-vyos-path-hack.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/global/resource-model-parent-vyos-path-hack.gotmpl */
		"service",

		"https",

		"api",

		"graphql",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *ServiceHTTPSAPIGraphqlAuthentication) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack-for-non-global.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack-for-non-global.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack-for-non-global.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack-for-non-global.gotmpl */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceHTTPSAPIGraphqlAuthentication) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"type":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Authentication type

    |  Format  |  Description    |
    |----------|-----------------|
    |  key     |  Use API keys   |
    |  token   |  Use JWT token  |
`,
			Description: `Authentication type

    |  Format  |  Description    |
    |----------|-----------------|
    |  key     |  Use API keys   |
    |  token   |  Use JWT token  |
`,

			// Default:          stringdefault.StaticString(`key`),
			Computed: true,
		},

		"expiration":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Token time to expire in seconds

    |  Format       |  Description                |
    |---------------|-----------------------------|
    |  60-31536000  |  Token lifetime in seconds  |
`,
			Description: `Token time to expire in seconds

    |  Format       |  Description                |
    |---------------|-----------------------------|
    |  60-31536000  |  Token lifetime in seconds  |
`,

			// Default:          stringdefault.StaticString(`3600`),
			Computed: true,
		},

		"secret_length":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Length of shared secret in bytes

    |  Format    |  Description                             |
    |------------|------------------------------------------|
    |  16-65535  |  Byte length of generated shared secret  |
`,
			Description: `Length of shared secret in bytes

    |  Format    |  Description                             |
    |------------|------------------------------------------|
    |  16-65535  |  Byte length of generated shared secret  |
`,

			// Default:          stringdefault.StaticString(`32`),
			Computed: true,
		},
	}
}
