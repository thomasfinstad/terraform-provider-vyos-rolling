// Package globalsystemflowaccountingnetflowtimeout code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalsystemflowaccountingnetflowtimeout

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r systemFlowAccountingNetflowTimeout) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

<div style="text-align: center">
<i>system</i>

<br>
&darr;
<br>
Flow accounting settings

<br>
&darr;
<br>
NetFlow settings

<br>
&darr;
<br>
<b>
NetFlow timeout values
</b>
</div>
`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
