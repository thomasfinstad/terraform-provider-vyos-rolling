/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedsystemsyslogconsolefacility code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedsystemsyslogconsolefacility

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

/* tools/generate-terraform-resource-full/templates/resources/common/metadata.gotmpl #metadata (facility) */
// Metadata method to define the resource type name.
func (r systemSyslogConsoleFacility) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_syslog_console_facility"
}
