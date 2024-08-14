// Package namedvpnipsecespgroupproposal code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnipsecespgroupproposal

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r vpnIPsecEspGroupProposal) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Virtual Private Network (VPN)  
⯯  
VPN IP security (IPsec) parameters  
⯯  
Encapsulating Security Payload (ESP) group name  
⯯  
**ESP group proposal**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}