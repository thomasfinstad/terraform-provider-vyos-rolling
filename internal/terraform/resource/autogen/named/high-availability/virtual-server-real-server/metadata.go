// Package namedhighavailabilityvirtualserverrealserver code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedhighavailabilityvirtualserverrealserver

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r highAvailabilityVirtualServerRealServer) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_high_availability_virtual_server_real_server"
}
