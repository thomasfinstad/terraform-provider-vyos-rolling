// Package namedqospolicyshaperclassmatch code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedqospolicyshaperclassmatch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r qosPolicyShaperClassMatch) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_qos_policy_shaper_class_match"
}
