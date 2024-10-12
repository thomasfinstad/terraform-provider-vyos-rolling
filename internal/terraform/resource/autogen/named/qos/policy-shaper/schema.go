// Package namedqospolicyshaper code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedqospolicyshaper

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r qosPolicyShaper) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Quality of Service (QoS)  
⯯  
Service Policy definitions  
⯯  
**Traffic shaping based policy (Hierarchy Token Bucket)**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
