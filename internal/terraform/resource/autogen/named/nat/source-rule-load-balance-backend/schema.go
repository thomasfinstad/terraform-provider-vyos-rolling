// Package namednatsourceruleloadbalancebackend code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namednatsourceruleloadbalancebackend

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r natSourceRuleLoadBalanceBackend) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Network Address Translation (NAT) parameters  
⯯  
Source NAT settings  
⯯  
Rule number for NAT  
⯯  
Apply NAT load balance  
⯯  
**Translated IP address**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
