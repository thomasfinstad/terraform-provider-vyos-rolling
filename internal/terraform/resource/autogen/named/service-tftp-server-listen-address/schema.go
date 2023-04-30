// Package namedservicetftpserverlistenaddress code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicetftpserverlistenaddress

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r serviceTftpServerListenAddress) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Trivial File Transfer Protocol (TFTP) server

Local IP addresses to listen on

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address to listen for incoming connections  |
|  ipv6  |  IPv6 address to listen for incoming connections  |

`,
		Attributes: r.model.ResourceAttributes(),
	}
}
