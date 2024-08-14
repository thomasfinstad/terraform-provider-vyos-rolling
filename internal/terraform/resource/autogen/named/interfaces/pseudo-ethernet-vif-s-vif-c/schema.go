// Package namedinterfacespseudoethernetvifsvifc code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacespseudoethernetvifsvifc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r interfacesPseudoEthernetVifSVifC) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `*interfaces*  
⯯  
Pseudo Ethernet Interface (Macvlan)  
⯯  
QinQ TAG-S Virtual Local Area Network (VLAN) ID  
⯯  
**QinQ TAG-C Virtual Local Area Network (VLAN) ID**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}