// Package namedinterfaceswirelessvifs code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfaceswirelessvifs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r interfacesWirelessVifS) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `*interfaces*  
⯯  
Wireless (WiFi/WLAN) Network Interface  
⯯  
**QinQ TAG-S Virtual Local Area Network (VLAN) ID**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}