/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedservicedhcpvsixserversharednetworknamesubnetstaticmapping code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicedhcpvsixserversharednetworknamesubnetstaticmapping

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

/* tools/generate-terraform-resource-full/templates/resources/common/schema.gotmpl */
// Schema method to define the schema for any resource configuration, plan, and state data.
func (r serviceDhcpvsixServerSharedNetworkNameSubnetStaticMapping) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `
*service*  
⯯  
DHCP for IPv6 (DHCPv6) server  
⯯  
DHCPv6 shared network name  
⯯  
IPv6 DHCP subnet for this shared network  
⯯  
**Hostname for static mapping reservation**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
