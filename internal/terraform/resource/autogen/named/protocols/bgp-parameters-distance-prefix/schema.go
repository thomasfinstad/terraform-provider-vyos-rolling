// Package namedprotocolsbgpparametersdistanceprefix code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsbgpparametersdistanceprefix

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r protocolsBgpParametersDistancePrefix) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `<div style="text-align: center">
<i>protocols</i>

<br>
&darr;
<br>
Border Gateway Protocol (BGP)

<br>
&darr;
<br>
BGP parameters

<br>
&darr;
<br>
Administratives distances for BGP routes

<br>
&darr;
<br>
<b>
Administrative distance for a specific BGP prefix
</b>
</div>
`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
