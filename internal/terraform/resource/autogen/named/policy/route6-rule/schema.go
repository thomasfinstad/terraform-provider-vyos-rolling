// Package namedpolicyroutesixrule code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedpolicyroutesixrule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r policyRoutesixRule) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Routing policy  
⯯  
Policy route rule set name for IPv6  
⯯  
**Policy rule number**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
