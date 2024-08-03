// Package namedservicednsforwardingauthoritativedomainrecordsaaaa code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicednsforwardingauthoritativedomainrecordsaaaa

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r serviceDNSForwardingAuthoritativeDomainRecordsAaaa) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `*service*  
⯯  
Domain Name System (DNS) related services  
⯯  
DNS forwarding  
⯯  
Domain to host authoritative records for  
⯯  
DNS zone records  
⯯  
**AAAA record**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
