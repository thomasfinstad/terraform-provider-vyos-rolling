/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedservicedhcpvsixrelaylisteninterface code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicedhcpvsixrelaylisteninterface

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

/* tools/generate-terraform-resource-full/templates/resources/common/schema.gotmpl */
// Schema method to define the schema for any resource configuration, plan, and state data.
func (r serviceDhcpvsixRelayListenInterface) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `
*service*  
⯯  
DHCPv6 Relay Agent parameters  
⯯  
**Interface for DHCPv6 Relay Agent to listen for requests**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
