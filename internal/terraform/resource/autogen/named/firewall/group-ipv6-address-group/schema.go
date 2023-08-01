// Package namedfirewallgroupipvsixaddressgroup code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallgroupipvsixaddressgroup

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r firewallGroupIPvsixAddressGroup) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `<div style="text-align: center">
Firewall

<br>
&darr;
<br>
Firewall group

<br>
&darr;
<br>
<b>
Firewall ipv6-address-group
</b>
</div>
`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
