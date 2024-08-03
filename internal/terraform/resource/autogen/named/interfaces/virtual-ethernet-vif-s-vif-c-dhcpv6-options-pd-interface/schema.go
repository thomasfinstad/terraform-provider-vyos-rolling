// Package namedinterfacesvirtualethernetvifsvifcdhcpvsixoptionspdinterface code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesvirtualethernetvifsvifcdhcpvsixoptionspdinterface

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r interfacesVirtualEthernetVifSVifCDhcpvsixOptionsPdInterface) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `*interfaces*  
⯯  
Virtual Ethernet (veth) Interface  
⯯  
QinQ TAG-S Virtual Local Area Network (VLAN) ID  
⯯  
QinQ TAG-C Virtual Local Area Network (VLAN) ID  
⯯  
DHCPv6 client settings/options  
⯯  
DHCPv6 prefix delegation interface statement  
⯯  
**Delegate IPv6 prefix from provider to this interface**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
