/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalservicehttpsallowclient code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicehttpsallowclient

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

/* tools/generate-terraform-resource-full/templates/resources/common/metadata.gotmpl #metadata (allow-client) */
// Metadata method to define the resource type name.
func (r serviceHTTPSAllowClient) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_https_allow_client"
}
