// Package namedqospolicyratecontrol code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedqospolicyratecontrol

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r qosPolicyRateControl) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_qos_policy_rate_control"
}
