// Package namedvrfnameprotocolsospfaccesslist code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsospfaccesslist

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r vrfNameProtocolsOspfAccessList) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `<div style="text-align: center">
Virtual Routing and Forwarding

<br>
&darr;
<br>
Virtual Routing and Forwarding instance

<br>
&darr;
<br>
Routing protocol parameters

<br>
&darr;
<br>
Open Shortest Path First (OSPF)

<br>
&darr;
<br>
<b>
Access list to filter networks in routing updates
</b>
</div>
`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
