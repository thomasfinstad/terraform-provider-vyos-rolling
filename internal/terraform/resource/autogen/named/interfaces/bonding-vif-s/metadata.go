/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedinterfacesbondingvifs code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesbondingvifs

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

/* tools/generate-terraform-resource-full/templates/resources/common/metadata.gotmpl #metadata (vif-s) */
// Metadata method to define the resource type name.
func (r interfacesBondingVifS) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interfaces_bonding_vif_s"
}
