/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedvrfnameprotocolsisisfastreroutelfalocaltiebreakernodeprotectingindex code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsisisfastreroutelfalocaltiebreakernodeprotectingindex

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

/* tools/generate-terraform-resource-full/templates/resources/common/schema.gotmpl */
// Schema method to define the schema for any resource configuration, plan, and state data.
func (r vrfNameProtocolsIsisFastRerouteLfaLocalTiebreakerNodeProtectingIndex) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `
Virtual Routing and Forwarding  
⯯  
Virtual Routing and Forwarding instance  
⯯  
Routing protocol parameters  
⯯  
Intermediate System to Intermediate System (IS-IS)  
⯯  
IS-IS fast reroute configuration  
⯯  
Loop free alternate functionality  
⯯  
Local loop free alternate options  
⯯  
Configure tiebreaker for multiple backups  
⯯  
Prefer node protecting backup path  
⯯  
**Set preference order among tiebreakers**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
