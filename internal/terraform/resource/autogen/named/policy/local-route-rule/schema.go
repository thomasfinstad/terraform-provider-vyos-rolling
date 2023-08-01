// Package namedpolicylocalrouterule code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedpolicylocalrouterule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r policyLocalRouteRule) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `<div style="text-align: center">
<i>policy</i>

<br>
&darr;
<br>
IPv4 policy route of local traffic

<br>
&darr;
<br>
<b>
Policy local-route rule set number
</b>
</div>
`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
