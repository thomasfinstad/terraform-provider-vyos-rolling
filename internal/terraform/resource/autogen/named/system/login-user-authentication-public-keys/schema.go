// Package namedsystemloginuserauthenticationpublickeys code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedsystemloginuserauthenticationpublickeys

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r systemLoginUserAuthenticationPublicKeys) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `*system*  
⯯  
System User Login Configuration  
⯯  
Local user account information  
⯯  
Authentication settings  
⯯  
**Remote access public keys**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
