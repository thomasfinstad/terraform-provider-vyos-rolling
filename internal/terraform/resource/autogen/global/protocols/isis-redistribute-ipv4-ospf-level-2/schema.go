/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalprotocolsisisredistributeipvfourospfleveltwo code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsisisredistributeipvfourospfleveltwo

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

/* tools/generate-terraform-resource-full/templates/resources/global/schema.gotmpl */
// Schema method to define the schema for any resource configuration, plan, and state data.
func (r protocolsIsisRedistributeIPvfourOspfLevelTwo) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
Redistribute OSPF routes into IS-IS  
⯯  
**Redistribute into level-2**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
