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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (commit-confirm) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &SystemConfigManagementCommitConfirm{}

// SystemConfigManagementCommitConfirm describes the resource data model.
// This is a basenode!
// Top level basenode type: `Node`
type SystemConfigManagementCommitConfirm struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafSystemConfigManagementCommitConfirmAction types.String `tfsdk:"action" vyos:"action,omitempty"`

	// TagNodes

	// Nodes
}

// SetID configures the resource ID
func (o *SystemConfigManagementCommitConfirm) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *SystemConfigManagementCommitConfirm) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *SystemConfigManagementCommitConfirm) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *SystemConfigManagementCommitConfirm) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"commit-confirm",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *SystemConfigManagementCommitConfirm) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (config-management) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (system) */
		"system", // Node

		"config-management", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *SystemConfigManagementCommitConfirm) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (config-management) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (system) */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemConfigManagementCommitConfirm) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"action":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (action) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Commit confirm revert action

    |  Format  |  Description                                     |
    |----------|--------------------------------------------------|
    |  reload  |  Reload previous configuration if not confirmed  |
    |  reboot  |  Reboot to saved configuration if not confirmed  |
`,
			Description: `Commit confirm revert action

    |  Format  |  Description                                     |
    |----------|--------------------------------------------------|
    |  reload  |  Reload previous configuration if not confirmed  |
    |  reboot  |  Reboot to saved configuration if not confirmed  |
`,

			// Default:          stringdefault.StaticString(`reboot`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
