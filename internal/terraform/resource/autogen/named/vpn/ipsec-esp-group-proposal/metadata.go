// Package namedvpnipsecespgroupproposal code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnipsecespgroupproposal

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r vpnIPsecEspGroupProposal) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_ipsec_esp_group_proposal"
}
