// Package globalfirewallipvsixforwardfilter code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalfirewallipvsixforwardfilter

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r firewallIPvsixForwardFilter) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_firewall_ipv6_forward_filter"
}
