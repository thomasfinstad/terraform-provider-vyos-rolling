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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (rule) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &PolicyCommunityListRule{}

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-tag-node-identifier (rule) */
// PolicyCommunityListRuleSelfIdentifier used by TagNodes to keep the id info
type PolicyCommunityListRuleSelfIdentifier struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (rule) */

	PolicyCommunityListRule types.Number `tfsdk:"rule" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (community-list) */

	PolicyCommunityList types.String `tfsdk:"community_list" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (policy) */
}

// PolicyCommunityListRule describes the resource data model.
// This is a basenode!
// Top level basenode type: `TagNode`
type PolicyCommunityListRule struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	SelfIdentifier *PolicyCommunityListRuleSelfIdentifier `tfsdk:"identifier" vyos:"-,self-id"`

	// LeafNodes
	LeafPolicyCommunityListRuleAction      types.String `tfsdk:"action" vyos:"action,omitempty"`
	LeafPolicyCommunityListRuleDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafPolicyCommunityListRuleRegex       types.String `tfsdk:"regex" vyos:"regex,omitempty"`

	// TagNodes

	// Nodes
}

// SetID configures the resource ID
func (o *PolicyCommunityListRule) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *PolicyCommunityListRule) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *PolicyCommunityListRule) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *PolicyCommunityListRule) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"rule",
		o.SelfIdentifier.PolicyCommunityListRule.ValueBigFloat().String(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *PolicyCommunityListRule) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (community-list) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (policy) */
		"policy", // Node

		"community-list",
		o.SelfIdentifier.PolicyCommunityList.ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *PolicyCommunityListRule) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (community-list) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (community-list) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (policy) */
		"policy", // Node

		"community-list",
		o.SelfIdentifier.PolicyCommunityList.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyCommunityListRule) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
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
					MarkdownDescription: `Rule for this BGP community list

    |  Format   |  Description                 |
    |-----------|------------------------------|
    |  1-65535  |  Community-list rule number  |
`,
					Description: `Rule for this BGP community list

    |  Format   |  Description                 |
    |-----------|------------------------------|
    |  1-65535  |  Community-list rule number  |
`,
					PlanModifiers: []planmodifier.Number{
						numberplanmodifier.RequiresReplace(),
					},
				},

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (community-list) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (policy) */

				"community_list": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Add a BGP community list entry

    |  Format  |  Description              |
    |----------|---------------------------|
    |  txt     |  BGP community-list name  |
`,
					Description: `Add a BGP community list entry

    |  Format  |  Description              |
    |----------|---------------------------|
    |  txt     |  BGP community-list name  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in community_list, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  community_list, value must match: ^[.:a-zA-Z0-9-_/]+$",
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

		"action":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (action) */
		schema.StringAttribute{
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

		"regex":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (regex) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Regular expression to match against a community-list

    |  Format                      |  Description                                                                 |
    |------------------------------|------------------------------------------------------------------------------|
    |  <aa:nn>                     |  Community number in AA:NN format where AA and NN are (0-65535)              |
    |  local-AS                    |  Well-known communities value NO_EXPORT_SUBCONFED 0xFFFFFF03                 |
    |  no-advertise                |  Well-known communities value NO_ADVERTISE 0xFFFFFF02                        |
    |  no-export                   |  Well-known communities value NO_EXPORT 0xFFFFFF01                           |
    |  internet                    |  Well-known communities value 0                                              |
    |  graceful-shutdown           |  Well-known communities value GRACEFUL_SHUTDOWN 0xFFFF0000                   |
    |  accept-own-nexthop          |  Well-known communities value ACCEPT_OWN_NEXTHOP 0xFFFF0008                  |
    |  accept-own                  |  Well-known communities value ACCEPT_OWN 0xFFFF0001 65535:1                  |
    |  route-filter-translated-v4  |  Well-known communities value ROUTE_FILTER_TRANSLATED_v4 0xFFFF0002 65535:2  |
    |  route-filter-v4             |  Well-known communities value ROUTE_FILTER_v4 0xFFFF0003 65535:3             |
    |  route-filter-translated-v6  |  Well-known communities value ROUTE_FILTER_TRANSLATED_v6 0xFFFF0004 65535:4  |
    |  route-filter-v6             |  Well-known communities value ROUTE_FILTER_v6 0xFFFF0005 65535:5             |
    |  llgr-stale                  |  Well-known communities value LLGR_STALE 0xFFFF0006 65535:6                  |
    |  no-llgr                     |  Well-known communities value NO_LLGR 0xFFFF0007 65535:7                     |
    |  blackhole                   |  Well-known communities value BLACKHOLE 0xFFFF029A 65535:666                 |
    |  no-peer                     |  Well-known communities value NOPEER 0xFFFFFF04 65535:65284                  |
    |  additive                    |  New value is appended to the existing value                                 |
`,
			Description: `Regular expression to match against a community-list

    |  Format                      |  Description                                                                 |
    |------------------------------|------------------------------------------------------------------------------|
    |  <aa:nn>                     |  Community number in AA:NN format where AA and NN are (0-65535)              |
    |  local-AS                    |  Well-known communities value NO_EXPORT_SUBCONFED 0xFFFFFF03                 |
    |  no-advertise                |  Well-known communities value NO_ADVERTISE 0xFFFFFF02                        |
    |  no-export                   |  Well-known communities value NO_EXPORT 0xFFFFFF01                           |
    |  internet                    |  Well-known communities value 0                                              |
    |  graceful-shutdown           |  Well-known communities value GRACEFUL_SHUTDOWN 0xFFFF0000                   |
    |  accept-own-nexthop          |  Well-known communities value ACCEPT_OWN_NEXTHOP 0xFFFF0008                  |
    |  accept-own                  |  Well-known communities value ACCEPT_OWN 0xFFFF0001 65535:1                  |
    |  route-filter-translated-v4  |  Well-known communities value ROUTE_FILTER_TRANSLATED_v4 0xFFFF0002 65535:2  |
    |  route-filter-v4             |  Well-known communities value ROUTE_FILTER_v4 0xFFFF0003 65535:3             |
    |  route-filter-translated-v6  |  Well-known communities value ROUTE_FILTER_TRANSLATED_v6 0xFFFF0004 65535:4  |
    |  route-filter-v6             |  Well-known communities value ROUTE_FILTER_v6 0xFFFF0005 65535:5             |
    |  llgr-stale                  |  Well-known communities value LLGR_STALE 0xFFFF0006 65535:6                  |
    |  no-llgr                     |  Well-known communities value NO_LLGR 0xFFFF0007 65535:7                     |
    |  blackhole                   |  Well-known communities value BLACKHOLE 0xFFFF029A 65535:666                 |
    |  no-peer                     |  Well-known communities value NOPEER 0xFFFFFF04 65535:65284                  |
    |  additive                    |  New value is appended to the existing value                                 |
`,
		},

		// TagNodes

		// Nodes

	}
}
