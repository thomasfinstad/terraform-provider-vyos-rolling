/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedvrfnameprotocolsospfinterfaceauthenticationmdfivekeyid code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsospfinterfaceauthenticationmdfivekeyid

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

/* tools/generate-terraform-resource-full/templates/resources/common/schema.gotmpl */
// Schema method to define the schema for any resource configuration, plan, and state data.
func (r vrfNameProtocolsOspfInterfaceAuthenticationMdfiveKeyID) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `
Virtual Routing and Forwarding  
⯯  
Virtual Routing and Forwarding instance  
⯯  
Routing protocol parameters  
⯯  
Open Shortest Path First (OSPF)  
⯯  
Interface configuration  
⯯  
Authentication  
⯯  
MD5 key id  
⯯  
**MD5 key id**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
