// Package namedservicednsdynamicinterface code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicednsdynamicinterface

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r serviceDNSDynamicInterface) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
Dynamic DNS

<br>
&darr;
<br>
<b>
Interface to send DDNS updates for
</b>
</div>
`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
