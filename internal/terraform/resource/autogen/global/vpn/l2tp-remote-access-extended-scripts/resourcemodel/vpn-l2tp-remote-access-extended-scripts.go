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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (extended-scripts) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &VpnLtwotpRemoteAccessExtendedScrIPts{}

// VpnLtwotpRemoteAccessExtendedScrIPts describes the resource data model.
// This is a basenode!
// Top level basenode type: `Node`
type VpnLtwotpRemoteAccessExtendedScrIPts struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafVpnLtwotpRemoteAccessExtendedScrIPtsOnPreUp  types.String `tfsdk:"on_pre_up" vyos:"on-pre-up,omitempty"`
	LeafVpnLtwotpRemoteAccessExtendedScrIPtsOnUp     types.String `tfsdk:"on_up" vyos:"on-up,omitempty"`
	LeafVpnLtwotpRemoteAccessExtendedScrIPtsOnDown   types.String `tfsdk:"on_down" vyos:"on-down,omitempty"`
	LeafVpnLtwotpRemoteAccessExtendedScrIPtsOnChange types.String `tfsdk:"on_change" vyos:"on-change,omitempty"`

	// TagNodes

	// Nodes
}

// SetID configures the resource ID
func (o *VpnLtwotpRemoteAccessExtendedScrIPts) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *VpnLtwotpRemoteAccessExtendedScrIPts) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *VpnLtwotpRemoteAccessExtendedScrIPts) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnLtwotpRemoteAccessExtendedScrIPts) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"extended-scripts",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *VpnLtwotpRemoteAccessExtendedScrIPts) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (remote-access) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (l2tp) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (vpn) */
		"vpn", // Node

		"l2tp", // Node

		"remote-access", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *VpnLtwotpRemoteAccessExtendedScrIPts) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (remote-access) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (l2tp) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (vpn) */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnLtwotpRemoteAccessExtendedScrIPts) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"on_pre_up":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (on-pre-up) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Script to run before session interface comes up

`,
			Description: `Script to run before session interface comes up

`,
		},

		"on_up":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (on-up) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Script to run when session interface is completely configured and started

`,
			Description: `Script to run when session interface is completely configured and started

`,
		},

		"on_down":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (on-down) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Script to run when session interface going to terminate

`,
			Description: `Script to run when session interface going to terminate

`,
		},

		"on_change":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (on-change) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Script to run when session interface changed by RADIUS CoA handling

`,
			Description: `Script to run when session interface changed by RADIUS CoA handling

`,
		},

		// TagNodes

		// Nodes

	}
}
