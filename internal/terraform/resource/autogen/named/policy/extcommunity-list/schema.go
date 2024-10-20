/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedpolicyextcommunitylist code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedpolicyextcommunitylist

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

/* tools/generate-terraform-resource-full/templates/resources/named/schema.gotmpl */
// Schema method to define the schema for any resource configuration, plan, and state data.
func (r policyExtcommunityList) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Routing policy  
⯯  
**Add a BGP extended community list entry**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
