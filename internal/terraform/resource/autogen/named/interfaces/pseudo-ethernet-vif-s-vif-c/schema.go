// Package namedinterfacespseudoethernetvifsvifc code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
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
		MarkdownDescription: `<div style="text-align: center">
<i>interfaces</i>

<br>
&darr;
<br>
Pseudo Ethernet Interface (Macvlan)

<br>
&darr;
<br>
QinQ TAG-S Virtual Local Area Network (VLAN) ID

<br>
&darr;
<br>
<b>
QinQ TAG-C Virtual Local Area Network (VLAN) ID
</b>
</div>
`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
