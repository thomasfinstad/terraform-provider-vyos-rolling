// Package namedserviceipoeserverauthenticationinterfacemac code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedserviceipoeserverauthenticationinterfacemac

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r serviceIPoeServerAuthenticationInterfaceMac) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	r.ResourceName = req.ProviderTypeName + "_service_ipoe_server_authentication_interface_mac"
	resp.TypeName = r.ResourceName
}