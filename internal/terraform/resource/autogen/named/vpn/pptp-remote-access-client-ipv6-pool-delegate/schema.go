// Package namedvpnpptpremoteaccessclientipvsixpooldelegate code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnpptpremoteaccessclientipvsixpooldelegate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r vpnPptpRemoteAccessClientIPvsixPoolDelegate) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Virtual Private Network (VPN)  
⯯  
Point to Point Tunneling Protocol (PPTP) Virtual Private Network (VPN)  
⯯  
Remote access PPTP VPN  
⯯  
Pool of client IPv6 addresses  
⯯  
**Subnet used to delegate prefix through DHCPv6-PD (RFC3633)**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
