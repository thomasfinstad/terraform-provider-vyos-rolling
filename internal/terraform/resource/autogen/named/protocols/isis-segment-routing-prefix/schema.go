/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedprotocolsisissegmentroutingprefix code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsisissegmentroutingprefix

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

/* tools/generate-terraform-resource-full/templates/resources/common/schema.gotmpl #schema (prefix) */
// Schema method to define the schema for any resource configuration, plan, and state data.
func (r protocolsIsisSegmentRoutingPrefix) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `
*protocols*  
⯯  
Intermediate System to Intermediate System (IS-IS)  
⯯  
Segment-Routing (SPRING) settings  
⯯  
**Static IPv4/IPv6 prefix segment/label mapping**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
