// Package namedprotocolsbgpaddressfamilyipvfourvpnnetwork code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsbgpaddressfamilyipvfourvpnnetwork

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r protocolsBgpAddressFamilyIPvfourVpnNetwork) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `<div style="text-align: center">
<i>protocols</i>

<br>
&darr;
<br>
Border Gateway Protocol (BGP)

<br>
&darr;
<br>
BGP address-family parameters

<br>
&darr;
<br>
Unicast VPN IPv4 BGP settings

<br>
&darr;
<br>
<b>
Import BGP network/prefix into unicast VPN IPv4 RIB
</b>
</div>
`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}