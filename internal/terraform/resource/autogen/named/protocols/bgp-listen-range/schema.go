/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedprotocolsbgplistenrange code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsbgplistenrange

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

/* tools/generate-terraform-resource-full/templates/resources/common/schema.gotmpl #schema (range) */
// Schema method to define the schema for any resource configuration, plan, and state data.
func (r protocolsBgpListenRange) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `
*protocols*  
⯯  
Border Gateway Protocol (BGP)  
⯯  
Listen for and accept BGP dynamic neighbors from range  
⯯  
**BGP dynamic neighbors listen range**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
