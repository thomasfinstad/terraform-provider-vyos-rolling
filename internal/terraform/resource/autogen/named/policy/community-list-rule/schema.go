// Package namedpolicycommunitylistrule code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedpolicycommunitylistrule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r policyCommunityListRule) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Routing policy  
⯯  
Add a BGP community list entry  
⯯  
**Rule for this BGP community list**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
