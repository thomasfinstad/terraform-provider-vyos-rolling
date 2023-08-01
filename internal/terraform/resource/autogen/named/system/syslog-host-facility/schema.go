// Package namedsystemsysloghostfacility code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedsystemsysloghostfacility

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r systemSyslogHostFacility) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `<div style="text-align: center">
<i>system</i>

<br>
&darr;
<br>
System logging

<br>
&darr;
<br>
Logging to a remote host

<br>
&darr;
<br>
<b>
Facility for logging
</b>
</div>
`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
