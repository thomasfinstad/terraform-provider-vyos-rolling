// Package namednatcgnatrule code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namednatcgnatrule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r natCgnatRule) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nat_cgnat_rule"
}
