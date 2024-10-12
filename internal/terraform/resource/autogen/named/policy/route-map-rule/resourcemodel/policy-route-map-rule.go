// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
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

var _ helpers.VyosTopResourceDataModel = &PolicyRouteMapRule{}

// PolicyRouteMapRule describes the resource data model.
type PolicyRouteMapRule struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafPolicyRouteMapRuleAction      types.String `tfsdk:"action" vyos:"action,omitempty"`
	LeafPolicyRouteMapRuleCall        types.String `tfsdk:"call" vyos:"call,omitempty"`
	LeafPolicyRouteMapRuleContinue    types.Number `tfsdk:"continue" vyos:"continue,omitempty"`
	LeafPolicyRouteMapRuleDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodePolicyRouteMapRuleMatch   *PolicyRouteMapRuleMatch   `tfsdk:"match" vyos:"match,omitempty"`
	NodePolicyRouteMapRuleOnMatch *PolicyRouteMapRuleOnMatch `tfsdk:"on_match" vyos:"on-match,omitempty"`
	NodePolicyRouteMapRuleSet     *PolicyRouteMapRuleSet     `tfsdk:"set" vyos:"set,omitempty"`
}

// SetID configures the resource ID
func (o *PolicyRouteMapRule) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *PolicyRouteMapRule) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *PolicyRouteMapRule) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *PolicyRouteMapRule) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"rule",
		o.SelfIdentifier.Attributes()["rule"].(types.Number).ValueBigFloat().String(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *PolicyRouteMapRule) GetVyosParentPath() []string {
	return []string{
		"policy",

		"route-map",

		o.SelfIdentifier.Attributes()["route_map"].(types.String).ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *PolicyRouteMapRule) GetVyosNamedParentPath() []string {
	return []string{
		"policy",

		"route-map",

		o.SelfIdentifier.Attributes()["route_map"].(types.String).ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteMapRule) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"rule": schema.NumberAttribute{
					Required: true,
					MarkdownDescription: `Rule for this route-map

    |  Format   |  Description            |
    |-----------|-------------------------|
    |  1-65535  |  Route-map rule number  |
`,
					Description: `Rule for this route-map

    |  Format   |  Description            |
    |-----------|-------------------------|
    |  1-65535  |  Route-map rule number  |
`,
					PlanModifiers: []planmodifier.Number{
						numberplanmodifier.RequiresReplace(),
					},
				},

				"route_map": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `IP route-map

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Route map name  |
`,
					Description: `IP route-map

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Route map name  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in route_map, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
								"illegal character in  route_map, value must match: ^[a-zA-Z0-9-_]*$",
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

		"action": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Action to take on entries matching this rule

    |  Format  |  Description              |
    |----------|---------------------------|
    |  permit  |  Permit matching entries  |
    |  deny    |  Deny matching entries    |
`,
			Description: `Action to take on entries matching this rule

    |  Format  |  Description              |
    |----------|---------------------------|
    |  permit  |  Permit matching entries  |
    |  deny    |  Deny matching entries    |
`,
		},

		"call": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Call another route-map on match

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Route map name  |
`,
			Description: `Call another route-map on match

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Route map name  |
`,
		},

		"continue": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Jump to a different rule in this route-map on a match

    |  Format   |  Description  |
    |-----------|---------------|
    |  1-65535  |  Rule number  |
`,
			Description: `Jump to a different rule in this route-map on a match

    |  Format   |  Description  |
    |-----------|---------------|
    |  1-65535  |  Rule number  |
`,
		},

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

		// Nodes

		"match": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleMatch{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Route parameters to match

`,
			Description: `Route parameters to match

`,
		},

		"on_match": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleOnMatch{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Exit policy on matches

`,
			Description: `Exit policy on matches

`,
		},

		"set": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleSet{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Route parameters

`,
			Description: `Route parameters

`,
		},
	}
}
