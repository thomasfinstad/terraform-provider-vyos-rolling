/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedprotocolsstatictableroutesix code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsstatictableroutesix

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

/* tools/generate-terraform-resource-full/templates/resources/common/metadata.gotmpl #metadata (route6) */
// Metadata method to define the resource type name.
func (r protocolsStaticTableRoutesix) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protocols_static_table_route6"
}
