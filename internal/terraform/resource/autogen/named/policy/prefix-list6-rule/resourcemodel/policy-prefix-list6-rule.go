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

var _ helpers.VyosTopResourceDataModel = &PolicyPrefixListsixRule{}

// PolicyPrefixListsixRule describes the resource data model.
type PolicyPrefixListsixRule struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafPolicyPrefixListsixRuleAction      types.String `tfsdk:"action" vyos:"action,omitempty"`
	LeafPolicyPrefixListsixRuleDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafPolicyPrefixListsixRuleGe          types.Number `tfsdk:"ge" vyos:"ge,omitempty"`
	LeafPolicyPrefixListsixRuleLe          types.Number `tfsdk:"le" vyos:"le,omitempty"`
	LeafPolicyPrefixListsixRulePrefix      types.String `tfsdk:"prefix" vyos:"prefix,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// SetID configures the resource ID
func (o *PolicyPrefixListsixRule) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *PolicyPrefixListsixRule) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *PolicyPrefixListsixRule) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *PolicyPrefixListsixRule) GetVyosPath() []string {
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
func (o *PolicyPrefixListsixRule) GetVyosParentPath() []string {
	return []string{
		"policy",

		"prefix-list6",

		o.SelfIdentifier.Attributes()["prefix_list6"].(types.String).ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *PolicyPrefixListsixRule) GetVyosNamedParentPath() []string {
	return []string{
		"policy",

		"prefix-list6",

		o.SelfIdentifier.Attributes()["prefix_list6"].(types.String).ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyPrefixListsixRule) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
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
					MarkdownDescription: `Rule for this prefix-list6

    |  Format   |  Description              |
    |-----------|---------------------------|
    |  1-65535  |  Prefix-list rule number  |
`,
					Description: `Rule for this prefix-list6

    |  Format   |  Description              |
    |-----------|---------------------------|
    |  1-65535  |  Prefix-list rule number  |
`,
					PlanModifiers: []planmodifier.Number{
						numberplanmodifier.RequiresReplace(),
					},
				},

				"prefix_list6": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `IPv6 prefix-list filter

    |  Format  |  Description               |
    |----------|----------------------------|
    |  txt     |  Name of IPv6 prefix-list  |
`,
					Description: `IPv6 prefix-list filter

    |  Format  |  Description               |
    |----------|----------------------------|
    |  txt     |  Name of IPv6 prefix-list  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in prefix_list6, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_]+$`),
								"illegal character in  prefix_list6, value must match: ^[.:a-zA-Z0-9-_]+$",
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

		"ge": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Prefix length to match a netmask greater than or equal to it

    |  Format  |  Description                  |
    |----------|-------------------------------|
    |  0-128   |  Netmask greater than length  |
`,
			Description: `Prefix length to match a netmask greater than or equal to it

    |  Format  |  Description                  |
    |----------|-------------------------------|
    |  0-128   |  Netmask greater than length  |
`,
		},

		"le": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Prefix length to match a netmask less than or equal to it

    |  Format  |  Description               |
    |----------|----------------------------|
    |  0-128   |  Netmask less than length  |
`,
			Description: `Prefix length to match a netmask less than or equal to it

    |  Format  |  Description               |
    |----------|----------------------------|
    |  0-128   |  Netmask less than length  |
`,
		},

		"prefix": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Prefix to match

    |  Format   |  Description  |
    |-----------|---------------|
    |  ipv6net  |  IPv6 prefix  |
`,
			Description: `Prefix to match

    |  Format   |  Description  |
    |-----------|---------------|
    |  ipv6net  |  IPv6 prefix  |
`,
		},

		// Nodes

	}
}
