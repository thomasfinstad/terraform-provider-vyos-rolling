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

var _ helpers.VyosTopResourceDataModel = &NatDestinationRuleLoadBalanceBackend{}

// NatDestinationRuleLoadBalanceBackend describes the resource data model.
type NatDestinationRuleLoadBalanceBackend struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafNatDestinationRuleLoadBalanceBackendWeight types.Number `tfsdk:"weight" vyos:"weight,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// SetID configures the resource ID
func (o *NatDestinationRuleLoadBalanceBackend) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *NatDestinationRuleLoadBalanceBackend) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *NatDestinationRuleLoadBalanceBackend) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *NatDestinationRuleLoadBalanceBackend) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"backend",
		o.SelfIdentifier.Attributes()["backend"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *NatDestinationRuleLoadBalanceBackend) GetVyosParentPath() []string {
	return []string{
		"nat",

		"destination",

		"rule",

		o.SelfIdentifier.Attributes()["rule"].(types.Number).ValueBigFloat().String(),

		"load-balance",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *NatDestinationRuleLoadBalanceBackend) GetVyosNamedParentPath() []string {
	return []string{
		"nat",

		"destination",

		"rule",

		o.SelfIdentifier.Attributes()["rule"].(types.Number).ValueBigFloat().String(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o NatDestinationRuleLoadBalanceBackend) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.MapNestedAttribute{
			Required: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"backend": schema.StringAttribute{
						Required: true,
						MarkdownDescription: `Translated IP address

    |  Format  |  Description            |
    |----------|-------------------------|
    |  ipv4    |  IPv4 address to match  |
`,
						Description: `Translated IP address

    |  Format  |  Description            |
    |----------|-------------------------|
    |  ipv4    |  IPv4 address to match  |
`,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						}, Validators: []validator.String{
							stringvalidator.All(
								helpers.StringNot(
									stringvalidator.RegexMatches(
										regexp.MustCompile(`^.*__.*$`),
										"double underscores in backend, conflicts with the internal resource id",
									),
								),
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
									"illegal character in  backend, value must match: ^[a-zA-Z0-9-_]*$",
								),
							),
						},
					},

					"rule": schema.NumberAttribute{
						Required: true,
						MarkdownDescription: `Rule number for NAT

    |  Format    |  Description         |
    |------------|----------------------|
    |  1-999999  |  Number of NAT rule  |
`,
						Description: `Rule number for NAT

    |  Format    |  Description         |
    |------------|----------------------|
    |  1-999999  |  Number of NAT rule  |
`,
						PlanModifiers: []planmodifier.Number{
							numberplanmodifier.RequiresReplace(),
						},
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
			MarkdownDescription: `Set probability for this output value

    |  Format  |  Description                            |
    |----------|-----------------------------------------|
    |  1-100   |  Set probability for this output value  |
`,
			Description: `Set probability for this output value

    |  Format  |  Description                            |
    |----------|-----------------------------------------|
    |  1-100   |  Set probability for this output value  |
`,
		},

		// Nodes

	}
}
