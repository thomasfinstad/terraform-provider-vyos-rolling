// Package globalfirewallglobaloptionsstatepolicyinvalid code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalfirewallglobaloptionsstatepolicyinvalid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r firewallGlobalOptionsStatePolicyInvalID) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_firewall_global_options_state_policy_invalid"
}
