/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedservicesnmpvthreegroup code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicesnmpvthreegroup

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

/* tools/generate-terraform-resource-full/templates/resources/common/metadata.gotmpl #metadata (group) */
// Metadata method to define the resource type name.
func (r serviceSnmpVthreeGroup) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_snmp_v3_group"
}
