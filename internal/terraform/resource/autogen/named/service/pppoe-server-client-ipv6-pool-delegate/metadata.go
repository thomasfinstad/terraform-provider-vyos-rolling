// Package namedservicepppoeserverclientipvsixpooldelegate code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicepppoeserverclientipvsixpooldelegate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r servicePppoeServerClientIPvsixPoolDelegate) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_pppoe_server_client_ipv6_pool_delegate"
}
