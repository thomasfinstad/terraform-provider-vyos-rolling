/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalprotocolsbgpaddressfamilyipvfourlabeledunicastmaximumpaths code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbgpaddressfamilyipvfourlabeledunicastmaximumpaths

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

/* tools/generate-terraform-resource-full/templates/resources/common/metadata.gotmpl #metadata (maximum-paths) */
// Metadata method to define the resource type name.
func (r protocolsBgpAddressFamilyIPvfourLabeledUnicastMaximumPaths) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protocols_bgp_address_family_ipv4_labeled_unicast_maximum_paths"
}
