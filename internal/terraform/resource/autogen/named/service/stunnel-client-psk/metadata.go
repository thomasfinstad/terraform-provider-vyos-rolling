// Package namedservicestunnelclientpsk code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicestunnelclientpsk

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r serviceStunnelClientPsk) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_stunnel_client_psk"
}