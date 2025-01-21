/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedprotocolsnhrptunnelmaptunnelip code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsnhrptunnelmaptunnelip

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

/* tools/generate-terraform-resource-full/templates/resources/common/schema.gotmpl #schema (tunnel-ip) */
// Schema method to define the schema for any resource configuration, plan, and state data.
func (r protocolsNhrpTunnelMapTunnelIP) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `
*protocols*  
⯯  
Next Hop Resolution Protocol (NHRP) parameters  
⯯  
Tunnel for NHRP  
⯯  
Map tunnel IP to NBMA   
⯯  
**Set a NHRP tunnel address**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
