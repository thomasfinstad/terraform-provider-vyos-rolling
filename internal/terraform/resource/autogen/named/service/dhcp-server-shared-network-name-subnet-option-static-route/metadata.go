// Package namedservicedhcpserversharednetworknamesubnetoptionstaticroute code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicedhcpserversharednetworknamesubnetoptionstaticroute

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r serviceDhcpServerSharedNetworkNameSubnetOptionStaticRoute) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_dhcp_server_shared_network_name_subnet_option_static_route"
}
