// Package globalsystemip code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalsystemip

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r systemIP) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_ip"
}
