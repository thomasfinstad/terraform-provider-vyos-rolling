/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalservicepppoeserverauthenticationradiusdynamicauthor code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicepppoeserverauthenticationradiusdynamicauthor

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

/* tools/generate-terraform-resource-full/templates/resources/common/metadata.gotmpl #metadata (dynamic-author) */
// Metadata method to define the resource type name.
func (r servicePppoeServerAuthenticationRadiusDynamicAuthor) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_pppoe_server_authentication_radius_dynamic_author"
}
