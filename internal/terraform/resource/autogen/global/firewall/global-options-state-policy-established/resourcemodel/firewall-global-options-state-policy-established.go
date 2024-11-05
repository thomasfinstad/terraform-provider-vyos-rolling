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

var _ helpers.VyosTopResourceDataModel = &FirewallGlobalOptionsStatePolicyEstablished{}

// FirewallGlobalOptionsStatePolicyEstablished describes the resource data model.
// This is a basenode!
// Top level basenode type: `Node`
type FirewallGlobalOptionsStatePolicyEstablished struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafFirewallGlobalOptionsStatePolicyEstablishedAction   types.String `tfsdk:"action" vyos:"action,omitempty"`
	LeafFirewallGlobalOptionsStatePolicyEstablishedLog      types.Bool   `tfsdk:"log" vyos:"log,omitempty"`
	LeafFirewallGlobalOptionsStatePolicyEstablishedLogLevel types.String `tfsdk:"log_level" vyos:"log-level,omitempty"`

	// TagNodes

	// Nodes
}

// SetID configures the resource ID
func (o *FirewallGlobalOptionsStatePolicyEstablished) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *FirewallGlobalOptionsStatePolicyEstablished) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *FirewallGlobalOptionsStatePolicyEstablished) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *FirewallGlobalOptionsStatePolicyEstablished) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"established",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *FirewallGlobalOptionsStatePolicyEstablished) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */
		"firewall", // Node

		"global-options", // Node

		"state-policy", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *FirewallGlobalOptionsStatePolicyEstablished) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallGlobalOptionsStatePolicyEstablished) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Action for packets

    |  Format  |  Description       |
    |----------|--------------------|
    |  accept  |  Action to accept  |
    |  drop    |  Action to drop    |
    |  reject  |  Action to reject  |
`,
			Description: `Action for packets

    |  Format  |  Description       |
    |----------|--------------------|
    |  accept  |  Action to accept  |
    |  drop    |  Action to drop    |
    |  reject  |  Action to reject  |
`,
		},

		"log":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Log packets hitting this rule

`,
			Description: `Log packets hitting this rule

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"log_level":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set log-level. Log must be enable.

    |  Format  |  Description         |
    |----------|----------------------|
    |  emerg   |  Emerg log level     |
    |  alert   |  Alert log level     |
    |  crit    |  Critical log level  |
    |  err     |  Error log level     |
    |  warn    |  Warning log level   |
    |  notice  |  Notice log level    |
    |  info    |  Info log level      |
    |  debug   |  Debug log level     |
`,
			Description: `Set log-level. Log must be enable.

    |  Format  |  Description         |
    |----------|----------------------|
    |  emerg   |  Emerg log level     |
    |  alert   |  Alert log level     |
    |  crit    |  Critical log level  |
    |  err     |  Error log level     |
    |  warn    |  Warning log level   |
    |  notice  |  Notice log level    |
    |  info    |  Info log level      |
    |  debug   |  Debug log level     |
`,
		},

		// TagNodes

		// Nodes

	}
}
