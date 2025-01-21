/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalvpnopenconnectlistenports code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnopenconnectlistenports

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

/* tools/generate-terraform-resource-full/templates/resources/common/schema.gotmpl #schema (listen-ports) */
// Schema method to define the schema for any resource configuration, plan, and state data.
func (r vpnOpenconnectListenPorts) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `~> This resource is global, having more than one resource of this type is likely to cause configuration drift / conflicts.

Virtual Private Network (VPN)  
⯯  
SSL VPN OpenConnect, AnyConnect compatible server  
⯯  
**Specify custom ports to use for client connections**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
