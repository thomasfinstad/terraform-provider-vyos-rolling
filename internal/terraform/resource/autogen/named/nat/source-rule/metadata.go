// Package namednatsourcerule code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namednatsourcerule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r natSourceRule) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nat_source_rule"
}
