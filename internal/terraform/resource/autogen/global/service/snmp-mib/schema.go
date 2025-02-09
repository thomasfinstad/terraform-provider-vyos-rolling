/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalservicesnmpmib code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicesnmpmib

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

/* tools/generate-terraform-resource-full/templates/resources/common/schema.gotmpl #schema (mib) */
// Schema method to define the schema for any resource configuration, plan, and state data.
func (r serviceSnmpMib) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `~> This resource is global, having more than one resource of this type is likely to cause configuration drift / conflicts.

*service*  
⯯  
Simple Network Management Protocol (SNMP)  
⯯  
**Management information base (MIB)**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
