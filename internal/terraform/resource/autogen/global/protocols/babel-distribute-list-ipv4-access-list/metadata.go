/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalprotocolsbabeldistributelistipvfouraccesslist code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbabeldistributelistipvfouraccesslist

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

/* tools/generate-terraform-resource-full/templates/resources/common/metadata.gotmpl #metadata (access-list) */
// Metadata method to define the resource type name.
func (r protocolsBabelDistributeListIPvfourAccessList) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protocols_babel_distribute_list_ipv4_access_list"
}
