// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// PolicyRouteMapRule describes the resource data model.
type PolicyRouteMapRule struct {
	ID types.Number `tfsdk:"identifier" vyos:",self-id"`

	ParentIDPolicyRouteMap types.String `tfsdk:"route_map" vyos:"route-map_identifier,parent-id"`

	// LeafNodes
	LeafPolicyRouteMapRuleAction      types.String `tfsdk:"action" vyos:"action,omitempty"`
	LeafPolicyRouteMapRuleCall        types.String `tfsdk:"call" vyos:"call,omitempty"`
	LeafPolicyRouteMapRuleContinue    types.Number `tfsdk:"continue" vyos:"continue,omitempty"`
	LeafPolicyRouteMapRuleDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodePolicyRouteMapRuleMatch   *PolicyRouteMapRuleMatch   `tfsdk:"match" vyos:"match,omitempty"`
	NodePolicyRouteMapRuleOnMatch *PolicyRouteMapRuleOnMatch `tfsdk:"on_match" vyos:"on-match,omitempty"`
	NodePolicyRouteMapRuleSet     *PolicyRouteMapRuleSet     `tfsdk:"set" vyos:"set,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *PolicyRouteMapRule) GetVyosPath() []string {
	return []string{
		"policy",

		"route-map",
		o.ParentIDPolicyRouteMap.ValueString(),

		"rule",
		o.ID.ValueBigFloat().String(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteMapRule) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Rule for this route-map

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-65535  |  Route-map rule number  |

`,
		},

		"route_map_identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `IP route-map

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  Route map name  |

`,
		},

		// LeafNodes

		"action": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Action to take on entries matching this rule

    |  Format  |  Description  |
    |----------|---------------|
    |  permit  |  Permit matching entries  |
    |  deny  |  Deny matching entries  |

`,
		},

		"call": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Call another route-map on match

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  Route map name  |

`,
		},

		"continue": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Jump to a different rule in this route-map on a match

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-65535  |  Rule number  |

`,
		},

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  Description  |

`,
		},

		// Nodes

		"match": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleMatch{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Route parameters to match

`,
		},

		"on_match": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleOnMatch{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Exit policy on matches

`,
		},

		"set": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleSet{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Route parameters

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PolicyRouteMapRule) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *PolicyRouteMapRule) UnmarshalJSON(_ []byte) error {
	return nil
}
