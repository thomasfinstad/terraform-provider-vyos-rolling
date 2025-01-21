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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (sync-group) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &HighAvailabilityVrrpSyncGroup{}

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-tag-node-identifier (sync-group) */
// HighAvailabilityVrrpSyncGroupSelfIdentifier used by TagNodes to keep the id info
type HighAvailabilityVrrpSyncGroupSelfIdentifier struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (sync-group) */

	HighAvailabilityVrrpSyncGroup types.String `tfsdk:"sync_group" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (vrrp) */

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (high-availability) */
}

// HighAvailabilityVrrpSyncGroup describes the resource data model.
// This is a basenode!
// Top level basenode type: `TagNode`
type HighAvailabilityVrrpSyncGroup struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	SelfIdentifier *HighAvailabilityVrrpSyncGroupSelfIdentifier `tfsdk:"identifier" vyos:"-,self-id"`

	// LeafNodes
	LeafHighAvailabilityVrrpSyncGroupMember types.List `tfsdk:"member" vyos:"member,omitempty"`

	// TagNodes

	// Nodes

	NodeHighAvailabilityVrrpSyncGroupHealthCheck *HighAvailabilityVrrpSyncGroupHealthCheck `tfsdk:"health_check" vyos:"health-check,omitempty"`

	NodeHighAvailabilityVrrpSyncGroupTransitionScrIPt *HighAvailabilityVrrpSyncGroupTransitionScrIPt `tfsdk:"transition_script" vyos:"transition-script,omitempty"`
}

// SetID configures the resource ID
func (o *HighAvailabilityVrrpSyncGroup) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *HighAvailabilityVrrpSyncGroup) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *HighAvailabilityVrrpSyncGroup) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *HighAvailabilityVrrpSyncGroup) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"sync-group",
		o.SelfIdentifier.HighAvailabilityVrrpSyncGroup.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *HighAvailabilityVrrpSyncGroup) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (vrrp) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (high-availability) */
		"high-availability", // Node

		"vrrp", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *HighAvailabilityVrrpSyncGroup) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (vrrp) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (high-availability) */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o HighAvailabilityVrrpSyncGroup) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"sync_group": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `VRRP sync group

`,
					Description: `VRRP sync group

`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in sync_group, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  sync_group, value must match: ^[.:a-zA-Z0-9-_/]+$",
							),
						),
					},
				},

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (vrrp) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (high-availability) */

			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"member":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (member) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Sync group member

    |  Format  |  Description      |
    |----------|-------------------|
    |  txt     |  VRRP group name  |
`,
			Description: `Sync group member

    |  Format  |  Description      |
    |----------|-------------------|
    |  txt     |  VRRP group name  |
`,
		},

		// TagNodes

		// Nodes

		"health_check": schema.SingleNestedAttribute{
			Attributes: HighAvailabilityVrrpSyncGroupHealthCheck{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Health check

`,
			Description: `Health check

`,
		},

		"transition_script": schema.SingleNestedAttribute{
			Attributes: HighAvailabilityVrrpSyncGroupTransitionScrIPt{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `VRRP transition scripts

`,
			Description: `VRRP transition scripts

`,
		},
	}
}
