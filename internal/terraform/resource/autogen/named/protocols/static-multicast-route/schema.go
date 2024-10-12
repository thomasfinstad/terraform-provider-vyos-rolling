// Package namedprotocolsstaticmulticastroute code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsstaticmulticastroute

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r protocolsStaticMulticastRoute) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `*protocols*  
⯯  
Static Routing  
⯯  
Multicast static route  
⯯  
**Configure static unicast route into MRIB for multicast RPF lookup**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
