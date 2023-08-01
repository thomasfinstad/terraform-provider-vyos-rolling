// Package namedinterfacesbondingvifsdhcpvsixoptionspd code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesbondingvifsdhcpvsixoptionspd

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r interfacesBondingVifSDhcpvsixOptionsPd) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `<div style="text-align: center">
<i>interfaces</i>

<br>
&darr;
<br>
Bonding Interface/Link Aggregation

<br>
&darr;
<br>
QinQ TAG-S Virtual Local Area Network (VLAN) ID

<br>
&darr;
<br>
DHCPv6 client settings/options

<br>
&darr;
<br>
<b>
DHCPv6 prefix delegation interface statement
</b>
</div>
`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
