// Package namedfirewallgroupdynamicgroupipvsixaddressgroup code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallgroupdynamicgroupipvsixaddressgroup

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r firewallGroupDynamicGroupIPvsixAddressGroup) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_firewall_group_dynamic_group_ipv6_address_group"
}
