// Package globalvpnltwotpremoteaccesslog code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnltwotpremoteaccesslog

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r vpnLtwotpRemoteAccessLog) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_l2tp_remote_access_log"
}
