// Package globalprotocolsisisredistributeipvfourstaticlevelone code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsisisredistributeipvfourstaticlevelone

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r protocolsIsisRedistributeIPvfourStaticLevelOne) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `~> This resource is global, having more than one resource of this type is likely to cause configuration drift / conflicts.

*protocols*  
⯯  
Intermediate System to Intermediate System (IS-IS)  
⯯  
Redistribute information from another routing protocol  
⯯  
Redistribute IPv4 routes  
⯯  
Redistribute static routes into IS-IS  
⯯  
**Redistribute into level-1**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
