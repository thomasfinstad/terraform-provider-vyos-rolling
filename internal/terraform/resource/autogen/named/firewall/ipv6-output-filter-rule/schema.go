/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedfirewallipvsixoutputfilterrule code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallipvsixoutputfilterrule

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

/* tools/generate-terraform-resource-full/templates/resources/common/schema.gotmpl #schema (rule) */
// Schema method to define the schema for any resource configuration, plan, and state data.
func (r firewallIPvsixOutputFilterRule) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `
Firewall  
⯯  
IPv6 firewall  
⯯  
IPv6 output firewall  
⯯  
IPv6 firewall output filter  
⯯  
**IPv6 Firewall output filter rule number**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
