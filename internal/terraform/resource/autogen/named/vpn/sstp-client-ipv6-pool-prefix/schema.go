// Package namedvpnsstpclientipvsixpoolprefix code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnsstpclientipvsixpoolprefix

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r vpnSstpClientIPvsixPoolPrefix) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `<div style="text-align: center">
<i>vpn</i>

<br>
&darr;
<br>
Secure Socket Tunneling Protocol (SSTP) server

<br>
&darr;
<br>
Pool of client IPv6 addresses

<br>
&darr;
<br>
<b>
Pool of addresses used to assign to clients
</b>
</div>
`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
