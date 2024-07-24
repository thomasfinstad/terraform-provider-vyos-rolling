// Package namedvrfnameprotocolsisisinterface code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsisisinterface

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r vrfNameProtocolsIsisInterface) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Virtual Routing and Forwarding  
⯯  
Virtual Routing and Forwarding instance  
⯯  
Routing protocol parameters  
⯯  
Intermediate System to Intermediate System (IS-IS)  
⯯  
**Interface params**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
