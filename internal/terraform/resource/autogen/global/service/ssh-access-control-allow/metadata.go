// Package globalservicetcpaccesscontrolallow code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicetcpaccesscontrolallow

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r serviceTCPAccessControlAllow) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_ssh_access_control_allow"
}
