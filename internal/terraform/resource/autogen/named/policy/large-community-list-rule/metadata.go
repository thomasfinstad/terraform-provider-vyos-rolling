// Package namedpolicylargecommunitylistrule code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedpolicylargecommunitylistrule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r policyLargeCommunityListRule) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policy_large_community_list_rule"
}
