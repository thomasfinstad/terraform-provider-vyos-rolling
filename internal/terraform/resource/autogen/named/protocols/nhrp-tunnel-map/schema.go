// Package namedprotocolsnhrptunnelmap code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsnhrptunnelmap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r protocolsNhrpTunnelMap) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `*protocols*  
⯯  
Next Hop Resolution Protocol (NHRP) parameters  
⯯  
Tunnel for NHRP  
⯯  
**Set an HUB tunnel address**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
