// Package globalprotocolsisisredistributeipvsixconnectedleveltwo code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsisisredistributeipvsixconnectedleveltwo

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r protocolsIsisRedistributeIPvsixConnectedLevelTwo) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `~> This resource is global, having more than one resource of this type is likely to cause configuration drift / conflicts.

*protocols*  
⯯  
Intermediate System to Intermediate System (IS-IS)  
⯯  
Redistribute information from another routing protocol  
⯯  
Redistribute IPv6 routes  
⯯  
Redistribute connected routes into IS-IS  
⯯  
**Redistribute into level-2**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
