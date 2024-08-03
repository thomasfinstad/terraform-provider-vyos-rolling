// Package namednatcgnatpoolexternal code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namednatcgnatpoolexternal

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r natCgnatPoolExternal) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `*nat*  
⯯  
Carrier-grade NAT (CGNAT) parameters  
⯯  
External and internal pool parameters  
⯯  
**External pool name**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
