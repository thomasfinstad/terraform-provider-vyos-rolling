// Package globalserviceconfigsyncsectionservice code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalserviceconfigsyncsectionservice

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r serviceConfigSyncSectionService) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_config_sync_section_service"
}
