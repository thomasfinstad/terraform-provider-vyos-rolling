// Package namednatcgnatpoolexternalrange code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namednatcgnatpoolexternalrange

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r natCgnatPoolExternalRange) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Network Address Translation (NAT) parameters  
⯯  
Carrier-grade NAT (CGNAT) parameters  
⯯  
External and internal pool parameters  
⯯  
External pool name  
⯯  
**Range of IP addresses**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
