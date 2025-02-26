/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedqospolicyshaperclassmatch code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedqospolicyshaperclassmatch

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

/* tools/generate-terraform-resource-full/templates/resources/common/schema.gotmpl #schema (match) */
// Schema method to define the schema for any resource configuration, plan, and state data.
func (r qosPolicyShaperClassMatch) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `
Quality of Service (QoS)  
⯯  
Service Policy definitions  
⯯  
Traffic shaping based policy (Hierarchy Token Bucket)  
⯯  
Class ID  
⯯  
**Class matching rule name**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
