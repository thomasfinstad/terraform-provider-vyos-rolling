/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedvrfnameprotocolsstaticroutesix code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsstaticroutesix

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

/* tools/generate-terraform-resource-full/templates/resources/common/metadata.gotmpl #metadata (route6) */
// Metadata method to define the resource type name.
func (r vrfNameProtocolsStaticRoutesix) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vrf_name_protocols_static_route6"
}
