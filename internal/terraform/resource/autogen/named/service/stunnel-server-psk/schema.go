/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedservicestunnelserverpsk code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicestunnelserverpsk

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

/* tools/generate-terraform-resource-full/templates/resources/common/schema.gotmpl #schema (psk) */
// Schema method to define the schema for any resource configuration, plan, and state data.
func (r serviceStunnelServerPsk) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `
*service*  
⯯  
Stunnel TLS Proxy  
⯯  
Stunnel server config  
⯯  
**Pre-shared key name**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
