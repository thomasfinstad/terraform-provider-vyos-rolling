/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedvpnltwotpremoteaccessclientipvsixpool code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnltwotpremoteaccessclientipvsixpool

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

/* tools/generate-terraform-resource-full/templates/resources/common/metadata.gotmpl */
// Metadata method to define the resource type name.
func (r vpnLtwotpRemoteAccessClientIPvsixPool) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_l2tp_remote_access_client_ipv6_pool"
}
