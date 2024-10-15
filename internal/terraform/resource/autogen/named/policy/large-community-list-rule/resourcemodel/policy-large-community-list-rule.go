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

var _ helpers.VyosTopResourceDataModel = &PolicyLargeCommunityListRule{}

// PolicyLargeCommunityListRule describes the resource data model.
type PolicyLargeCommunityListRule struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafPolicyLargeCommunityListRuleAction      types.String `tfsdk:"action" vyos:"action,omitempty"`
	LeafPolicyLargeCommunityListRuleDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafPolicyLargeCommunityListRuleRegex       types.String `tfsdk:"regex" vyos:"regex,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// SetID configures the resource ID
func (o *PolicyLargeCommunityListRule) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *PolicyLargeCommunityListRule) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *PolicyLargeCommunityListRule) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *PolicyLargeCommunityListRule) GetVyosPath() []string {
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
func (o *PolicyLargeCommunityListRule) GetVyosParentPath() []string {
	return []string{
		"policy",

		"large-community-list",

		o.SelfIdentifier.Attributes()["large_community_list"].(types.String).ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *PolicyLargeCommunityListRule) GetVyosNamedParentPath() []string {
	return []string{
		"policy",

		"large-community-list",

		o.SelfIdentifier.Attributes()["large_community_list"].(types.String).ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyLargeCommunityListRule) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
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
					MarkdownDescription: `Rule for this BGP extended community list

    |  Format   |  Description                       |
    |-----------|------------------------------------|
    |  1-65535  |  Large community-list rule number  |
`,
					Description: `Rule for this BGP extended community list

    |  Format   |  Description                       |
    |-----------|------------------------------------|
    |  1-65535  |  Large community-list rule number  |
`,
					PlanModifiers: []planmodifier.Number{
						numberplanmodifier.RequiresReplace(),
					},
				},

				"large_community_list": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Add a BGP large community list entry

    |  Format  |  Description                    |
    |----------|---------------------------------|
    |  txt     |  BGP large-community-list name  |
`,
					Description: `Add a BGP large community list entry

    |  Format  |  Description                    |
    |----------|---------------------------------|
    |  txt     |  BGP large-community-list name  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in large_community_list, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_]+$`),
								"illegal character in  large_community_list, value must match: ^[.:a-zA-Z0-9-_]+$",
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

		"regex": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Regular expression to match against a large community list

    |  Format     |  Description                                            |
    |-------------|---------------------------------------------------------|
    |  ASN:NN:NN  |  BGP large-community-list filter                        |
    |  IP:NN:NN   |  BGP large-community-list filter (IPv4 address format)  |
`,
			Description: `Regular expression to match against a large community list

    |  Format     |  Description                                            |
    |-------------|---------------------------------------------------------|
    |  ASN:NN:NN  |  BGP large-community-list filter                        |
    |  IP:NN:NN   |  BGP large-community-list filter (IPv4 address format)  |
`,
		},

		// Nodes

	}
}
