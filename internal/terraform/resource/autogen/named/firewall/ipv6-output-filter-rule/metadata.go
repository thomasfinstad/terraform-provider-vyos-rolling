// Package namedfirewallipvsixoutputfilterrule code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallipvsixoutputfilterrule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r firewallIPvsixOutputFilterRule) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_firewall_ipv6_output_filter_rule"
}
