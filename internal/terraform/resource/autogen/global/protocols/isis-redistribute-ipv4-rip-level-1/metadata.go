/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalprotocolsisisredistributeipvfourriplevelone code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsisisredistributeipvfourriplevelone

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

/* tools/generate-terraform-resource-full/templates/resources/common/metadata.gotmpl #metadata (level-1) */
// Metadata method to define the resource type name.
func (r protocolsIsisRedistributeIPvfourRIPLevelOne) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protocols_isis_redistribute_ipv4_rip_level_1"
}
