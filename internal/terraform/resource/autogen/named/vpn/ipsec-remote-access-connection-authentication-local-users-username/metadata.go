// Package namedvpnipsecremoteaccessconnectionauthenticationlocalusersusername code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnipsecremoteaccessconnectionauthenticationlocalusersusername

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r vpnIPsecRemoteAccessConnectionAuthenticationLocalUsersUsername) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_ipsec_remote_access_connection_authentication_local_users_username"
}
