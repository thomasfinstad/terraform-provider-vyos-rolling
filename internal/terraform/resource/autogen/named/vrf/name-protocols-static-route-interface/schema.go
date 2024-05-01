// Package namedvrfnameprotocolsstaticrouteinterface code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsstaticrouteinterface

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r vrfNameProtocolsStaticRouteInterface) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Virtual Routing and Forwarding  
⯯  
Virtual Routing and Forwarding instance  
⯯  
Routing protocol parameters  
⯯  
Static Routing  
⯯  
Static IPv4 route  
⯯  
**Next-hop IPv4 router interface**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
