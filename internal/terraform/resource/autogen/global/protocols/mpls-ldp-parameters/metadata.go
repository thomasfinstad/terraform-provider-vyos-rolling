/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalprotocolsmplsldpparameters code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsmplsldpparameters

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

/* tools/generate-terraform-resource-full/templates/resources/common/metadata.gotmpl #metadata (parameters) */
// Metadata method to define the resource type name.
func (r protocolsMplsLdpParameters) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protocols_mpls_ldp_parameters"
}
