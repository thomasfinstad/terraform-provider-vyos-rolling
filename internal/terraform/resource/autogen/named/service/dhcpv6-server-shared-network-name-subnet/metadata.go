// Package namedservicedhcpvsixserversharednetworknamesubnet code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicedhcpvsixserversharednetworknamesubnet

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r serviceDhcpvsixServerSharedNetworkNameSubnet) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_dhcpv6_server_shared_network_name_subnet"
}
