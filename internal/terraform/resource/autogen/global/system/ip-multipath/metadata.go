// Package globalsystemipmultipath code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalsystemipmultipath

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r systemIPMultIPath) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_ip_multipath"
}
