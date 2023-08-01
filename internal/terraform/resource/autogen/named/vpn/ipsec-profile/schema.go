// Package namedvpnipsecprofile code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnipsecprofile

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r vpnIPsecProfile) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `<div style="text-align: center">
Virtual Private Network (VPN)

<br>
&darr;
<br>
VPN IP security (IPsec) parameters

<br>
&darr;
<br>
<b>
VPN IPsec profile
</b>
</div>
`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}