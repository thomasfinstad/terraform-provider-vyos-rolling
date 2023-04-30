// Package namedloadbalancingwanrule code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedloadbalancingwanrule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r loadBalancingWanRule) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Configure load-balancing

Configure Wide Area Network (WAN) load-balancing

Rule number (1-9999)

|  Format  |  Description  |
|----------|---------------|
|  u32:1-9999  |  Rule number  |

`,
		Attributes: r.model.ResourceAttributes(),
	}
}
