/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedvpnipsecauthenticationpsk code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnipsecauthenticationpsk

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

/* tools/generate-terraform-resource-full/templates/resources/common/metadata.gotmpl #metadata (psk) */
// Metadata method to define the resource type name.
func (r vpnIPsecAuthenticationPsk) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_ipsec_authentication_psk"
}
