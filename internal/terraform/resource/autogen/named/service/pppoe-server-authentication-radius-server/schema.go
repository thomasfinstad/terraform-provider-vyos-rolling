// Package namedservicepppoeserverauthenticationradiusserver code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicepppoeserverauthenticationradiusserver

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r servicePppoeServerAuthenticationRadiusServer) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `<div style="text-align: center">
<i>service</i>

<br>
&darr;
<br>
Point to Point over Ethernet (PPPoE) Server

<br>
&darr;
<br>
Authentication for remote access PPPoE Server

<br>
&darr;
<br>
RADIUS based user authentication

<br>
&darr;
<br>
<b>
RADIUS server configuration
</b>
</div>
`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
