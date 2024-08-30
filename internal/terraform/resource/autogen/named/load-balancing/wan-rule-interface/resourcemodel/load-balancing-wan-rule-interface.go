// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

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

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &LoadBalancingWanRuleInterface{}

// LoadBalancingWanRuleInterface describes the resource data model.
type LoadBalancingWanRuleInterface struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafLoadBalancingWanRuleInterfaceWeight types.Number `tfsdk:"weight" vyos:"weight,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// SetID configures the resource ID
func (o *LoadBalancingWanRuleInterface) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *LoadBalancingWanRuleInterface) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *LoadBalancingWanRuleInterface) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *LoadBalancingWanRuleInterface) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"interface",
		o.SelfIdentifier.Attributes()["interface"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *LoadBalancingWanRuleInterface) GetVyosParentPath() []string {
	return []string{
		"load-balancing",

		"wan",

		"rule",

		o.SelfIdentifier.Attributes()["rule"].(types.Number).ValueBigFloat().String(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *LoadBalancingWanRuleInterface) GetVyosNamedParentPath() []string {
	return []string{
		"load-balancing",

		"wan",

		"rule",

		o.SelfIdentifier.Attributes()["rule"].(types.Number).ValueBigFloat().String(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o LoadBalancingWanRuleInterface) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"interface": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Interface name [REQUIRED]

`,
					Description: `Interface name [REQUIRED]

`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in interface, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
								"illegal character in  interface, value must match: ^[a-zA-Z0-9-_]*$",
							),
						),
					},
				},

				"rule": schema.NumberAttribute{
					Required: true,
					MarkdownDescription: `Rule number (1-9999)

    |  Format  |  Description  |
    |----------|---------------|
    |  1-9999  |  Rule number  |
`,
					Description: `Rule number (1-9999)

    |  Format  |  Description  |
    |----------|---------------|
    |  1-9999  |  Rule number  |
`,
					PlanModifiers: []planmodifier.Number{
						numberplanmodifier.RequiresReplace(),
					},
				},
			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"weight": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Load-balance weight

    |  Format  |  Description       |
    |----------|--------------------|
    |  1-255   |  Interface weight  |
`,
			Description: `Load-balance weight

    |  Format  |  Description       |
    |----------|--------------------|
    |  1-255   |  Interface weight  |
`,

			// Default:          stringdefault.StaticString(`1`),
			Computed: true,
		},

		// Nodes

	}
}
