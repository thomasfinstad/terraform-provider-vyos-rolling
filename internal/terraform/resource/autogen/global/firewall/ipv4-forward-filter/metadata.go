// Package globalfirewallipvfourforwardfilter code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalfirewallipvfourforwardfilter

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r firewallIPvfourForwardFilter) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_firewall_ipv4_forward_filter"
}
