// Package namedprotocolsnhrptunneldynamicmap code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsnhrptunneldynamicmap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r protocolsNhrpTunnelDynamicMap) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `<div style="text-align: center">
<i>protocols</i>

<br>
&darr;
<br>
Next Hop Resolution Protocol (NHRP) parameters

<br>
&darr;
<br>
Tunnel for NHRP

<br>
&darr;
<br>
<b>
Set an HUB tunnel address
</b>
</div>
`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}