/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalprotocolsbgpaddressfamilyipvfourunicastlabelvpn code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbgpaddressfamilyipvfourunicastlabelvpn

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

/* tools/generate-terraform-resource-full/templates/resources/common/metadata.gotmpl #metadata (vpn) */
// Metadata method to define the resource type name.
func (r protocolsBgpAddressFamilyIPvfourUnicastLabelVpn) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protocols_bgp_address_family_ipv4_unicast_label_vpn"
}
