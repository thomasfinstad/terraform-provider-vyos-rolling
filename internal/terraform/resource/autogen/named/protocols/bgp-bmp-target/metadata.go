// Package namedprotocolsbgpbmptarget code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsbgpbmptarget

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r protocolsBgpBmpTarget) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protocols_bgp_bmp_target"
}
