// Package namedvpnsstpauthenticationradiusserver code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnsstpauthenticationradiusserver

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r vpnSstpAuthenticationRadiusServer) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Secure Socket Tunneling Protocol (SSTP) server

Authentication for remote access SSTP Server

RADIUS based user authentication

RADIUS server configuration

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  RADIUS server IPv4 address  |

`,
		Attributes: r.model.ResourceAttributes(),
	}
}
