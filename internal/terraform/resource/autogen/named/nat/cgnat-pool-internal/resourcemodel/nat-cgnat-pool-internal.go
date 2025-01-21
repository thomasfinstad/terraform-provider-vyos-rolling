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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (internal) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &NatCgnatPoolInternal{}

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-tag-node-identifier (internal) */
// NatCgnatPoolInternalSelfIdentifier used by TagNodes to keep the id info
type NatCgnatPoolInternalSelfIdentifier struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (internal) */

	NatCgnatPoolInternal types.String `tfsdk:"internal" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (pool) */

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (cgnat) */

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (nat) */
}

// NatCgnatPoolInternal describes the resource data model.
// This is a basenode!
// Top level basenode type: `TagNode`
type NatCgnatPoolInternal struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	SelfIdentifier *NatCgnatPoolInternalSelfIdentifier `tfsdk:"identifier" vyos:"-,self-id"`

	// LeafNodes
	LeafNatCgnatPoolInternalRange types.List `tfsdk:"range" vyos:"range,omitempty"`

	// TagNodes

	// Nodes
}

// SetID configures the resource ID
func (o *NatCgnatPoolInternal) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *NatCgnatPoolInternal) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *NatCgnatPoolInternal) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *NatCgnatPoolInternal) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"internal",
		o.SelfIdentifier.NatCgnatPoolInternal.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *NatCgnatPoolInternal) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (pool) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (cgnat) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (nat) */
		"nat", // Node

		"cgnat", // Node

		"pool", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *NatCgnatPoolInternal) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (pool) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (cgnat) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (nat) */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o NatCgnatPoolInternal) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"internal": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Internal pool name

    |  Format  |  Description         |
    |----------|----------------------|
    |  txt     |  Internal pool name  |
`,
					Description: `Internal pool name

    |  Format  |  Description         |
    |----------|----------------------|
    |  txt     |  Internal pool name  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in internal, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  internal, value must match: ^[.:a-zA-Z0-9-_/]+$",
							),
						),
					},
				},

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (pool) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (cgnat) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (nat) */

			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"range":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (range) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Range of IP addresses

    |  Format     |  Description         |
    |-------------|----------------------|
    |  ipv4net    |  IPv4 prefix         |
    |  ipv4range  |  IPv4 address range  |
`,
			Description: `Range of IP addresses

    |  Format     |  Description         |
    |-------------|----------------------|
    |  ipv4net    |  IPv4 prefix         |
    |  ipv4range  |  IPv4 address range  |
`,
		},

		// TagNodes

		// Nodes

	}
}
