/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedvpnipsecremoteaccessconnectionauthenticationlocalusersusername code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnipsecremoteaccessconnectionauthenticationlocalusersusername

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

/* tools/generate-terraform-resource-full/templates/resources/common/schema.gotmpl #schema (username) */
// Schema method to define the schema for any resource configuration, plan, and state data.
func (r vpnIPsecRemoteAccessConnectionAuthenticationLocalUsersUsername) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `
Virtual Private Network (VPN)  
⯯  
VPN IP security (IPsec) parameters  
⯯  
IKEv2 remote access VPN  
⯯  
IKEv2 VPN connection name  
⯯  
Authentication for remote access  
⯯  
Local user authentication  
⯯  
**Username used for authentication**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
