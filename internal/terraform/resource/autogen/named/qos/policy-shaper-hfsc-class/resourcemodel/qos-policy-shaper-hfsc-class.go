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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/numberplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (class) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &QosPolicyShaperHfscClass{}

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-tag-node-identifier (class) */
// QosPolicyShaperHfscClassSelfIdentifier used by TagNodes to keep the id info
type QosPolicyShaperHfscClassSelfIdentifier struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (class) */

	QosPolicyShaperHfscClass types.Number `tfsdk:"class" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (shaper-hfsc) */

	QosPolicyShaperHfsc types.String `tfsdk:"shaper_hfsc" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (policy) */

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (qos) */
}

// QosPolicyShaperHfscClass describes the resource data model.
// This is a basenode!
// Top level basenode type: `TagNode`
type QosPolicyShaperHfscClass struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	SelfIdentifier *QosPolicyShaperHfscClassSelfIdentifier `tfsdk:"identifier" vyos:"-,self-id"`

	// LeafNodes
	LeafQosPolicyShaperHfscClassDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafQosPolicyShaperHfscClassMatchGroup  types.List   `tfsdk:"match_group" vyos:"match-group,omitempty"`

	// TagNodes

	ExistsTagQosPolicyShaperHfscClassMatch bool `tfsdk:"-" vyos:"match,child"`

	// Nodes

	NodeQosPolicyShaperHfscClassLinkshare *QosPolicyShaperHfscClassLinkshare `tfsdk:"linkshare" vyos:"linkshare,omitempty"`

	NodeQosPolicyShaperHfscClassRealtime *QosPolicyShaperHfscClassRealtime `tfsdk:"realtime" vyos:"realtime,omitempty"`

	NodeQosPolicyShaperHfscClassUpperlimit *QosPolicyShaperHfscClassUpperlimit `tfsdk:"upperlimit" vyos:"upperlimit,omitempty"`
}

// SetID configures the resource ID
func (o *QosPolicyShaperHfscClass) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *QosPolicyShaperHfscClass) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *QosPolicyShaperHfscClass) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *QosPolicyShaperHfscClass) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"class",
		o.SelfIdentifier.QosPolicyShaperHfscClass.ValueBigFloat().String(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *QosPolicyShaperHfscClass) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (shaper-hfsc) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (policy) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (qos) */
		"qos", // Node

		"policy", // Node

		"shaper-hfsc",
		o.SelfIdentifier.QosPolicyShaperHfsc.ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *QosPolicyShaperHfscClass) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (shaper-hfsc) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (shaper-hfsc) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (policy) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (qos) */
		"qos", // Node

		"policy", // Node

		"shaper-hfsc",
		o.SelfIdentifier.QosPolicyShaperHfsc.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyShaperHfscClass) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"class": schema.NumberAttribute{
					Required: true,
					MarkdownDescription: `Class ID

    |  Format  |  Description       |
    |----------|--------------------|
    |  1-4095  |  Class Identifier  |
`,
					Description: `Class ID

    |  Format  |  Description       |
    |----------|--------------------|
    |  1-4095  |  Class Identifier  |
`,
					PlanModifiers: []planmodifier.Number{
						numberplanmodifier.RequiresReplace(),
					},
				},

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (shaper-hfsc) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (policy) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (qos) */

				"shaper_hfsc": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Hierarchical Fair Service Curve's policy

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Policy name  |
`,
					Description: `Hierarchical Fair Service Curve's policy

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Policy name  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in shaper_hfsc, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  shaper_hfsc, value must match: ^[.:a-zA-Z0-9-_/]+$",
							),
						),
					},
				},
			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

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

		"match_group":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (match-group) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Filter group for QoS policy

    |  Format  |  Description       |
    |----------|--------------------|
    |  txt     |  Match group name  |
`,
			Description: `Filter group for QoS policy

    |  Format  |  Description       |
    |----------|--------------------|
    |  txt     |  Match group name  |
`,
		},

		// TagNodes

		// Nodes

		"linkshare": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperHfscClassLinkshare{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Linkshare class settings

`,
			Description: `Linkshare class settings

`,
		},

		"realtime": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperHfscClassRealtime{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Realtime class settings

`,
			Description: `Realtime class settings

`,
		},

		"upperlimit": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperHfscClassUpperlimit{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Upperlimit class settings

`,
			Description: `Upperlimit class settings

`,
		},
	}
}
