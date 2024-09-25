// Package namedservicedhcpvsixserversharednetworknamesubnet code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicedhcpvsixserversharednetworknamesubnet

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r serviceDhcpvsixServerSharedNetworkNameSubnet) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `*service*  
⯯  
DHCP for IPv6 (DHCPv6) server  
⯯  
DHCPv6 shared network name  
⯯  
**IPv6 DHCP subnet for this shared network**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}