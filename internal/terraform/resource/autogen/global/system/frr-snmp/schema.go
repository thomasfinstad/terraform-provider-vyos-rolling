/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalsystemfrrsnmp code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalsystemfrrsnmp

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

/* tools/generate-terraform-resource-full/templates/resources/common/schema.gotmpl #schema (snmp) */
// Schema method to define the schema for any resource configuration, plan, and state data.
func (r systemFrrSnmp) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `~> This resource is global, having more than one resource of this type is likely to cause configuration drift / conflicts.

*system*  
⯯  
Configure FRRouting parameters  
⯯  
**Enable SNMP integration for next daemons**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
