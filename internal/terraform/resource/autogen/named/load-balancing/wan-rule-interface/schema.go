// Package namedloadbalancingwanruleinterface code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedloadbalancingwanruleinterface

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r loadBalancingWanRuleInterface) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
Rule number (1-9999)

<br>
&darr;
<br>
<b>
Interface name [REQUIRED]
</b>
</div>
`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
