// Package namedvpnsstpauthenticationlocalusersusername code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnsstpauthenticationlocalusersusername

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r vpnSstpAuthenticationLocalUsersUsername) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Virtual Private Network (VPN)  
⯯  
Secure Socket Tunneling Protocol (SSTP) server  
⯯  
Authentication for remote access SSTP Server  
⯯  
Local user authentication for PPPoE server  
⯯  
**User name for authentication**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
