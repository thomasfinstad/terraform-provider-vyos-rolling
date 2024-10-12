// Package namedinterfacesbondingvifdhcpvsixoptionspd code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesbondingvifdhcpvsixoptionspd

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r interfacesBondingVifDhcpvsixOptionsPd) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `*interfaces*  
⯯  
Bonding Interface/Link Aggregation  
⯯  
Virtual Local Area Network (VLAN) ID  
⯯  
DHCPv6 client settings/options  
⯯  
**DHCPv6 prefix delegation interface statement**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
