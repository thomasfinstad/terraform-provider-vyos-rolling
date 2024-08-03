// Package namedvpnipsecremoteaccessradiusserver code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnipsecremoteaccessradiusserver

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r vpnIPsecRemoteAccessRadiusServer) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Virtual Private Network (VPN)  
⯯  
VPN IP security (IPsec) parameters  
⯯  
IKEv2 remote access VPN  
⯯  
RADIUS based user authentication  
⯯  
**RADIUS server configuration**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
