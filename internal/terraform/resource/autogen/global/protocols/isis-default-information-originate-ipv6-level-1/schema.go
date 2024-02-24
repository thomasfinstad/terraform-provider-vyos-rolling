// Package globalprotocolsisisdefaultinformationoriginateipvsixlevelone code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsisisdefaultinformationoriginateipvsixlevelone

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r protocolsIsisDefaultInformationOriginateIPvsixLevelOne) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

<div style="text-align: center">
<i>protocols</i>

<br>
&darr;
<br>
Intermediate System to Intermediate System (IS-IS)

<br>
&darr;
<br>
Control distribution of default information

<br>
&darr;
<br>
Distribute a default route

<br>
&darr;
<br>
Distribute default route for IPv6

<br>
&darr;
<br>
<b>
Distribute default route into level-1
</b>
</div>
`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
