// Package globalvpnipsecremoteaccessradius code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnipsecremoteaccessradius

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r vpnIPsecRemoteAccessRadius) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_ipsec_remote_access_radius"
}
