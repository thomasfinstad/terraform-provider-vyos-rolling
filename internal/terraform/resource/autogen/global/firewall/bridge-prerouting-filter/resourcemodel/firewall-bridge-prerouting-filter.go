/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &FirewallBrIDgePreroutingFilter{}

// FirewallBrIDgePreroutingFilter describes the resource data model.
// This is a basenode!
// Top level basenode type: `Node`
type FirewallBrIDgePreroutingFilter struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafFirewallBrIDgePreroutingFilterDefaultAction types.String `tfsdk:"default_action" vyos:"default-action,omitempty"`
	LeafFirewallBrIDgePreroutingFilterDefaultLog    types.Bool   `tfsdk:"default_log" vyos:"default-log,omitempty"`
	LeafFirewallBrIDgePreroutingFilterDescrIPtion   types.String `tfsdk:"description" vyos:"description,omitempty"`

	// TagNodes

	ExistsTagFirewallBrIDgePreroutingFilterRule bool `tfsdk:"-" vyos:"rule,child"`

	// Nodes
}

// SetID configures the resource ID
func (o *FirewallBrIDgePreroutingFilter) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *FirewallBrIDgePreroutingFilter) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *FirewallBrIDgePreroutingFilter) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *FirewallBrIDgePreroutingFilter) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"filter",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *FirewallBrIDgePreroutingFilter) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */
		"firewall", // Node

		"bridge", // Node

		"prerouting", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *FirewallBrIDgePreroutingFilter) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallBrIDgePreroutingFilter) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"default_action":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Default-action for rule-set

    |  Format  |  Description                       |
    |----------|------------------------------------|
    |  drop    |  Drop if no prior rules are hit    |
    |  accept  |  Accept if no prior rules are hit  |
`,
			Description: `Default-action for rule-set

    |  Format  |  Description                       |
    |----------|------------------------------------|
    |  drop    |  Drop if no prior rules are hit    |
    |  accept  |  Accept if no prior rules are hit  |
`,

			// Default:          stringdefault.StaticString(`accept`),
			Computed: true,
		},

		"default_log":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Log packets hitting default-action

`,
			Description: `Log packets hitting default-action

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"description":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Description  |
`,
			Description: `Description

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Description  |
`,
		},

		// TagNodes

		// Nodes

	}
}
