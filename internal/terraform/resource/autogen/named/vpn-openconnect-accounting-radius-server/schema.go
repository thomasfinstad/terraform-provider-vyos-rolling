// Package namedvpnopenconnectaccountingradiusserver code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnopenconnectaccountingradiusserver

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r vpnOpenconnectAccountingRadiusServer) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `SSL VPN OpenConnect, AnyConnect compatible server

Accounting for users OpenConnect VPN Sessions

RADIUS accounting for users OpenConnect VPN sessions OpenConnect authentication mode radius

RADIUS server configuration

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  RADIUS server IPv4 address  |

`,
		Attributes: r.model.ResourceAttributes(),
	}
}
