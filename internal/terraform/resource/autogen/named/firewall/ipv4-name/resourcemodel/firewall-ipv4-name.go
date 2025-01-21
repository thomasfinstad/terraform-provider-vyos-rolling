/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (name) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &FirewallIPvfourName{}

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-tag-node-identifier (name) */
// FirewallIPvfourNameSelfIdentifier used by TagNodes to keep the id info
type FirewallIPvfourNameSelfIdentifier struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (name) */

	FirewallIPvfourName types.String `tfsdk:"name" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (ipv4) */

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (firewall) */
}

// FirewallIPvfourName describes the resource data model.
// This is a basenode!
// Top level basenode type: `TagNode`
type FirewallIPvfourName struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	SelfIdentifier *FirewallIPvfourNameSelfIdentifier `tfsdk:"identifier" vyos:"-,self-id"`

	// LeafNodes
	LeafFirewallIPvfourNameDefaultAction     types.String `tfsdk:"default_action" vyos:"default-action,omitempty"`
	LeafFirewallIPvfourNameDefaultLog        types.Bool   `tfsdk:"default_log" vyos:"default-log,omitempty"`
	LeafFirewallIPvfourNameDescrIPtion       types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafFirewallIPvfourNameDefaultJumpTarget types.String `tfsdk:"default_jump_target" vyos:"default-jump-target,omitempty"`

	// TagNodes

	ExistsTagFirewallIPvfourNameRule bool `tfsdk:"-" vyos:"rule,child"`

	// Nodes
}

// SetID configures the resource ID
func (o *FirewallIPvfourName) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *FirewallIPvfourName) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *FirewallIPvfourName) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *FirewallIPvfourName) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"name",
		o.SelfIdentifier.FirewallIPvfourName.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *FirewallIPvfourName) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (ipv4) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (firewall) */
		"firewall", // Node

		"ipv4", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *FirewallIPvfourName) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (ipv4) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (firewall) */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvfourName) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"name": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `IPv4 custom firewall

`,
					Description: `IPv4 custom firewall

`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in name, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  name, value must match: ^[.:a-zA-Z0-9-_/]+$",
							),
						),
					},
				},

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (ipv4) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (firewall) */

			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"default_action":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (default-action) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Default-action for rule-set

    |  Format    |  Description                                                                    |
    |------------|---------------------------------------------------------------------------------|
    |  drop      |  Drop if no prior rules are hit                                                 |
    |  jump      |  Jump to another chain if no prior rules are hit                                |
    |  reject    |  Drop and notify source if no prior rules are hit                               |
    |  return    |  Return from the current chain and continue at the next rule of the last chain  |
    |  accept    |  Accept if no prior rules are hit                                               |
    |  continue  |  Continue parsing next rule                                                     |
`,
			Description: `Default-action for rule-set

    |  Format    |  Description                                                                    |
    |------------|---------------------------------------------------------------------------------|
    |  drop      |  Drop if no prior rules are hit                                                 |
    |  jump      |  Jump to another chain if no prior rules are hit                                |
    |  reject    |  Drop and notify source if no prior rules are hit                               |
    |  return    |  Return from the current chain and continue at the next rule of the last chain  |
    |  accept    |  Accept if no prior rules are hit                                               |
    |  continue  |  Continue parsing next rule                                                     |
`,

			// Default:          stringdefault.StaticString(`drop`),
			Computed: true,
		},

		"default_log":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (default-log) */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (description) */
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

		"default_jump_target":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (default-jump-target) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set jump target. Action jump must be defined in default-action to use this setting

`,
			Description: `Set jump target. Action jump must be defined in default-action to use this setting

`,
		},

		// TagNodes

		// Nodes

	}
}
