// Package globalsystemipvsix code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalsystemipvsix

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r systemIPvsix) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_ipv6"
}
