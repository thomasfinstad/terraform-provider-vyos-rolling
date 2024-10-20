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
var _ helpers.VyosTopResourceDataModel = &ServiceIPoeServerAuthenticationRadiusDynamicAuthor{}

// ServiceIPoeServerAuthenticationRadiusDynamicAuthor describes the resource data model.
type ServiceIPoeServerAuthenticationRadiusDynamicAuthor struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceIPoeServerAuthenticationRadiusDynamicAuthorServer types.String `tfsdk:"server" vyos:"server,omitempty"`
	LeafServiceIPoeServerAuthenticationRadiusDynamicAuthorPort   types.Number `tfsdk:"port" vyos:"port,omitempty"`
	LeafServiceIPoeServerAuthenticationRadiusDynamicAuthorKey    types.String `tfsdk:"key" vyos:"key,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *ServiceIPoeServerAuthenticationRadiusDynamicAuthor) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceIPoeServerAuthenticationRadiusDynamicAuthor) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceIPoeServerAuthenticationRadiusDynamicAuthor) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceIPoeServerAuthenticationRadiusDynamicAuthor) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"dynamic-author",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceIPoeServerAuthenticationRadiusDynamicAuthor) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/global/resource-model-parent-vyos-path-hack.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/global/resource-model-parent-vyos-path-hack.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/global/resource-model-parent-vyos-path-hack.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/global/resource-model-parent-vyos-path-hack.gotmpl */
		"service",

		"ipoe-server",

		"authentication",

		"radius",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *ServiceIPoeServerAuthenticationRadiusDynamicAuthor) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack-for-non-global.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack-for-non-global.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack-for-non-global.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack-for-non-global.gotmpl */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceIPoeServerAuthenticationRadiusDynamicAuthor) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"server":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address for Dynamic Authorization Extension server (DM/CoA)

    |  Format  |  Description                                    |
    |----------|-------------------------------------------------|
    |  ipv4    |  IPv4 address for dynamic authorization server  |
`,
			Description: `IP address for Dynamic Authorization Extension server (DM/CoA)

    |  Format  |  Description                                    |
    |----------|-------------------------------------------------|
    |  ipv4    |  IPv4 address for dynamic authorization server  |
`,
		},

		"port":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Port for Dynamic Authorization Extension server (DM/CoA)

    |  Format   |  Description  |
    |-----------|---------------|
    |  1-65535  |  TCP port     |
`,
			Description: `Port for Dynamic Authorization Extension server (DM/CoA)

    |  Format   |  Description  |
    |-----------|---------------|
    |  1-65535  |  TCP port     |
`,

			// Default:          stringdefault.StaticString(`1700`),
			Computed: true,
		},

		"key":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Shared secret for Dynamic Authorization Extension server

`,
			Description: `Shared secret for Dynamic Authorization Extension server

`,
		},
	}
}
