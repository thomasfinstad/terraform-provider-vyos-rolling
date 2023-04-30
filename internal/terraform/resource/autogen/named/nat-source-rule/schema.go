// Package namednatsourcerule code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namednatsourcerule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r natSourceRule) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Network Address Translation (NAT) parameters

Source NAT settings

Rule number for NAT

|  Format  |  Description  |
|----------|---------------|
|  u32:1-999999  |  Number of NAT rule  |

`,
		Attributes: r.model.ResourceAttributes(),
	}
}
