// Package namedfirewallbridgeoutputfilterrule code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallbridgeoutputfilterrule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r firewallBrIDgeOutputFilterRule) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_firewall_bridge_output_filter_rule"
}
