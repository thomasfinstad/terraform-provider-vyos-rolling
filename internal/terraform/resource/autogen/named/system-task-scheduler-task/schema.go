// Package namedsystemtaskschedulertask code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedsystemtaskschedulertask

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r systemTaskSchedulerTask) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Task scheduler settings

Scheduled task

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Task name  |

`,
		Attributes: r.model.ResourceAttributes(),
	}
}
