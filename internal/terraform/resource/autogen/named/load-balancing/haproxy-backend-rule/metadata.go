// Package namedloadbalancinghaproxybackendrule code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedloadbalancinghaproxybackendrule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r loadBalancingHaproxyBackendRule) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_load_balancing_haproxy_backend_rule"
}
