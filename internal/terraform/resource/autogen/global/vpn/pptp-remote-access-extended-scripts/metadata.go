/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalvpnpptpremoteaccessextendedscripts code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnpptpremoteaccessextendedscripts

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

/* tools/generate-terraform-resource-full/templates/resources/common/metadata.gotmpl #metadata (extended-scripts) */
// Metadata method to define the resource type name.
func (r vpnPptpRemoteAccessExtendedScrIPts) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_pptp_remote_access_extended_scripts"
}
