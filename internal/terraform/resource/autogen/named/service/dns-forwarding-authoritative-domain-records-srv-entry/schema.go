// Package namedservicednsforwardingauthoritativedomainrecordssrventry code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicednsforwardingauthoritativedomainrecordssrventry

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r serviceDNSForwardingAuthoritativeDomainRecordsSrvEntry) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `<div style="text-align: center">
<i>service</i>

<br>
&darr;
<br>
Domain Name System related services

<br>
&darr;
<br>
DNS forwarding

<br>
&darr;
<br>
Domain to host authoritative records for

<br>
&darr;
<br>
DNS zone records

<br>
&darr;
<br>
"SRV" record

<br>
&darr;
<br>
<b>
Service entry
</b>
</div>
`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
