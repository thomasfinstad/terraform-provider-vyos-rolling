// Package namedserviceipoeserverclientipvsixpoolprefix code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedserviceipoeserverclientipvsixpoolprefix

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r serviceIPoeServerClientIPvsixPoolPrefix) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_ipoe_server_client_ipv6_pool_prefix"
}
