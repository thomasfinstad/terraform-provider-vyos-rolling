// Package globalservicedhcpvsixserver code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicedhcpvsixserver

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r serviceDhcpvsixServer) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_dhcpv6_server"
}
