// Package namedservicedhcpserversharednetworknameoptionstaticroute code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicedhcpserversharednetworknameoptionstaticroute

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r serviceDhcpServerSharedNetworkNameOptionStaticRoute) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `*service*  
⯯  
Dynamic Host Configuration Protocol (DHCP) for DHCP server  
⯯  
Name of DHCP shared network  
⯯  
DHCP option  
⯯  
**Classless static route destination subnet**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
