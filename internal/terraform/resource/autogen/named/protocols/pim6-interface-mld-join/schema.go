// Package namedprotocolspimsixinterfacemldjoin code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolspimsixinterfacemldjoin

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r protocolsPimsixInterfaceMldJoin) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `*protocols*  
⯯  
Protocol Independent Multicast for IPv6 (PIMv6) and MLD  
⯯  
PIMv6 interface  
⯯  
Multicast Listener Discovery (MLD)  
⯯  
**MLD join multicast group**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
