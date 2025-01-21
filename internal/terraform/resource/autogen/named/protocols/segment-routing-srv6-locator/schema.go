/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedprotocolssegmentroutingsrvsixlocator code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolssegmentroutingsrvsixlocator

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

/* tools/generate-terraform-resource-full/templates/resources/common/schema.gotmpl #schema (locator) */
// Schema method to define the schema for any resource configuration, plan, and state data.
func (r protocolsSegmentRoutingSrvsixLocator) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `
*protocols*  
⯯  
Segment Routing  
⯯  
Segment-Routing SRv6 configuration  
⯯  
**Segment Routing SRv6 locator**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
