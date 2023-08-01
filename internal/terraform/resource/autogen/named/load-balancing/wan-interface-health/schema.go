// Package namedloadbalancingwaninterfacehealth code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedloadbalancingwaninterfacehealth

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r loadBalancingWanInterfaceHealth) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `<div style="text-align: center">
Configure load-balancing

<br>
&darr;
<br>
Configure Wide Area Network (WAN) load-balancing

<br>
&darr;
<br>
<b>
Interface name
</b>
</div>
`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
