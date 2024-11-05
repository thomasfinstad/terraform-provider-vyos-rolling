/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedinterfacesvirtualethernetvifsdhcpvsixoptionspd code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesvirtualethernetvifsdhcpvsixoptionspd

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

/* tools/generate-terraform-resource-full/templates/resources/common/schema.gotmpl */
// Schema method to define the schema for any resource configuration, plan, and state data.
func (r interfacesVirtualEthernetVifSDhcpvsixOptionsPd) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `
*interfaces*  
⯯  
Virtual Ethernet (veth) Interface  
⯯  
QinQ TAG-S Virtual Local Area Network (VLAN) ID  
⯯  
DHCPv6 client settings/options  
⯯  
**DHCPv6 prefix delegation interface statement**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
