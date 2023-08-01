// Package namedprotocolsstatictableroutesix code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsstatictableroutesix

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r protocolsStaticTableRoutesix) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `<div style="text-align: center">
Routing protocols

<br>
&darr;
<br>
Static Routing

<br>
&darr;
<br>
Policy route table number

<br>
&darr;
<br>
<b>
Static IPv6 route
</b>
</div>
`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
