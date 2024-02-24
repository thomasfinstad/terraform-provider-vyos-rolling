// Package globalvpnsstpauthenticationradius code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnsstpauthenticationradius

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r vpnSstpAuthenticationRadius) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

<div style="text-align: center">
<i>vpn</i>

<br>
&darr;
<br>
Secure Socket Tunneling Protocol (SSTP) server

<br>
&darr;
<br>
Authentication for remote access SSTP Server

<br>
&darr;
<br>
<b>
RADIUS based user authentication
</b>
</div>
`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
