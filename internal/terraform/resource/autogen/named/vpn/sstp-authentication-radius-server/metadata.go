/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedvpnsstpauthenticationradiusserver code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnsstpauthenticationradiusserver

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

/* tools/generate-terraform-resource-full/templates/resources/common/metadata.gotmpl #metadata (server) */
// Metadata method to define the resource type name.
func (r vpnSstpAuthenticationRadiusServer) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_sstp_authentication_radius_server"
}
