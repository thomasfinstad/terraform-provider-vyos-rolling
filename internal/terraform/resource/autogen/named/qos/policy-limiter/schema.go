// Package namedqospolicylimiter code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedqospolicylimiter

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r qosPolicyLimiter) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
<b>
Traffic input limiting policy
</b>
</div>
`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
