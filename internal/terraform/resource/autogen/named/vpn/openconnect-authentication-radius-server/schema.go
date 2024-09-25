// Package namedvpnopenconnectauthenticationradiusserver code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnopenconnectauthenticationradiusserver

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r vpnOpenconnectAuthenticationRadiusServer) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Virtual Private Network (VPN)  
⯯  
SSL VPN OpenConnect, AnyConnect compatible server  
⯯  
Authentication for remote access SSL VPN Server  
⯯  
RADIUS based user authentication  
⯯  
**RADIUS server configuration**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
