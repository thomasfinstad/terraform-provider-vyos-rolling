/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedprotocolsripinterfaceauthenticationmdfive code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsripinterfaceauthenticationmdfive

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

/* tools/generate-terraform-resource-full/templates/resources/common/metadata.gotmpl #metadata (md5) */
// Metadata method to define the resource type name.
func (r protocolsRIPInterfaceAuthenticationMdfive) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protocols_rip_interface_authentication_md5"
}
