/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedprotocolsospfvthreearea code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsospfvthreearea

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

/* tools/generate-terraform-resource-full/templates/resources/common/metadata.gotmpl #metadata (area) */
// Metadata method to define the resource type name.
func (r protocolsOspfvthreeArea) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protocols_ospfv3_area"
}
