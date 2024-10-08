// Package globalprotocolsripredistributeisis code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsripredistributeisis

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r protocolsRIPRedistributeIsis) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `~> This resource is global, having more than one resource of this type is likely to cause configuration drift / conflicts.

*protocols*  
⯯  
Routing Information Protocol (RIP) parameters  
⯯  
Redistribute information from another routing protocol  
⯯  
**Redistribute IS-IS routes**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
