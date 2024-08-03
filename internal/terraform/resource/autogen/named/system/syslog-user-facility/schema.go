// Package namedsystemsysloguserfacility code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedsystemsysloguserfacility

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r systemSyslogUserFacility) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `*system*  
⯯  
System logging  
⯯  
Logging to specific terminal of given user  
⯯  
**Facility for logging**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
