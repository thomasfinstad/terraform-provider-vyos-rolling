/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalprotocolsisisfastreroutelfalocalprioritylimitmedium code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsisisfastreroutelfalocalprioritylimitmedium

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

/* tools/generate-terraform-resource-full/templates/resources/common/metadata.gotmpl #metadata (medium) */
// Metadata method to define the resource type name.
func (r protocolsIsisFastRerouteLfaLocalPriorityLimitMedium) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protocols_isis_fast_reroute_lfa_local_priority_limit_medium"
}
