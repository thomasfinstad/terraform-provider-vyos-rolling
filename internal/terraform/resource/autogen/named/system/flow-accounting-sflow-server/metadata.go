/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedsystemflowaccountingsflowserver code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedsystemflowaccountingsflowserver

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

/* tools/generate-terraform-resource-full/templates/resources/common/metadata.gotmpl #metadata (server) */
// Metadata method to define the resource type name.
func (r systemFlowAccountingSflowServer) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_flow_accounting_sflow_server"
}
