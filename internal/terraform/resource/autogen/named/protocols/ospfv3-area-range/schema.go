// Package namedprotocolsospfvthreearearange code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsospfvthreearearange

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r protocolsOspfvthreeAreaRange) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `*protocols*  
⯯  
Open Shortest Path First (OSPF) for IPv6  
⯯  
OSPFv3 Area  
⯯  
**Specify IPv6 prefix (border routers only)**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}