// Package globalvpnltwotpremoteaccessipsecsettingsauthentication code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnltwotpremoteaccessipsecsettingsauthentication

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r vpnLtwotpRemoteAccessIPsecSettingsAuthentication) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

Virtual Private Network (VPN)  
⯯  
L2TP Virtual Private Network (VPN)  
⯯  
Remote access L2TP VPN  
⯯  
Internet Protocol Security (IPsec) for remote access L2TP VPN  
⯯  
**IPsec authentication settings**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
