// Package namedservicedhcpserversharednetworkname code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicedhcpserversharednetworkname

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r serviceDhcpServerSharedNetworkName) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `<div style="text-align: center">
<i>service</i>

<br>
&darr;
<br>
Dynamic Host Configuration Protocol (DHCP) for DHCP server

<br>
&darr;
<br>
<b>
Name of DHCP shared network
</b>
</div>
`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
