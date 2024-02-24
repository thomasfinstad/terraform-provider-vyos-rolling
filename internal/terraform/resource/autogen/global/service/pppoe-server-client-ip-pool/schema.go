// Package globalservicepppoeserverclientippool code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicepppoeserverclientippool

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r servicePppoeServerClientIPPool) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

<div style="text-align: center">
<i>service</i>

<br>
&darr;
<br>
Point to Point over Ethernet (PPPoE) Server

<br>
&darr;
<br>
<b>
Pool of client IP addresses (must be within a /24)
</b>
</div>
`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}