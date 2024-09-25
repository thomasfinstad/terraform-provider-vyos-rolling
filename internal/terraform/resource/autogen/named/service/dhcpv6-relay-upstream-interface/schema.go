// Package namedservicedhcpvsixrelayupstreaminterface code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicedhcpvsixrelayupstreaminterface

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r serviceDhcpvsixRelayUpstreamInterface) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `*service*  
⯯  
DHCPv6 Relay Agent parameters  
⯯  
**Interface for DHCPv6 Relay Agent forward requests out**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
