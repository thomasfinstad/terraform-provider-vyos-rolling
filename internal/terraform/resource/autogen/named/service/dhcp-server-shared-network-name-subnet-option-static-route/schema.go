/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedservicedhcpserversharednetworknamesubnetoptionstaticroute code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicedhcpserversharednetworknamesubnetoptionstaticroute

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

/* tools/generate-terraform-resource-full/templates/resources/named/schema.gotmpl */
// Schema method to define the schema for any resource configuration, plan, and state data.
func (r serviceDhcpServerSharedNetworkNameSubnetOptionStaticRoute) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `*service*  
⯯  
Dynamic Host Configuration Protocol (DHCP) for DHCP server  
⯯  
Name of DHCP shared network  
⯯  
DHCP subnet for shared network  
⯯  
DHCP option  
⯯  
**Classless static route destination subnet**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
