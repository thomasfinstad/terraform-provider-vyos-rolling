// Package namedvpnopenconnectauthenticationlocalusersusername code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
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
		MarkdownDescription: `Virtual Private Network (VPN)  
⯯  
SSL VPN OpenConnect, AnyConnect compatible server  
⯯  
Authentication for remote access SSL VPN Server  
⯯  
Local user authentication  
⯯  
**Username used for authentication**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
