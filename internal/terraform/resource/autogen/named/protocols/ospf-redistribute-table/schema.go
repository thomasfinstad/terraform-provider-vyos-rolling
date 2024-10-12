// Package namedprotocolsospfredistributetable code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsospfredistributetable

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r protocolsOspfRedistributeTable) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `*protocols*  
⯯  
Open Shortest Path First (OSPF)  
⯯  
Redistribute information from another routing protocol  
⯯  
**Redistribute non-main Kernel Routing Table**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
