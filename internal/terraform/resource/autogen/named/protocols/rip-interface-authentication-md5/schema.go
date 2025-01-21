/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedprotocolsripinterfaceauthenticationmdfive code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsripinterfaceauthenticationmdfive

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

/* tools/generate-terraform-resource-full/templates/resources/common/schema.gotmpl #schema (md5) */
// Schema method to define the schema for any resource configuration, plan, and state data.
func (r protocolsRIPInterfaceAuthenticationMdfive) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `
*protocols*  
⯯  
Routing Information Protocol (RIP) parameters  
⯯  
Interface name  
⯯  
Authentication  
⯯  
**MD5 key id**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
