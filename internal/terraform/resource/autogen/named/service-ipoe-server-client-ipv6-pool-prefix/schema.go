// Package namedserviceipoeserverclientipvsixpoolprefix code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedserviceipoeserverclientipvsixpoolprefix

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r serviceIPoeServerClientIPvsixPoolPrefix) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Internet Protocol over Ethernet (IPoE) Server

Pool of client IPv6 addresses

Pool of addresses used to assign to clients

|  Format  |  Description  |
|----------|---------------|
|  ipv6net  |  IPv6 address and prefix length  |

`,
		Attributes: r.model.ResourceAttributes(),
	}
}
