// Package namedservicepppoeserverauthenticationradiusserver code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicepppoeserverauthenticationradiusserver

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r servicePppoeServerAuthenticationRadiusServer) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_pppoe_server_authentication_radius_server"
}