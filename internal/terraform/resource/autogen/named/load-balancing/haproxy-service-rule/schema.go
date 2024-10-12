// Package namedloadbalancinghaproxyservicerule code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedloadbalancinghaproxyservicerule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r loadBalancingHaproxyServiceRule) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `*load-balancing*  
⯯  
Configure haproxy  
⯯  
Frontend service name  
⯯  
**Proxy rule number**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
