/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (login) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &SystemLogin{}

// SystemLogin describes the resource data model.
// This is a basenode!
// Top level basenode type: `Node`
type SystemLogin struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafSystemLoginMaxLoginSession types.Number `tfsdk:"max_login_session" vyos:"max-login-session,omitempty"`
	LeafSystemLoginTimeout         types.Number `tfsdk:"timeout" vyos:"timeout,omitempty"`

	// TagNodes

	ExistsTagSystemLoginUser bool `tfsdk:"-" vyos:"user,child"`

	// Nodes

	ExistsNodeSystemLoginRadius bool `tfsdk:"-" vyos:"radius,child"`

	ExistsNodeSystemLoginTacacs bool `tfsdk:"-" vyos:"tacacs,child"`

	ExistsNodeSystemLoginBanner bool `tfsdk:"-" vyos:"banner,child"`
}

// SetID configures the resource ID
func (o *SystemLogin) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *SystemLogin) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *SystemLogin) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *SystemLogin) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"login",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *SystemLogin) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (system) */
		"system", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *SystemLogin) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (system) */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemLogin) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"max_login_session":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (max-login-session) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum number of all login sessions

    |  Format   |  Description                           |
    |-----------|----------------------------------------|
    |  1-65536  |  Maximum number of all login sessions  |
`,
			Description: `Maximum number of all login sessions

    |  Format   |  Description                           |
    |-----------|----------------------------------------|
    |  1-65536  |  Maximum number of all login sessions  |
`,
		},

		"timeout":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (timeout) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Session timeout

    |  Format    |  Description                 |
    |------------|------------------------------|
    |  5-604800  |  Session timeout in seconds  |
`,
			Description: `Session timeout

    |  Format    |  Description                 |
    |------------|------------------------------|
    |  5-604800  |  Session timeout in seconds  |
`,
		},

		// TagNodes

		// Nodes

	}
}
