// Package namedprotocolspiminterface code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolspiminterface

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r protocolsPimInterface) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `*protocols*  
⯯  
Protocol Independent Multicast (PIM) and IGMP  
⯯  
**PIM interface**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
