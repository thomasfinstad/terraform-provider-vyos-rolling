/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalsystemipnht code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalsystemipnht

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

/* tools/generate-terraform-resource-full/templates/resources/common/schema.gotmpl #schema (nht) */
// Schema method to define the schema for any resource configuration, plan, and state data.
func (r systemIPNht) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `~> This resource is global, having more than one resource of this type is likely to cause configuration drift / conflicts.

*system*  
⯯  
IPv4 Settings  
⯯  
**Filter Next Hop tracking route resolution**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
