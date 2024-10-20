/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedinterfacesltwotpvthree code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesltwotpvthree

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

/* tools/generate-terraform-resource-full/templates/resources/named/schema.gotmpl */
// Schema method to define the schema for any resource configuration, plan, and state data.
func (r interfacesLtwotpvthree) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `*interfaces*  
⯯  
**Layer 2 Tunnel Protocol Version 3 (L2TPv3) Interface**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
