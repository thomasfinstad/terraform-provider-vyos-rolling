// Package resourcemodel code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

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

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &NatStaticRule{}

// NatStaticRule describes the resource data model.
type NatStaticRule struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"rule_id" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafNatStaticRuleDescrIPtion      types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafNatStaticRuleInboundInterface types.String `tfsdk:"inbound_interface" vyos:"inbound-interface,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeNatStaticRuleDestination *NatStaticRuleDestination `tfsdk:"destination" vyos:"destination,omitempty"`
	NodeNatStaticRuleTranSLAtion *NatStaticRuleTranSLAtion `tfsdk:"translation" vyos:"translation,omitempty"`
}

// SetID configures the resource ID
func (o *NatStaticRule) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *NatStaticRule) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *NatStaticRule) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *NatStaticRule) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"rule",
		o.SelfIdentifier.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *NatStaticRule) GetVyosParentPath() []string {
	return []string{
		"nat",

		"static",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *NatStaticRule) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o NatStaticRule) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"rule_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Rule number for NAT

`,
			Description: `Rule number for NAT

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			}, Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in rule_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illegal character in  rule_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"description": schema.StringAttribute{
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

		"inbound_interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Inbound interface of NAT traffic

`,
			Description: `Inbound interface of NAT traffic

`,
		},

		// Nodes

		"destination": schema.SingleNestedAttribute{
			Attributes: NatStaticRuleDestination{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `NAT destination parameters

`,
			Description: `NAT destination parameters

`,
		},

		"translation": schema.SingleNestedAttribute{
			Attributes: NatStaticRuleTranSLAtion{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Translation address or prefix

`,
			Description: `Translation address or prefix

`,
		},
	}
}
