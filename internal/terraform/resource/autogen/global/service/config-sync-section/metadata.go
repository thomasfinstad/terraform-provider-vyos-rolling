// Package globalserviceconfigsyncsection code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalserviceconfigsyncsection

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r serviceConfigSyncSection) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_config_sync_section"
}
