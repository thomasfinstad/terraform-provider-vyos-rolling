/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedsystemtaskschedulertask code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedsystemtaskschedulertask

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

/* tools/generate-terraform-resource-full/templates/resources/common/metadata.gotmpl #metadata (task) */
// Metadata method to define the resource type name.
func (r systemTaskSchedulerTask) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_task_scheduler_task"
}
