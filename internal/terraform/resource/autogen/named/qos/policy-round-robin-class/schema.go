// Package namedqospolicyroundrobinclass code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedqospolicyroundrobinclass

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r qosPolicyRoundRobinClass) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `<div style="text-align: center">
Quality of Service (QoS)

<br>
&darr;
<br>
Service Policy definitions

<br>
&darr;
<br>
Deficit Round Robin Scheduler

<br>
&darr;
<br>
<b>
Class ID
</b>
</div>
`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
