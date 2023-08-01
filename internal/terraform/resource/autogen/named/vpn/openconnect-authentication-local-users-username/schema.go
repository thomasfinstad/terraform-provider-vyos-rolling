// Package namedvpnopenconnectauthenticationlocalusersusername code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnopenconnectauthenticationlocalusersusername

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r vpnOpenconnectAuthenticationLocalUsersUsername) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `<div style="text-align: center">
<i>vpn</i>

<br>
&darr;
<br>
SSL VPN OpenConnect, AnyConnect compatible server

<br>
&darr;
<br>
Authentication for remote access SSL VPN Server

<br>
&darr;
<br>
Local user authentication

<br>
&darr;
<br>
<b>
Username used for authentication
</b>
</div>
`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}