/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalprotocolsospftimersthrottlespf code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsospftimersthrottlespf

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

/* tools/generate-terraform-resource-full/templates/resources/common/schema.gotmpl #schema (spf) */
// Schema method to define the schema for any resource configuration, plan, and state data.
func (r protocolsOspfTimersThroTTLeSpf) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `~> This resource is global, having more than one resource of this type is likely to cause configuration drift / conflicts.

*protocols*  
⯯  
Open Shortest Path First (OSPF)  
⯯  
Adjust routing timers  
⯯  
Throttling adaptive timers  
⯯  
**OSPF SPF timers**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
