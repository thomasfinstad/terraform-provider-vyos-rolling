// Package namedpolicylocalrouterule code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedpolicylocalrouterule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r policyLocalRouteRule) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policy_local_route_rule"
}