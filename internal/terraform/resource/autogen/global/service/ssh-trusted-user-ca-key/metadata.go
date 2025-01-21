/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalservicetcptrustedusercakey code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicetcptrustedusercakey

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

/* tools/generate-terraform-resource-full/templates/resources/common/metadata.gotmpl #metadata (trusted-user-ca-key) */
// Metadata method to define the resource type name.
func (r serviceTCPTrustedUserCaKey) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_ssh_trusted_user_ca_key"
}
