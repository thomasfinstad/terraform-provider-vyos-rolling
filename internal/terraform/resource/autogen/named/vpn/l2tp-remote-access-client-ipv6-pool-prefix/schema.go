// Package namedvpnltwotpremoteaccessclientipvsixpoolprefix code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnltwotpremoteaccessclientipvsixpoolprefix

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r vpnLtwotpRemoteAccessClientIPvsixPoolPrefix) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Virtual Private Network (VPN)  
⯯  
L2TP Virtual Private Network (VPN)  
⯯  
Remote access L2TP VPN  
⯯  
Pool of client IPv6 addresses  
⯯  
**Pool of addresses used to assign to clients**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
