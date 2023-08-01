// Package namedinterfacesethernetvif code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesethernetvif

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r interfacesEthernetVif) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `<div style="text-align: center">
Network interfaces

<br>
&darr;
<br>
Ethernet Interface

<br>
&darr;
<br>
<b>
Virtual Local Area Network (VLAN) ID
</b>
</div>
`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
