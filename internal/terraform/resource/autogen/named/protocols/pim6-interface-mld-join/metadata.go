/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedprotocolspimsixinterfacemldjoin code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolspimsixinterfacemldjoin

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

/* tools/generate-terraform-resource-full/templates/resources/common/metadata.gotmpl #metadata (join) */
// Metadata method to define the resource type name.
func (r protocolsPimsixInterfaceMldJoin) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protocols_pim6_interface_mld_join"
}
