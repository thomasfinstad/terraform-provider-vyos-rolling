// Package namedhighavailabilityvrrpsyncgroup code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedhighavailabilityvrrpsyncgroup

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r highAvailabilityVrrpSyncGroup) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `High availability settings  
⯯  
Virtual Router Redundancy Protocol settings  
⯯  
**VRRP sync group**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
