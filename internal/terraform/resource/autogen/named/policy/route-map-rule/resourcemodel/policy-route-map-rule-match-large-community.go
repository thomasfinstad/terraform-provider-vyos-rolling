// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// PolicyRouteMapRuleMatchLargeCommunity describes the resource data model.
type PolicyRouteMapRuleMatchLargeCommunity struct {
	// LeafNodes
	LeafPolicyRouteMapRuleMatchLargeCommunityLargeCommunityList types.String `tfsdk:"large_community_list" vyos:"large-community-list,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteMapRuleMatchLargeCommunity) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"large_community_list": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `BGP large-community-list to match

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PolicyRouteMapRuleMatchLargeCommunity) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *PolicyRouteMapRuleMatchLargeCommunity) UnmarshalJSON(_ []byte) error {
	return nil
}
