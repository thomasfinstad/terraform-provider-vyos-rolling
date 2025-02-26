/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedfirewallbridgeforwardfilterrule code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallbridgeforwardfilterrule

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

/* tools/generate-terraform-resource-full/templates/resources/common/schema.gotmpl #schema (rule) */
// Schema method to define the schema for any resource configuration, plan, and state data.
func (r firewallBrIDgeForwardFilterRule) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `
Firewall  
⯯  
Bridge firewall  
⯯  
Bridge forward firewall  
⯯  
Bridge firewall forward filter  
⯯  
**Bridge Firewall forward filter rule number**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
