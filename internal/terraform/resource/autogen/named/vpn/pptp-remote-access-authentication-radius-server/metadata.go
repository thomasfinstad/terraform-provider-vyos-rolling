/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedvpnpptpremoteaccessauthenticationradiusserver code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnpptpremoteaccessauthenticationradiusserver

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

/* tools/generate-terraform-resource-full/templates/resources/common/metadata.gotmpl */
// Metadata method to define the resource type name.
func (r vpnPptpRemoteAccessAuthenticationRadiusServer) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_pptp_remote_access_authentication_radius_server"
}
