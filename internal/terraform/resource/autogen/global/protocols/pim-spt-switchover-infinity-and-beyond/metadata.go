/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalprotocolspimsptswitchoverinfinityandbeyond code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolspimsptswitchoverinfinityandbeyond

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

/* tools/generate-terraform-resource-full/templates/resources/common/metadata.gotmpl #metadata (infinity-and-beyond) */
// Metadata method to define the resource type name.
func (r protocolsPimSptSwitchoverInfinityAndBeyond) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protocols_pim_spt_switchover_infinity_and_beyond"
}
