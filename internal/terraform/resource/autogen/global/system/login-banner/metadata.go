/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalsystemloginbanner code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalsystemloginbanner

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

/* tools/generate-terraform-resource-full/templates/resources/common/metadata.gotmpl */
// Metadata method to define the resource type name.
func (r systemLoginBanner) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_login_banner"
}
